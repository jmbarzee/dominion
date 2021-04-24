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

		// DomainIdentity holds Identifying information for the Domain
		ident.DomainIdentity

		dominion *dominion.DominionGuard

		// services stores the members of a Dominion in a wrapped sync.map as
		//     ServiceType -> Service
		services service.ServiceMap

		stopBroadcastSelf context.CancelFunc
	}
)

// NewDomain creates a new Domain, to correctly build the Domain, just initilize
func NewDomain(c config.DomainConfig) (*Domain, error) {
	if err := system.Setup(c.ID, "domain"); err != nil {
		return nil, err
	}

	return &Domain{
		services:       service.NewServiceMap(),
		DomainIdentity: c.DomainIdentity,
	}, nil
}

func (d Domain) Run(ctx context.Context) error {
	system.Logf("I seek to join the Dominion\n")
	system.Logf(d.DomainIdentity.String())
	system.Logf("The Dominion ever expands!\n")

	// Start Auto Connecting Routines
	go system.RoutineOperation(ctx, "checkIsolation", config.GetDomainConfig().IsolationCheck, d.checkIsolation)
	go system.RoutineOperation(ctx, "checkServices", config.GetDomainConfig().ServiceCheck, d.checkServices)

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

func (d *Domain) updateDominion(ident ident.DominionIdentity) error {
	if d.dominion == nil {
		system.Logf("Joining Dominion %v:", ident.Address.String())
		d.dominion = dominion.NewDominionGuard(ident)
		return nil
	} else {
		return d.dominion.LatchWrite(func(dominion *dominion.Dominion) error {
			if dominion.Address.String() != ident.Address.String() {
				return fmt.Errorf("Dominion Address doesn't known dominion")
			} else {
				dominion.LastContact = time.Now()
				return nil
			}
		})
	}
}
