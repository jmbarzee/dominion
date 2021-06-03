package microdomain

import (
	"context"
	"errors"

	grpc "github.com/jmbarzee/dominion/grpc"
	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/system"
)

// Heartbeat implements grpc and allows the domain to use grpc.
// Heartbeat serves as the heartbeat from a dominion.
func (d *MicroDomain) Heartbeat(ctx context.Context, request *grpc.HeartbeatRequest) (*grpc.HeartbeatReply, error) {
	rpcName := "Heartbeat"
	system.LogRPCf(rpcName, "Receiving request")

	dominionIdent, err := ident.NewIdentity(request.GetDominion())
	if err != nil {
		system.LogRPCf(rpcName, "Error: %s", err.Error())
		return nil, err
	}

	if err := d.updateDominion(dominionIdent); err != nil {
		system.LogRPCf(rpcName, "Error: %s", err.Error())
		return nil, err
	}

	// Prepare reply
	reply := &grpc.HeartbeatReply{
		Domain:   ident.NewGRPCIdentity(d.Identity),
		Services: ident.NewGRPCServiceIdentityList(d.packageService()),
	}
	system.LogRPCf(rpcName, "Sending reply")
	return reply, nil
}

// StartService implements grpc and initiates a service in the domain.
func (d *MicroDomain) StartService(ctx context.Context, request *grpc.StartServiceRequest) (*grpc.StartServiceReply, error) {
	return nil, errors.New("microDomains cannot start services")
}
