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
	// Service should be embedded into the implementation of a specific service
	// the specific service should implement myService.Run(ctx)
	// and should return myService.HostService() (blocking) as its final line
	Service struct {
		// UnimplementedServiceServer is embedded to enable forwards compatability
		pb.UnimplementedServiceServer

		// ServiceIdentity holds the identifying information for the Service
		ident.ServiceIdentity

		Server *grpc.Server

		Dominion *dominion.DominionGuard

		Config config.ServiceConfig
	}
)

// NewService builds a service from a ServiceConfig
func NewService(c config.ServiceConfig) (*Service, error) {
	if err := system.Setup(c.LogFile); err != nil {
		return nil, err
	}
	server := grpc.NewServer()

	service := &Service{
		ServiceIdentity: c.ServiceIdentity,
		Server:          server,
		Dominion:        dominion.NewDominionGuard(c.Identity),
		Config:          c,
	}

	pb.RegisterServiceServer(service.Server, service)
	return service, nil
}

func (s Service) GetService() ident.ServiceIdentity {
	return s.ServiceIdentity
}
