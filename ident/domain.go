package ident

import (
	"fmt"

	"github.com/jmbarzee/dominion/grpc"
)

// Identity contains the all shareable information about a domain
type DomainIdentity struct {
	// Identity holds identifying information
	Identity

	// Traits is the list of traits of the domain
	Traits []string
}

func (i DomainIdentity) String() string {
	s := `{
	Identity: ` + i.Identity.String() + `
	Traits: [`
	for _, trait := range i.Traits {
		s += trait + ","
	}
	s += `]
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

// NewDomainIdentity creates a DomainIdentity from a grpc.DomainIdentity
func NewDomainIdentity(grpcDIdent *grpc.DomainIdentity) (DomainIdentity, error) {
	identity, err := NewIdentity(grpcDIdent.Identity)
	if err != nil {
		return DomainIdentity{}, fmt.Errorf("error parsing Identity: %w", err)
	}

	return DomainIdentity{
		Identity: identity,
		Traits:   grpcDIdent.GetTraits(),
	}, nil
}

// NewGRPCServiceIdentity creates a grpc.Identity from a Identity
func NewGRPCDomainIdentity(dIdent DomainIdentity) *grpc.DomainIdentity {
	return &grpc.DomainIdentity{
		Identity: NewGRPCIdentity(dIdent.Identity),
		Traits:   dIdent.Traits,
	}
}

// NewDomainIdentityList creates a list of new DomainIdentitys from a list of grpc.DomainIdentity
func NewDomainIdentityList(grpcDIdents []*grpc.DomainIdentity) ([]DomainIdentity, error) {
	dIdents := make([]DomainIdentity, len(grpcDIdents))
	for i, grpcDIdent := range grpcDIdents {
		dIdent, err := NewDomainIdentity(grpcDIdent)
		if err != nil {
			return nil, fmt.Errorf("error parsing DomainIdentity: %w", err)
		}
		dIdents[i] = dIdent
	}
	return dIdents, nil
}

// NewGRPCDomainIdentityList creates a list of new DomainIdentitys from a list of grpc.DomainIdentity
func NewGRPCDomainIdentityList(dIdents []DomainIdentity) []*grpc.DomainIdentity {

	grpcDIdents := make([]*grpc.DomainIdentity, len(dIdents))
	for i, dIdent := range dIdents {
		grpcDIdents[i] = NewGRPCDomainIdentity(dIdent)
	}
	return grpcDIdents
}
