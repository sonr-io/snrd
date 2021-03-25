package node

import (
	"errors"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	dt "github.com/sonr-io/core/pkg/data"
)

// ^ User Node Info ^ //
// @ ID Returns Host ID
func (n *Node) ID() peer.ID {
	return n.host.ID()
}

// ^ Close Ends All Network Communication ^
func (n *Node) Pause() {
	// Check if Response Is Invited
	dt.GetState().Pause()
}

// ^ Close Ends All Network Communication ^
func (n *Node) Resume() {
	dt.GetState().Resume()
}

// ^ Close Ends All Network Communication ^
func (n *Node) Close() {
	n.host.Close()
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) InviteLink(req *md.InviteRequest, p *md.Peer) error {
	// @ 3. Send Invite to Peer
	if n.local.HasPeer(req.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := n.local.FindPeerInTopic(req.To.Id.Peer)
		if err != nil {
			return err
		}

		// Get URL Data
		urlInfo, err := dt.GetPageInfoFromUrl(req.Url)
		if err != nil {
			log.Println(err)
			urlInfo = &md.URLLink{
				Link: req.Url,
			}
		}

		// Build Invite Message
		invMsg := md.AuthInvite{
			IsRemote: req.IsRemote,
			From:     p,
			Payload:  md.Payload_URL,
			Card: &md.TransferCard{
				// SQL Properties
				Payload:  md.Payload_URL,
				Received: int32(time.Now().Unix()),
				Platform: p.Platform,

				// Transfer Properties
				Status: md.TransferCard_DIRECT,

				// Owner Properties
				Username:  p.Profile.Username,
				FirstName: p.Profile.FirstName,
				LastName:  p.Profile.LastName,

				// Data Properties
				Url: urlInfo,
			},
		}
		// Run Routine
		go n.local.Invite(id, &invMsg, p, nil)
	} else {
		return errors.New("Invalid Peer")
	}
	return nil
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) InviteContact(req *md.InviteRequest, p *md.Peer, c *md.Contact) error {
	// @ 3. Send Invite to Peer
	if n.local.HasPeer(req.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := n.local.FindPeerInTopic(req.To.Id.Peer)
		if err != nil {
			return err
		}

		// Build Invite Message
		invMsg := md.AuthInvite{
			IsRemote: req.IsRemote,
			From:     p,
			Payload:  md.Payload_CONTACT,
			Card: &md.TransferCard{
				// SQL Properties
				Payload:  md.Payload_CONTACT,
				Received: int32(time.Now().Unix()),
				Preview:  p.Profile.Picture,
				Platform: p.Platform,

				// Transfer Properties
				Status: md.TransferCard_DIRECT,

				// Owner Properties
				Username:  p.Profile.Username,
				FirstName: p.Profile.FirstName,
				LastName:  p.Profile.LastName,

				// Data Properties
				Contact: c,
			},
		}

		// Run Routine
		go n.local.Invite(id, &invMsg, p, nil)
	} else {
		return errors.New("Invalid Peer")
	}
	return nil
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) InviteFile(card *md.TransferCard, req *md.InviteRequest, p *md.Peer, cf *sf.ProcessedFile) error {
	card.Status = md.TransferCard_INVITE

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:    p,
		Payload: card.Payload,
		Card:    card,
	}

	// Get PeerID
	id, _, err := n.local.FindPeerInTopic(req.To.Id.Peer)
	if err != nil {
		return err
	}

	// Run Routine
	go n.local.Invite(id, &invMsg, p, cf)
	return nil
}

// ^ Respond to an Invitation ^ //
func (n *Node) Respond(decision bool, fs *sf.FileSystem, p *md.Peer, c *md.Contact) {
	n.local.RespondToInvite(decision, fs, p, c)
}

// ^ Send Direct Message to Peer in Lobby ^ //
func (n *Node) Message(msg string, to string, p *md.Peer) error {
	if n.local.HasPeer(to) {
		// Inform Lobby
		if err := n.local.Send(&md.LobbyEvent{
			Event:   md.LobbyEvent_MESSAGE,
			From:    p,
			Id:      p.Id.Peer,
			Message: msg,
			To:      to,
		}); err != nil {
			return err
		}
	}
	return nil
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (n *Node) Update(p *md.Peer) error {
	// Inform Lobby
	if err := n.local.Send(&md.LobbyEvent{
		Event: md.LobbyEvent_UPDATE,
		From:  p,
		Id:    p.Id.Peer,
	}); err != nil {
		return err
	}
	return nil
}
