package dominion

import (
	"context"
	"fmt"

	"github.com/blang/semver"
	"github.com/jmbarzee/dominion/dominion/config"
	"github.com/jmbarzee/dominion/dominion/domain"
	"github.com/jmbarzee/dominion/identity"
	"github.com/jmbarzee/dominion/system"
)

// Dominion is the leader of the Domains
// The Dominion is responsisble for:
//  - Listen for broadcasts from new/lonely Domains
//  - Heartbeats to connected Domains
//  - Command Domains to start new services
type Dominion struct {

	// domains is a map of domains the dominion currently contains
	domains domain.DomainMap

	identity.DominionIdentity
}

// NewDominion creates a new dominion, to correctly build the dominion, just initilize
func NewDominion(config config.DominionConfig) (*Dominion, error) {

	if err := system.Setup("dominion", "dominion"); err != nil {
		return nil, err
	}

	ident, err := NewDominionIdentity(config.Port)
	if err != nil {
		return nil, err
	}

	return &Dominion{
		domains:          domain.NewDomainMap(),
		DominionIdentity: ident,
	}, nil
}

// NewDominionIdentity creates a new DominionIdentity
func NewDominionIdentity(port int) (identity.DominionIdentity, error) {
	// Initialize Version
	version, err := semver.Parse(system.Version)
	if err != nil {
		return identity.DominionIdentity{}, fmt.Errorf("failed to semver.Parse(%v): %v\n", version, err.Error())
	}

	// Initialize IP
	ip, err := system.GetOutboundIP()
	if err != nil {
		return identity.DominionIdentity{}, fmt.Errorf("failed to find Local IP: %v\n", err.Error())
	}

	return identity.DominionIdentity{
		Version: version,
		Address: identity.Address{
			IP:   ip,
			Port: port,
		},
	}, nil
}

// Run begins all the Dominion routines.
// Run doesn't return unless the Dominion has been ended
func (d Dominion) Run(ctx context.Context) error {

	// Intro log
	system.Logf("I am the Dominion\n")
	system.Logf(d.DominionIdentity.String())
	system.Logf("The Dominion ever expands!\n")

	// Start Routines
	go system.RoutineCheck(ctx, "checkDomains", config.GetDominionConfig().DomainCheck, d.checkDomains)
	go system.RoutineCheck(ctx, "checkServices", config.GetDominionConfig().ServiceCheck, d.checkServices)
	go d.listenForBroadcasts(ctx)

	return d.hostDominion(ctx)
}
func (d *Dominion) packageDomains() []identity.DomainIdentity {
	domainIdents := make([]identity.DomainIdentity, 0)

	d.domains.Range(func(uuid string, domainGuard *domain.DomainGuard) bool {
		domainGuard.LatchRead(func(domain domain.Domain) error {
			domainIdents = append(domainIdents, domain.DomainIdentity)
			return nil
		})
		return true
	})

	return domainIdents
}

func (d *Dominion) findService(serviceTypeRequested string) []identity.ServiceIdentity {
	serviceIdents := make([]identity.ServiceIdentity, 0)

	d.domains.Range(func(uuid string, domainGuard *domain.DomainGuard) bool {
		domainGuard.LatchRead(func(domain domain.Domain) error {
			serviceIdent, ok := domain.Services[serviceTypeRequested]
			if ok {
				serviceIdents = append(serviceIdents, serviceIdent)
			}
			return nil
		})
		return true
	})

	return serviceIdents
}

func (d *Dominion) findServiceCanidates(serviceTypeRequested string) []identity.DomainIdentity {
	traitsNeeded := config.GetServicesConfig().Services[serviceTypeRequested].Traits
	domainIdents := make([]identity.DomainIdentity, 0)

	d.domains.Range(func(uuid string, domainGuard *domain.DomainGuard) bool {
		domainGuard.LatchRead(func(domain domain.Domain) error {
			if domain.HasTraits(traitsNeeded) {
				domainIdents = append(domainIdents, domain.DomainIdentity)
			}
			return nil
		})
		return true
	})

	return domainIdents
}
