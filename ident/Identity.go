package ident

import (
	"fmt"
	"net"

	"github.com/blang/semver"
	"github.com/google/uuid"
	"github.com/jmbarzee/dominion/grpc"
)

// Identity is a set of identifying information
type Identity struct {
	// ID is a unique identifier
	ID uuid.UUID
	// Version is the version of code being run
	Version semver.Version
	// Address is a network address
	Address Address
}

// NewIdentity creates a Identity from a grcp.Identity
func NewIdentity(grpcIdent *grpc.Identity) (Identity, error) {
	id, err := uuid.FromBytes(grpcIdent.GetID())
	if id == uuid.Nil {
		return Identity{}, fmt.Errorf("error parsing ID: %w", err)
	}

	version, err := semver.Parse(grpcIdent.GetVersion())
	if err != nil {
		return Identity{}, fmt.Errorf("error parsing Version: %w", err)
	}

	return Identity{
		ID:      id,
		Version: version,
		Address: NewAddress(grpcIdent.GetAddress()),
	}, nil
}

// NewGRPCIdentity creates a grcp.ServiceIdentity from a Identity
func NewGRPCIdentity(ident Identity) *grpc.Identity {
	return &grpc.Identity{
		ID:      ident.ID[:],
		Version: ident.Version.String(),
		Address: NewGRPCAddress(ident.Address),
	}
}

func (i Identity) String() string {
	return fmt.Sprintf("{ %s, %s, %s}", i.ID.String(), i.Version.String(), i.Address.String())
}

// Address is a IP and port combination
type Address struct {
	// IP is the ip which the Domain will be responding on
	IP net.IP
	// Port is the port which the Domain will be responding on
	Port int
}

// NewAddress creates a Address from a grcp.Address
func NewAddress(grpcAddr *grpc.Address) Address {
	return Address{
		IP:   grpcAddr.GetIP(),
		Port: int(grpcAddr.GetPort()),
	}
}

// NewGRPCAddress creates a grcp.ServiceIdentity from a Address
func NewGRPCAddress(addr Address) *grpc.Address {
	return &grpc.Address{
		IP:   addr.IP,
		Port: int32(addr.Port),
	}
}

func (a Address) String() string {
	return fmt.Sprintf("%s:%v", a.IP.String(), a.Port)
}
