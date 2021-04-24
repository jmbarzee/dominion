package service

import (
	pb "github.com/jmbarzee/dominion/grpc"
	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/service/config"
	"github.com/jmbarzee/dominion/service/dominion"
	"github.com/jmbarzee/dominion/system"
	"google.golang.org/grpc"
)

type (
	// Service offers all the shared features of services
	// Service should be emmbeded into the implementation of a specific service
	// the specific service should implement myService.Run(ctx)
	// and should return myService.HostService() (blocking) as its final line
	Service struct {
		// UnimplementedServiceServer is embedded to enable forwards compatability
		pb.UnimplementedServiceServer

		// ServiceIdentity holds the identifying information for the Service
		ident.ServiceIdentity

		Server *grpc.Server

		Dominion *dominion.DominionGuard
	}
)

// NewService builds a service from a ServiceConfig
func NewService(c config.ServiceConfig) (*Service, error) {
	if err := system.Setup(c.ID, c.Type); err != nil {
		return nil, err
	}
	server := grpc.NewServer()

	service := &Service{
		ServiceIdentity: c.ServiceIdentity,
		Server:          server,
		Dominion:        dominion.NewDominionGuard(c.DominionIdentity),
	}

	pb.RegisterServiceServer(service.Server, service)
	return service, nil
}
