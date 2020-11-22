package sonr

import (
	"context"
	"errors"
	"fmt"

	"github.com/libp2p/go-libp2p-core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/internal/lobby"
	pb "github.com/sonr-io/core/internal/models"
	st "github.com/sonr-io/core/internal/stream"
)

// ^ CurrentFile returns last file in Processed Files ^ //
func (sn *Node) currentFile() *pb.Metadata {
	return sn.files[len(sn.files)-1]
}

// ^ InitStreams sets Auth/Data Streams with Handlers ^ //
func (sn *Node) initStreams() {
	// Assign Callbacks from Node to Stream
	sn.authStream.Call = st.StreamCallback{
		Invited:   sn.call.OnInvited,
		Responded: sn.call.OnResponded,
		Error:     sn.Error,
	}

	// Set Handlers
	sn.host.SetStreamHandler(protocol.ID("/sonr/auth"), sn.authStream.SetStream)
	//sn.host.SetStreamHandler(protocol.ID("/sonr/transfer"), sn.authStream.HandleTransferStream)
}

// ^ SetDiscovery initializes discovery protocols and creates pubsub service ^ //
func (sn *Node) setDiscovery(ctx context.Context, connEvent *pb.RequestMessage) error {
	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, sn.host)
	if err != nil {
		return err
	}
	fmt.Println("GossipSub Created")

	// Assign Callbacks from Node to Lobby
	lobbyCallbackRef := lobby.LobbyCallback{
		Refreshed: sn.call.OnRefreshed,
		Error:     sn.Error,
	}

	// Enter Lobby
	if sn.lobby, err = lobby.Enter(ctx, lobbyCallbackRef, ps, sn.host.ID(), connEvent.Olc); err != nil {
		return err
	}
	fmt.Println("Lobby Entered")
	return nil
}

// ^ SetUser sets node info from connEvent and host ^ //
func (sn *Node) setPeer(connEvent *pb.RequestMessage) error {
	// Check for Host
	if sn.host == nil {
		err := errors.New("setPeer: Host has not been called")
		return err
	}

	// Set Contact
	sn.Peer = &pb.Peer{
		Id:         sn.host.ID().String(),
		Olc:        connEvent.Olc,
		Device:     connEvent.Device,
		FirstName:  connEvent.Contact.FirstName,
		LastName:   connEvent.Contact.LastName,
		ProfilePic: connEvent.Contact.ProfilePic,
	}

	// Set Profile
	return nil
}
