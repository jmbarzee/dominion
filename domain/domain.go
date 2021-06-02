package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/jmbarzee/dominion/domain/config"
	"github.com/jmbarzee/dominion/domain/dominion"
	"github.com/jmbarzee/dominion/domain/service"
	grpc "github.com/jmbarzee/dominion/grpc"
	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/system"
)

type (
	Domain struct {
		// UnimplementedDomainServer is embedded to enable forwards compatability
		grpc.UnimplementedDomainServer

		// Identity holds Identifying information for the Domain
		ident.Identity

		dominion *dominion.DominionGuard

		// services stores the members of a Dominion in a wrapped sync.map as
		//     ServiceType -> Service
		services service.ServiceMap

		// stopBroadcastSelf is a handle to end network broadcasting
		stopBroadcastSelf context.CancelFunc

		// config is the initial configuration for the Dominion
		config config.DomainConfig
	}
)

// NewDomain creates a new Domain, to correctly build the Domain, just initialize
func NewDomain(c config.DomainConfig) (*Domain, error) {
	if err := system.Setup(c.LogFile); err != nil {
		return nil, err
	}

	return &Domain{
		services: service.NewServiceMap(),
		Identity: c.Identity,
		config:   c,
	}, nil
}

func (d Domain) Run(ctx context.Context) error {
	system.Logf("I seek to join the Dominion\n")
	system.Logf(d.Identity.String())
	system.Logf("The Dominion ever expands!\n")

	// Start Auto Connecting Routines
	go system.RoutineOperation(ctx, "checkIsolation", d.config.IsolationCheck, d.checkIsolation)
	go system.RoutineOperation(ctx, "checkServices", d.config.ServiceCheck, d.checkServices)

	return d.hostDomain(ctx)
}

func (d Domain) packageServices() []ident.ServiceIdentity {
	services := []ident.ServiceIdentity{}
	d.services.Range(func(serviceType string, serviceGuard *service.ServiceGuard) bool {
		serviceGuard.LatchRead(func(service service.Service) error {
			services = append(services, service.ServiceIdentity)
			return nil
		})
		return true
	})
	return services
}

func (d *Domain) updateDominion(ident ident.Identity) error {
	if d.dominion == nil {
		system.Logf("Joining Dominion %v:", ident.Address.String())
		d.dominion = dominion.NewDominionGuard(ident)
		return nil
	} else {
		return d.dominion.LatchWrite(func(dominion *dominion.Dominion) error {
			if dominion.Address.String() != ident.Address.String() {
				return fmt.Errorf("dominion Address doesn't known dominion")
			} else {
				dominion.LastContact = time.Now()
				return nil
			}
		})
	}
}
