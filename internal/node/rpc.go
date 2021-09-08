package node

import (
	context "context"
	"fmt"
	"net"

	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
)

const RPC_SERVER_PORT = 60214

type NodeRPCService struct {
	NodeServiceServer
	*Node

	// Properties
	grpcServer *grpc.Server
	listener   net.Listener

	// Channels
	decisionEvents chan *common.DecisionEvent
	inviteEvents   chan *common.InviteEvent
	progressEvents chan *common.ProgressEvent
	completeEvents chan *common.CompleteEvent
}

// NewRPCService creates a new RPC service for the node.
func NewRPCService(n *Node) (*NodeRPCService, error) {
	// Bind RPC Service
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", RPC_SERVER_PORT))
	if err != nil {
		logger.Error("Failed to bind to port", zap.Error(err))
		return nil, err
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	nrc := &NodeRPCService{
		Node:           n,
		decisionEvents: make(chan *common.DecisionEvent),
		inviteEvents:   make(chan *common.InviteEvent),
		progressEvents: make(chan *common.ProgressEvent),
		completeEvents: make(chan *common.CompleteEvent),
		grpcServer:     grpcServer,
		listener:       listener,
	}

	// Register RPC Service
	RegisterNodeServiceServer(grpcServer, nrc)
	if err := grpcServer.Serve(listener); err != nil {
		logger.Error("Failed to Register node service server", zap.Error(err))
		return nil, err
	}
	return nrc, nil
}

// Supply supplies the node with the given amount of resources.
func (n *NodeRPCService) Supply(ctx context.Context, req *SupplyRequest) (*SupplyResponse, error) {
	// Call Internal Supply
	err := n.Node.Supply(req.GetPaths())
	if err != nil {
		return &SupplyResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	// Send Response
	return &SupplyResponse{
		Success: true,
	}, nil
}

// Edit method edits the node's user profile.
func (n *NodeRPCService) Edit(ctx context.Context, req *EditRequest) (*EditResponse, error) {
	return nil, nil
}

// Share method sends supplied files/urls with a peer
func (n *NodeRPCService) Share(ctx context.Context, req *ShareRequest) (*ShareResponse, error) {
	return nil, nil
}

// Respond method responds to a received InviteRequest.
func (n *NodeRPCService) Respond(ctx context.Context, req *RespondRequest) (*RespondResponse, error) {
	// // Unmarshal Data to Request
	// resp := &data.DecisionRequest{}
	// if err := proto.Unmarshal(buf, resp); err != nil {
	// 	n.handleError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
	// 	return
	// }

	// // Send Response
	// n.client.Respond(resp.ToResponse())
	return nil, nil
}

// OnDecision method sends a decision event to the client.
func (n *NodeRPCService) OnDecision(e *Empty, stream NodeService_OnDecisionServer) error {
	for {
		select {
		case m := <-n.decisionEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}

// OnInvite method sends an invite event to the client.
func (n *NodeRPCService) OnInvite(e *Empty, stream NodeService_OnInviteServer) error {
	for {
		select {
		case m := <-n.inviteEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}

// OnProgress method sends a progress event to the client.
func (n *NodeRPCService) OnProgress(e *Empty, stream NodeService_OnProgressServer) error {
	for {
		select {
		case m := <-n.progressEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}

// OnComplete method sends a complete event to the client.
func (n *NodeRPCService) OnComplete(e *Empty, stream NodeService_OnCompleteServer) error {
	for {
		select {
		case m := <-n.completeEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}
