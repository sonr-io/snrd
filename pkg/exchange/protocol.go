package exchange

import (
	"context"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/common"
	"google.golang.org/protobuf/proto"
)

var (
	logger         = golog.Child("protocols/exchange")
	ErrParameters  = errors.New("Failed to create new ExchangeProtocol, invalid parameters")
	ErrInvalidPeer = errors.New("Peer object provided to ExchangeProtocol is Nil")
)

// ExchangeProtocol handles Global Sonr Exchange Protocol
type ExchangeProtocol struct {
	node           api.NodeImpl
	ctx            context.Context
	host           *host.SNRHost // host
	mode           DNSMode
	namebaseClient *NamebaseAPIClient
	authRecord     host.Record
	nameRecord     host.Record
}

// NewProtocol creates new ExchangeProtocol
func NewProtocol(ctx context.Context, host *host.SNRHost, node api.NodeImpl, options ...Option) (*ExchangeProtocol, error) {
	// Set options
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}

	// Create Exchange Protocol
	exchProtocol := &ExchangeProtocol{
		ctx:            ctx,
		host:           host,
		node:           node,
		namebaseClient: initClient(ctx),
		mode:           opts.Mode,
	}
	logger.Debug("✅  ExchangeProtocol is Activated \n")

	// Set Peer in Exchange
	peer, err := node.Peer()
	if err != nil {
		logger.Error("Failed to get Profile", err)
		return nil, err
	}
	exchProtocol.Put(peer)
	return exchProtocol, nil
}

// FindPeerId method returns PeerID by SName
func (p *ExchangeProtocol) Get(sname string) (*common.Peer, error) {
	// Get Peer from KadDHT store
	buf, err := p.host.GetValue(p.ctx, sname)
	if err != nil {
		logger.Error("Failed to get item from KadDHT", err)
		return nil, err
	}

	// Unmarshal Peer from buffer
	peerData := &common.Peer{}
	err = proto.Unmarshal(buf, peerData)
	if err != nil {
		return nil, err
	}
	return peerData, err
}

// Put method updates peer instance in the store
func (p *ExchangeProtocol) Put(peer *common.Peer) error {
	// Create a cid manually by specifying the 'prefix' parameters
	key, err := peer.CID()
	if err != nil {
		return err
	}

	// Marshal Peer
	buf, err := peer.Buffer()
	if err != nil {
		logger.Error("Failed to Marshal Peer", err)
		return err
	}

	// Add Peer to KadDHT store
	err = p.host.PutValue(p.ctx, key, buf)
	if err != nil {
		logger.Error("Failed to put Item in KDHT", err)
		return err
	}
	return nil
}

// Verify method uses resolver to check if Peer is registered,
// returns true if Peer is registered
func (p *ExchangeProtocol) Verify(sname string) (bool, host.Record, error) {
	// Check if NamebaseClient is Nil
	empty := host.Record{}
	if p.namebaseClient == nil {
		return false, empty, errors.Wrap(ErrParameters, "NamebaseClient is Nil")
	}

	// Verify Peer is registered
	recs, err := p.host.LookupTXT(p.ctx, sname)
	if err != nil {
		logger.Error("Failed to resolve DNS record for SName", err)
		return false, empty, err
	}

	// Get Name Record
	rec, err := recs.GetNameRecord()
	if err != nil {
		logger.Error("Failed to get Name Record", err)
		return false, empty, err
	}

	// Check peer record
	pubKey, err := rec.PubKey()
	if err != nil {
		logger.Error("Failed to get public key from record", err)
		return false, rec, err
	}

	compId, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		logger.Error("Failed to extract PeerID from PublicKey", err)
		return false, rec, err
	}

	ok, err := compareRecordtoID(rec, compId)
	if err != nil {
		logger.Error("Failed to compare PeerID to record", err)
		return false, rec, err
	}
	return ok, rec, nil
}

// RegisterDomain registers a domain with Namebase.
func (p *ExchangeProtocol) Register(sName string, records ...host.Record) (host.DomainMap, error) {
	// Check if NamebaseClient is Nil
	if p.namebaseClient == nil {
		return nil, errors.Wrap(ErrParameters, "NamebaseClient is Nil")
	}

	// Put records into Namebase
	req := host.NewNBAddRequest(records...)
	ok, err := p.namebaseClient.PutRecords(req)
	if err != nil {
		logger.Error("Failed to Register SName", err)
		return nil, err
	}

	// API Call was Unsuccessful
	if !ok {
		return nil, err
	}

	// Get records from Namebase
	recs, err := p.namebaseClient.FindRecords(sName)
	if err != nil {
		logger.Error("Failed to Find SName after Registering", err)
		return nil, err
	}

	// Map records to DomainMap
	m := make(host.DomainMap)
	for _, r := range recs {
		m[r.Host] = r.Value
	}
	return m, nil
}

// compareRecordtoID compares a record to a PeerID
func compareRecordtoID(r host.Record, target peer.ID) (bool, error) {
	// Check peer record
	pid, err := r.PeerID()
	if err != nil {
		logger.Error("Failed to extract PeerID from PublicKey", err)
		return false, err
	}
	return pid == target, nil
}

// Close method closes the ExchangeProtocol
func (p *ExchangeProtocol) Close() error {
	// Check for isTemporary
	if p.mode.IsTemp() {
		logger.Info("Deleted Temporary SName from DNS Table")
		// Check if NamebaseClient is Nil
		if p.namebaseClient == nil {
			return errors.Wrap(ErrParameters, "NamebaseClient is Nil")
		}

		// Delete SName from Namebase
		req := host.NewNBDeleteRequest(p.authRecord.ToDeleteRecord(), p.nameRecord.ToDeleteRecord())
		_, err := p.namebaseClient.PutRecords(req)
		if err != nil {
			logger.Error("Failed to Delete SName", err)
			return err
		}
	}
	return nil
}
