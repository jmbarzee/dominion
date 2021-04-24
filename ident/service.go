package ident

import (
	"fmt"

	grpc "github.com/jmbarzee/dominion/grpc"
)

// ServiceIdentity represents a service running under a domain
type ServiceIdentity struct {
	// Identity holds identifying information
	Identity
	// Type is the type of the service
	Type string
}

func (i ServiceIdentity) String() string {
	return `{ Identity: ` + i.Identity.String() + `, Type: ` + i.Type + `}`
}

// NewServiceIdentity creates a ServiceIdentity from a grpc.ServiceIdentity
func NewServiceIdentity(grpcSIdent *grpc.ServiceIdentity) (ServiceIdentity, error) {
	identity, err := NewIdentity(grpcSIdent.Identity)
	if err != nil {
		return ServiceIdentity{}, fmt.Errorf("error parsing Identity: %w", err)
	}
	return ServiceIdentity{
		Identity: identity,
		Type:     grpcSIdent.GetType(),
	}, nil
}

// NewGRPCServiceIdentity creates a grpc.ServiceIdentity from a ServiceIdentity
func NewGRPCServiceIdentity(sIdent ServiceIdentity) *grpc.ServiceIdentity {
	return &grpc.ServiceIdentity{
		Identity: NewGRPCIdentity(sIdent.Identity),
		Type:     sIdent.Type,
	}
}

// NewServiceIdentityList creates a list of new ServiceIdentitys from a list of grpc.ServiceIdentity
func NewServiceIdentityList(grpcSIdents []*grpc.ServiceIdentity) ([]ServiceIdentity, error) {
	sIdents := make([]ServiceIdentity, len(grpcSIdents))
	for i, grpcSIdent := range grpcSIdents {
		sIdent, err := NewServiceIdentity(grpcSIdent)
		if err != nil {
			return nil, fmt.Errorf("error parsing ServiceIdentity: %w", err)
		}
		sIdents[i] = sIdent
	}
	return sIdents, nil
}

// NewGRPCServiceIdentityList creates a list of new ServiceIdentitys from a list of grpc.ServiceIdentity
func NewGRPCServiceIdentityList(sIdents []ServiceIdentity) []*grpc.ServiceIdentity {

	grpcSIdents := make([]*grpc.ServiceIdentity, len(sIdents))
	for i, sIdent := range sIdents {
		grpcSIdents[i] = NewGRPCServiceIdentity(sIdent)
	}
	return grpcSIdents
}
