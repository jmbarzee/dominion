package domain

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/jmbarzee/dominion/domain/dominion"
	service "github.com/jmbarzee/dominion/domain/service"
	grpc "github.com/jmbarzee/dominion/grpc"
	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/system"
	"github.com/jmbarzee/dominion/system/connect"
)

// Heartbeat implements grpc and allows the domain to use grpc.
// Heartbeat serves as the heartbeat from a dominion.
func (d *Domain) Heartbeat(ctx context.Context, request *grpc.HeartbeatRequest) (*grpc.HeartbeatReply, error) {
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
		Services: ident.NewGRPCServiceIdentityList(d.packageServices()),
	}
	system.LogRPCf(rpcName, "Sending reply")
	return reply, nil
}

func (d *Domain) rpcHeartbeat(ctx context.Context, serviceGuard *service.ServiceGuard) {
	rpcName := "Heartbeat"
	serviceType := ""
	err := serviceGuard.LatchWrite(func(service *service.Service) error {
		serviceType = service.Type
		cc := connect.NewConnectionConfig(d.config.DialTimeout)

		if err := connect.CheckConnection(ctx, service, cc); err != nil {
			return fmt.Errorf("Failed to check connection: %w", err)
		}

		// Prepare request
		request := &grpc.ServiceHeartbeatRequest{
			Domain: ident.NewGRPCIdentity(d.Identity),
		}

		// Send RPC
		system.LogRPCf(rpcName, "Sending request")
		client := grpc.NewServiceClient(service.Conn)
		reply, err := client.Heartbeat(ctx, request)
		if err != nil {
			return err
		}
		system.LogRPCf(rpcName, "Received reply")

		// Update domain
		service.LastContact = time.Now()
		serviceIdent, err := ident.NewServiceIdentity(reply.GetService())
		if err != nil {
			return err
		}
		service.ServiceIdentity = serviceIdent
		return nil
	})

	if err != nil {
		system.Logf("Failed to heartbeat \"%v\": %v: Dropping service", serviceType, err.Error())
		d.services.Delete(serviceType)
	}
}

// StartService implements grpc and initiates a service in the domain.
func (d *Domain) StartService(ctx context.Context, request *grpc.StartServiceRequest) (*grpc.StartServiceReply, error) {
	rpcName := "StartService"
	system.LogRPCf(rpcName, "Receiving request")
	serviceIdent, err := d.startService(request.GetType(), request.GetDockerImage())
	if err != nil {
		err := fmt.Errorf("Failed to start service: %w", err)
		system.Errorf("Error starting service: %w", err)
		return nil, err
	}

	reply := &grpc.StartServiceReply{
		Service: ident.NewGRPCServiceIdentity(serviceIdent),
	}

	system.LogRPCf(rpcName, "Sending reply")
	return reply, nil
}

func (d *Domain) startService(serviceType, dockerImage string) (ident.ServiceIdentity, error) {
	if _, ok := d.services.Load(serviceType); ok {
		return ident.ServiceIdentity{}, fmt.Errorf("Service already exists! (%s)", serviceType)
	}

	var dominionIdent ident.Identity
	d.dominion.LatchRead(func(dominion *dominion.Dominion) error {
		dominionIdent = dominion.Identity
		return nil
	})

	domainIdent := d.Identity

	serviceIdent := ident.ServiceIdentity{
		Identity: ident.Identity{
			ID:      domainIdent.ID, // TODO: IDs need to be unique and generated idempotently
			Version: domainIdent.Version,
			Address: ident.Address{
				IP:   domainIdent.Address.IP,
				Port: getRandomPort(d.Address.Port),
			},
		},
		Type: serviceType,
	}

	err := service.Start(serviceIdent, domainIdent, dominionIdent, dockerImage)
	if err != nil {
		return ident.ServiceIdentity{}, err
	}

	// Give the service a little bit of time to start
	// TODO: jmbarzee is there a better solution?
	time.Sleep(d.config.ServiceCheck * 3)

	d.services.Store(serviceType, service.NewServiceGuard(serviceIdent))
	system.Logf("Started service: %s", serviceIdent.String())
	return serviceIdent, nil
}

// getRandomPort returns a random port based on a seed
func getRandomPort(seedPort int) int {
	uint16Max := (1 << 16) - 1
	return (rand.Intn(uint16Max) + seedPort) % uint16Max
}
