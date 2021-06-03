package microdomain

import (
	"context"
	"fmt"
	"time"

	grpc "github.com/jmbarzee/dominion/grpc"
	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/microdomain/config"
	"github.com/jmbarzee/dominion/microdomain/dominion"
	"github.com/jmbarzee/dominion/system"
)

type (
	MicroDomain struct {
		// UnimplementedDomainServer is embedded to enable forwards compatability
		grpc.UnimplementedDomainServer

		// Identity holds Identifying information for the Domain
		ident.Identity

		dominion *dominion.DominionGuard

		service Service

		// stopBroadcastSelf is a handle to end network broadcasting
		stopBroadcastSelf context.CancelFunc

		// config is the initial configuration for the Dominion
		config config.DomainConfig
	}

	Service interface {
		Run(ctx context.Context)
		GetService() ident.ServiceIdentity
	}
)

// NewMicroDomain creates a new Domain, to correctly build the Domain, just initialize
func NewMicroDomain(c config.DomainConfig, service Service) (*MicroDomain, error) {
	if err := system.Setup(c.LogFile); err != nil {
		return nil, err
	}

	return &MicroDomain{
		Identity: c.Identity,
		config:   c,
		service:  service,
	}, nil
}

func (d MicroDomain) Run(ctx context.Context) error {
	system.Logf("I seek to join the Dominion\n")
	system.Logf(d.Identity.String())
	system.Logf("The Dominion ever expands!\n")

	// Start Auto Connecting Routines
	go system.RoutineOperation(ctx, "checkIsolation", d.config.IsolationCheck, d.checkIsolation)
	go d.service.Run(ctx)

	return d.hostDomain(ctx)
}

func (d MicroDomain) packageService() []ident.ServiceIdentity {
	services := []ident.ServiceIdentity{
		d.service.GetService(),
	}
	return services
}

func (d *MicroDomain) updateDominion(ident ident.Identity) error {
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
