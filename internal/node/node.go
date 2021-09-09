package node

import (
	"container/list"
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/transfer"
	"github.com/sonr-io/core/tools/emitter"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// Node Emission Events
const (
	Event_STATUS = "status"
)

// Node type - a p2p host implementing one or more p2p protocols
type Node struct {
	// Emitter is the event emitter for this node
	*emitter.Emitter

	// StateMachine state
	*state.StateMachine

	// Host and context
	*host.SHost

	// Properties
	ctx context.Context

	// Queue - the transfer queue
	queue *list.List

	// Profile - the node's profile
	profile *common.Profile

	// TransferProtocol - the transfer protocol
	*transfer.TransferProtocol

	// ExchangeProtocol - the exchange protocol
	*exchange.ExchangeProtocol
}

// NewNode Creates a node with its implemented protocols
func NewNode(ctx context.Context, host *host.SHost, loc *common.Location) *Node {
	// Initialize Node
	node := &Node{
		Emitter: emitter.New(2048),
		SHost:   host,
		ctx:     ctx,
		queue:   list.New(),
	}

	// Create Transfer Protocol
	node.TransferProtocol = transfer.NewProtocol(host, node.Emitter)
	node.Emit(Event_STATUS, true, "Transfer Protocol Set")

	// Set Exchange Protocol
	exch, err := exchange.NewProtocol(host, loc, node.Emitter)
	if err != nil {
		logger.Error("Failed to start ExchangeProtocol", zap.Error(err))
		return node
	}

	node.ExchangeProtocol = exch
	return node
}

// Edit method updates Node's profile
func (n *Node) Edit(p *common.Profile) error {
	buf, err := proto.Marshal(p)
	if err != nil {
		logger.Error("Failed to marshal Profile", zap.Error(err))
		return err
	}

	err = n.ExchangeProtocol.Update(p.GetSName(), buf)
	if err != nil {
		logger.Error("Failed to edit Profile", zap.Error(err))
		return err
	}
	return nil
}

// Supply a transfer item to the queue
func (n *Node) Supply(paths []string) error {
	// Create Transfer
	tr := common.Transfer{
		Metadata: n.NewMetadata(),
	}

	// Initialize Transfer Items and add iterate over paths
	items := make([]*common.Transfer_Item, len(paths))
	for _, path := range paths {
		// Check if path is a url
		if common.IsUrl(path) {
			items = append(items, common.NewTransferUrlItem(path))
		} else {
			// Create File Item
			item, err := common.NewTransferFileItem(path)
			if err != nil {
				logger.Error("Failed to edit Profile", zap.Error(err))
				n.Emit(Event_STATUS, err)
				return err
			}

			// Add item to transfer
			items = append(items, item)
		}
	}

	// Add items to transfer
	tr.Items = items
	n.queue.PushBack(&tr)
	return nil
}

// Invite a peer to have a transfer
func (n *Node) Invite(id peer.ID) error {
	// Create Invite Request
	req := &transfer.InviteRequest{
		Transfer: n.queue.Front().Value.(*common.Transfer),
		Metadata: n.NewMetadata(),
	}

	// Invite peer
	err := n.TransferProtocol.Invite(id, req)
	if err != nil {
		logger.Error("Failed to invite peer", zap.Error(err))
		n.Emit(Event_STATUS, err)
		return err
	}
	return nil
}

// Respond to an invite request
func (n *Node) Respond(decs bool) error {
	// Create Invite Response

	// n.TransferProtocol.Respond(id)
	return nil
}
