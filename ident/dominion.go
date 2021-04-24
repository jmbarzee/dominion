package ident

import (
	"fmt"

	"github.com/jmbarzee/dominion/grpc"
)

type (
	// DominionIdentity represents a dominion
	DominionIdentity struct {
		// Identity holds identifying information
		Identity
	}
)

func (i DominionIdentity) String() string {
	return `{
	Identity: ` + i.Identity.String() + `
}`
}

// NewDominionIdentity creates a DominionIdentity from a grpc.DominionIdentity
func NewDominionIdentity(grpcDIdent *grpc.DominionIdentity) (DominionIdentity, error) {
	identity, err := NewIdentity(grpcDIdent.Identity)
	if err != nil {
		return DominionIdentity{}, fmt.Errorf("error parsing Identity: %w", err)
	}

	return DominionIdentity{
		Identity: identity,
	}, nil
}

// NewPBDominionIdentity creates a grpc.DominionIdentity from a DominionIdentity
func NewGRPCDominionIdentity(dIdent DominionIdentity) *grpc.DominionIdentity {
	return &grpc.DominionIdentity{
		Identity: NewGRPCIdentity(dIdent.Identity),
	}
}
