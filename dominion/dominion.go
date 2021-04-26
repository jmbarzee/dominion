package dominion

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmbarzee/dominion/dominion/config"
	"github.com/jmbarzee/dominion/dominion/domain"
	"github.com/jmbarzee/dominion/grpc"
	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/system"
)

// Dominion is the leader of the Domains
// The Dominion is responsisble for:
//  - Listen for broadcasts from new/lonely Domains
//  - Heartbeats to connected Domains
//  - Command Domains to start new services
type Dominion struct {
	// UnimplementedDominionServer is embedded to enable forwards compatability
	grpc.UnimplementedDominionServer

	// DominionIdentity holds the identifying information for the Dominion
	ident.DominionIdentity

	// domains is a map of domains the dominion currently contains
	domains domain.DomainMap
}

// NewDominion creates a new dominion, to correctly build the dominion, just initilize
func NewDominion(c config.DominionConfig) (*Dominion, error) {

	if err := system.Setup(c.ID, "dominion"); err != nil {
		return nil, err
	}

	return &Dominion{
		domains:          domain.NewDomainMap(),
		DominionIdentity: c.DominionIdentity,
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
	go system.RoutineOperation(ctx, "checkDomains", config.GetDominionConfig().DomainCheck, d.checkDomains)
	go system.RoutineOperation(ctx, "checkServices", config.GetDominionConfig().ServiceCheck, d.checkServices)
	go d.listenForBroadcasts(ctx)

	return d.hostDominion(ctx)
}
func (d *Dominion) packageDomainRecords() []ident.DomainRecord {
	domainRecords := make([]ident.DomainRecord, 0)

	d.domains.Range(func(id uuid.UUID, domainGuard *domain.DomainGuard) bool {
		domainGuard.LatchRead(func(domain domain.Domain) error {
			record := ident.DomainRecord{
				DomainIdentity: domain.DomainIdentity,
				Services:       domain.Services,
			}
			domainRecords = append(domainRecords, record)
			return nil
		})
		return true
	})

	return domainRecords
}

func (d *Dominion) findService(serviceTypeRequested string) []ident.ServiceIdentity {
	serviceIdents := make([]ident.ServiceIdentity, 0)

	d.domains.Range(func(id uuid.UUID, domainGuard *domain.DomainGuard) bool {
		domainGuard.LatchRead(func(domain domain.Domain) error {
			var serviceIdent ident.ServiceIdentity
			ok := false
			for _, sIdent := range domain.Services {
				if sIdent.Type == serviceTypeRequested {
					serviceIdent = sIdent
					ok = true
				}
			}
			if ok {
				serviceIdents = append(serviceIdents, serviceIdent)
			}
			return nil
		})
		return true
	})

	return serviceIdents
}

func (d *Dominion) findServiceCanidates(serviceTypeRequested string) []ident.DomainIdentity {
	traitsNeeded := config.GetServicesConfig().Services[serviceTypeRequested].Traits
	domainIdents := make([]ident.DomainIdentity, 0)

	d.domains.Range(func(id uuid.UUID, domainGuard *domain.DomainGuard) bool {
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
