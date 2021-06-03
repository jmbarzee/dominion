package service

import (
	"time"

	"github.com/jmbarzee/dominion/ident"
	"google.golang.org/grpc"
)

// Service is a representation of a service on the same machine
// Service implements system.Connectable
type Service struct {
	//ServiceIdentity holds the identifying information of the service
	ident.ServiceIdentity

	// conn is the protocol buffer connection to the member
	Conn *grpc.ClientConn

	// LastContact is the last time a service replied to a rpc
	LastContact time.Time
}

// GetAddress returns the target address for the connection
func (s Service) GetAddress() ident.Address {
	return s.Identity.Address
}

// GetConnection returns a current gRCP connection (for checking)
func (s Service) GetConnection() *grpc.ClientConn {
	return s.Conn
}

// SetConnection replaces the connection of the device (any existing connection will be closed prior to this)
func (s *Service) SetConnection(newConn *grpc.ClientConn) {
	s.Conn = newConn
}
