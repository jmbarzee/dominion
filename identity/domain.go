package identity

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/google/uuid"
	pb "github.com/jmbarzee/dominion/grpc"
	"github.com/jmbarzee/dominion/system"
)

// Identity contains the all shareable information about a domain
type DomainIdentity struct {
	// Address is the network address of the domain
	Address Address
	// Version is the version of code being run by the domain
	Version semver.Version
	// ID is a unique identifier for a Domain
	ID uuid.UUID
	// Traits is the list of traits of the domain
	Traits []string

	// Services is the list of services the Domain currently offers
	Services map[string]ServiceIdentity
}

func (i DomainIdentity) String() string {
	s := `{
	Address: ` + i.Address.String() + `
	Version: ` + i.Version.String() + `
	ID: ` + i.ID.String() + `
	Traits: [`
	for _, trait := range i.Traits {
		s += trait + ","
	}
	s += `]
	Services: {
`
	for serviceType, service := range i.Services {
		s += `		` + serviceType + ": " + service.String() + "\n"
	}
	s += `	}
}`
	return s
}

func (d DomainIdentity) HasTraits(traits []string) bool {
	hasTraits := true
	for _, trait := range traits {
		if !d.HasTrait(trait) {
			hasTraits = false
			break
		}
	}
	return hasTraits
}

func (d DomainIdentity) HasTrait(trait string) bool {
	for _, ownTrait := range d.Traits {
		if ownTrait == trait {
			return true
		}
	}
	return false
}

// NewDomainIdentity creates a DomainIdentity from a pb.DomainIdentity
func NewDomainIdentity(pbdIdent *pb.DomainIdentity) DomainIdentity {

	version, err := semver.Parse(pbdIdent.GetVersion())
	if err != nil {
		system.Panic(fmt.Errorf("error parseing version from \"%v\": %w", pbdIdent.GetVersion(), err))
	}

	id, err := uuid.FromBytes(pbdIdent.GetID())
	if err != nil {
		system.Panic(fmt.Errorf("error parseing id from \"%b\": %w", pbdIdent.GetID(), err))
	}

	return DomainIdentity{
		Address:  NewAddress(pbdIdent.GetAddress()),
		Version:  version,
		ID:       id,
		Traits:   pbdIdent.GetTraits(),
		Services: NewServiceIdentityMap(pbdIdent.GetServices()),
	}
}

// NewPBServiceIdentity creates a pb.Identity from a Identity
func NewPBDomainIdentity(dIdent DomainIdentity) *pb.DomainIdentity {
	return &pb.DomainIdentity{
		Address:  NewPBAddress(dIdent.Address),
		ID:       dIdent.ID[:],
		Version:  dIdent.Version.String(),
		Traits:   dIdent.Traits,
		Services: NewPBServiceIdentityMap(dIdent.Services),
	}
}

// NewDomainIdentityList creates a list of new DomainIdentitys from a list of pb.DomainIdentity
func NewDomainIdentityList(pbdIdents []*pb.DomainIdentity) []DomainIdentity {
	dIdents := make([]DomainIdentity, len(pbdIdents))
	for i, pbdIdent := range pbdIdents {
		dIdents[i] = NewDomainIdentity(pbdIdent)
	}
	return dIdents
}

// NewPBDomainIdentityList creates a list of new DomainIdentitys from a list of pb.DomainIdentity
func NewPBDomainIdentityList(dIdents []DomainIdentity) []*pb.DomainIdentity {

	pbdIdents := make([]*pb.DomainIdentity, len(dIdents))
	for i, dIdent := range dIdents {
		pbdIdents[i] = NewPBDomainIdentity(dIdent)
	}
	return pbdIdents
}
