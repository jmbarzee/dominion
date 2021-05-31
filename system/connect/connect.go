package connect

import (
	"context"
	"fmt"
	"time"

	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/system"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// Connectable is anything which holds a gRPC connection
type Connectable interface {
	// GetAddress returns the target address for the connection
	GetAddress() ident.Address
	// GetConnection returns a current gRCP connection (for checking)
	GetConnection() *grpc.ClientConn
	// SetConnection replaces the connection of the device (any existing connection will be closed prior to this)
	SetConnection(*grpc.ClientConn)
}

// ConnectionConfig configures the opperation of the connection operations
type ConnectionConfig interface {
	// GetTimeout returns the timeout for dialing a new connection
	GetTimeout() time.Duration
}

// CheckConnection checks the gRPC connection of a connectable
func CheckConnection(ctx context.Context, c Connectable, cc ConnectionConfig) error {
	if c.GetConnection() == nil {
		if err := reconnect(ctx, c, cc); err != nil {
			return err
		}
	}

	failures := 0

	conn := c.GetConnection()

	for {
		state := conn.GetState()
		switch state {
		case connectivity.Idle:
			// connection is waiting for rpcs
			return nil
		case connectivity.Connecting:
			conn.WaitForStateChange(ctx, connectivity.Connecting)
			break
		case connectivity.Ready:
			return nil
		case connectivity.TransientFailure:
			conn.WaitForStateChange(ctx, connectivity.TransientFailure)
			failures++
			break
		case connectivity.Shutdown:
			return reconnect(ctx, c, cc)
		default:
			return fmt.Errorf("connection has un recognized state:%v", state)
		}
		if failures >= 3 {
			system.Logf("too many connection failures %v, reconnecting %v", failures, c.GetAddress())
			return reconnect(ctx, c, cc)
		}
	}
}

func reconnect(ctx context.Context, c Connectable, cc ConnectionConfig) error {
	if c.GetConnection() != nil {
		err := c.GetConnection().Close()
		if err != nil {
			return fmt.Errorf("failed to close connection during reconnect: %w", err)
		}
	}

	subCtx, cancel := context.WithTimeout(ctx, cc.GetTimeout())
	defer cancel()

	// connect
	conn, err := grpc.DialContext(
		subCtx,
		c.GetAddress().String(),
		grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("Failed to dial connection during reconnect: %w", err)
	}

	// update service
	c.SetConnection(conn)
	return nil
}
