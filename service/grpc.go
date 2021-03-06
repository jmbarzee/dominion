package service

import (
	"context"

	pb "github.com/jmbarzee/dominion/grpc"
	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/service/config"
	"github.com/jmbarzee/dominion/service/dominion"
	"github.com/jmbarzee/dominion/system"
	"github.com/jmbarzee/dominion/system/connect"
)

// Heartbeat implements grpc and allows the domain to use grpc.
// Heartbeat serves as the heartbeat from a dominion.
func (s *Service) Heartbeat(ctx context.Context, request *pb.ServiceHeartbeatRequest) (*pb.ServiceHeartbeatReply, error) {
	rpcName := "Heartbeat"
	system.LogRPCf(rpcName, "Receiving request")

	// Prepare reply
	reply := &pb.ServiceHeartbeatReply{
		Service: ident.NewGRPCServiceIdentity(s.ServiceIdentity),
	}
	system.LogRPCf(rpcName, "Sending reply")
	return reply, nil
}

// RPCGetServices requests a list of services from the dominion
func (s Service) RPCGetServices(ctx context.Context, serviceType string) ([]ident.ServiceIdentity, error) {
	rpcName := "GetServices"
	services := []ident.ServiceIdentity{}

	err := s.Dominion.LatchWrite(func(dominion *dominion.Dominion) error {

		cc := connect.NewConnectionConfig(config.DefaultServiceDialTimeout)

		if err := connect.CheckConnection(ctx, dominion, cc); err != nil {
			return err
		}

		serviceRequest := &pb.GetServicesRequest{
			Type: serviceType,
		}

		system.LogRPCf(rpcName, "Sending request")
		dominionClient := pb.NewDominionClient(dominion.Conn)
		reply, err := dominionClient.GetServices(ctx, serviceRequest)
		if err != nil {
			return err
		}
		system.LogRPCf(rpcName, "Received reply")

		serviceList, err := ident.NewServiceIdentityList(reply.GetServices())
		if err != nil {
			return err
		}
		services = serviceList
		return nil
	})
	if err != nil {
		return nil, err
	}

	return services, nil
}
