package dominion

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/dominion/dominion/domain"
	grpc "github.com/jmbarzee/dominion/grpc"
	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/system"
	"github.com/jmbarzee/dominion/system/connect"
)

// GetServices implements grpc and serves as the directory of services hosted on all domains.
// GetServices is called by services hosted on a single domain to find their dependencies.
func (d *Dominion) GetServices(ctx context.Context, request *grpc.GetServicesRequest) (*grpc.GetServicesReply, error) {
	rpcName := "GetServices"
	system.LogRPCf(rpcName, "Receiving request")
	reply := &grpc.GetServicesReply{
		Services: ident.NewGRPCServiceIdentityList(d.findService(request.Type)),
	}
	system.LogRPCf(rpcName, "Sending reply")
	return reply, nil
}

// GetDomains implements grpc and returns all domains and their services
func (d *Dominion) GetDomains(ctx context.Context, request *grpc.Empty) (*grpc.GetDomainsReply, error) {
	rpcName := "GetDomains"
	system.LogRPCf(rpcName, "Receiving request")

	reply := &grpc.GetDomainsReply{
		DomainRecords: ident.NewGRPCDomainRecordList(d.packageDomainRecords()),
	}
	system.LogRPCf(rpcName, "Sending reply")
	return reply, nil
}

func (d *Dominion) rpcHeartbeat(ctx context.Context, domainGuard *domain.DomainGuard) {
	rpcName := "Heartbeat"
	id := uuid.Nil
	err := domainGuard.LatchWrite(func(domain *domain.Domain) error {
		id = domain.Identity.ID
		cc := connect.NewConnectionConfig(d.config.DialTimeout)

		if err := connect.CheckConnection(ctx, domain, cc); err != nil {
			return fmt.Errorf("failed to check connection: %w", err)
		}

		// Prepare request
		request := &grpc.HeartbeatRequest{
			Dominion: ident.NewGRPCIdentity(d.Identity),
		}

		// Send RPC
		system.LogRPCf(rpcName, "Sending request")
		client := grpc.NewDomainClient(domain.Conn)
		reply, err := client.Heartbeat(ctx, request)
		if err != nil {
			return err
		}
		system.LogRPCf(rpcName, "Received reply")

		// Update domain
		domain.LastContact = time.Now()
		domainIdent, err := ident.NewIdentity(reply.GetDomain())
		if err != nil {
			return err
		}
		domain.Identity = domainIdent
		return nil
	})

	if err != nil {
		system.Errorf("Dropping Domain %s: %w", id, err)
		d.domains.Delete(id)
	}
}

func (d *Dominion) rpcStartService(ctx context.Context, domainGuard *domain.DomainGuard, serviceType string) error {
	rpcName := "StartService"
	return domainGuard.LatchWrite(func(domain *domain.Domain) error {
		cc := connect.NewConnectionConfig(d.config.DialTimeout)

		if err := connect.CheckConnection(ctx, domain, cc); err != nil {
			return fmt.Errorf("failed to check connection: %w", err)
		}

		// Prepare request
		request := &grpc.StartServiceRequest{
			Type: serviceType,
		}

		// Send RPC
		system.LogRPCf(rpcName, "Sending request")
		client := grpc.NewDomainClient(domain.Conn)
		reply, err := client.StartService(ctx, request)
		if err != nil {
			return err
		}
		system.LogRPCf(rpcName, "Received reply")

		// Update domain
		domain.LastContact = time.Now()
		serviceIdent, err := ident.NewServiceIdentity(reply.GetService())
		if err != nil {
			return err
		}

		domain.AddService(serviceIdent)
		return nil

	})
}
