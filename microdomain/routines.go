package microdomain

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/grandcat/zeroconf"
	pb "github.com/jmbarzee/dominion/grpc"
	"github.com/jmbarzee/dominion/microdomain/dominion"
	"github.com/jmbarzee/dominion/system"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (d *MicroDomain) checkIsolation(ctx context.Context, _ time.Time) {
	var shouldBeBroadcasting bool
	if d.dominion == nil {
		shouldBeBroadcasting = true
	} else {
		d.dominion.LatchRead(func(dominion *dominion.Dominion) error {
			if time.Since(dominion.LastContact) < d.config.IsolationCheck*10 {
				// Dominion hasn't expired
				if d.stopBroadcastSelf != nil {
					d.stopBroadcastSelf()
					d.stopBroadcastSelf = nil
				}
			} else {
				// Dominion has expired
				shouldBeBroadcasting = true
			}
			return nil
		})
	}

	if shouldBeBroadcasting && d.stopBroadcastSelf == nil {
		var ctxBroadcast context.Context
		ctxBroadcast, d.stopBroadcastSelf = context.WithCancel(ctx)
		go d.broadcastSelf(ctxBroadcast)
	}
}

// broadcastSelf uses zero conf to broadcast to a network.
func (d *MicroDomain) broadcastSelf(ctx context.Context) {
	routineName := "broadcastSelf"
	system.LogRoutinef(routineName, "Starting routine")

	// setup broadcasting
	server, err := zeroconf.Register(d.ID.String(), "dominion", "local.", d.Address.Port, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		system.Panic(fmt.Errorf("failed to broadcast: %w", err))
	}
	system.LogRoutinef(routineName, "Started broadcasting .oO \n")

	<-ctx.Done()
	server.Shutdown()
	system.LogRoutinef(routineName, "Stopping routine")
}

func (d *MicroDomain) hostDomain(ctx context.Context) error {
	routineName := "hostDomain"
	system.LogRoutinef(routineName, "Starting routine")

	address := fmt.Sprintf("%s:%v", "", d.Address.Port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("hostDomain() Failed to listen: %w", err)
	}

	server := grpc.NewServer()
	pb.RegisterDomainServer(server, d)
	// Register reflection service on gRPC server.
	go func() {
		<-ctx.Done()
		server.GracefulStop()
		system.LogRoutinef(routineName, "Stopped grpc server gracefully.")
	}()

	reflection.Register(server)
	err = server.Serve(lis)
	system.LogRoutinef(routineName, "Stopping routine")
	return err
}
