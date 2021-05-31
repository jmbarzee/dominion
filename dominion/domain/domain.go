package domain

import (
	"fmt"
	"time"

	"github.com/jmbarzee/dominion/ident"
	"google.golang.org/grpc"
)

// Domain is a representation of a domain service somewhere on the network
// Domain implements system.Connectable
type Domain struct {
	// DomainIdentity holds the identifying information of the domain
	ident.DomainIdentity

	// Services
	Services []ident.ServiceIdentity

	// conn is the protocol buffer connection to the member
	Conn *grpc.ClientConn

	// LastContact is the last time a domain replied to a rpc
	LastContact time.Time
}

// GetAddress returns the target address for the connection
func (d Domain) GetAddress() ident.Address {
	return d.Identity.Address
}

// GetConnection returns a current gRCP connection (for checking)
func (d Domain) GetConnection() *grpc.ClientConn {
	return d.Conn
}

// SetConnection replaces the connection of the device (any existing connection will be closed prior to this)
func (d *Domain) SetConnection(newConn *grpc.ClientConn) {
	d.Conn = newConn
}

// AddService adds a serviceIdentity if that service doesn't share an ID with any other service controled by the domain
func (d Domain) AddService(serviceIdent ident.ServiceIdentity) error {
	for _, sIdent := range d.Services {
		if sIdent.ID == serviceIdent.ID {
			return fmt.Errorf("could not add service, found matching ID '%s'", serviceIdent.ID.String())
		}
	}
	d.Services = append(d.Services, serviceIdent)
	return nil
}
