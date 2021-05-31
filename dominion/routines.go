package dominion

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/grandcat/zeroconf"
	"github.com/jmbarzee/dominion/dominion/config"
	"github.com/jmbarzee/dominion/dominion/domain"
	pb "github.com/jmbarzee/dominion/grpc"
	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/system"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// listenForBroadcasts listens for other Legionnaires.
func (d *Dominion) listenForBroadcasts(ctx context.Context) {
	routineName := "listenForBroadcasts"
	system.LogRoutinef(routineName, "Starting routine")

	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		system.Panic(fmt.Errorf("failed to initialize resolver: %w", err))
	}

	entries := make(chan *zeroconf.ServiceEntry)
	err = resolver.Browse(ctx, "dominion", "local.", entries)
	if err != nil {
		system.Panic(fmt.Errorf("failed to browse: %w", err))
	}

	system.LogRoutinef(routineName, "Listening for broadcasts...")

Loop:
	for {
		select {
		case entry, ok := <-entries:
			if !ok {
				// channel closed
				break Loop
			}
			if len(entry.AddrIPv4) <= 0 {
				break
			}

			id, err := uuid.Parse(entry.Instance)
			if err != nil {
				system.LogRoutinef(routineName, "Failed to parse uuid: %v", err.Error())
				break
			}
			ip := entry.AddrIPv4[0]
			port := entry.Port
			system.LogRoutinef(routineName, "Found broadcast - uuid:%v ip:%v port:%v", id, ip, port)

			// Write a temp ident. Should be populated on first heartbeat
			newDomainGuard := domain.NewDomainGuard(ident.DomainIdentity{
				Identity: ident.Identity{
					Address: ident.Address{
						IP:   ip,
						Port: port,
					},
					ID: id,
				},
			})

			// Add the new member
			d.domains.Store(id, newDomainGuard)

		case <-ctx.Done():
			break Loop
		}
	}

	system.LogRoutinef(routineName, "Stopping routine")
}

func (d *Dominion) checkDomains(ctx context.Context, _ time.Time) {
	d.domains.Range(func(id uuid.UUID, domainGuard *domain.DomainGuard) bool {
		domainGuard.LatchRead(func(domain domain.Domain) error {
			if time.Since(domain.LastContact) > d.config.DomainCheck*10 {
				// its been a while, make sure they are still alive
				go d.rpcHeartbeat(ctx, domainGuard)
			}
			return nil
		})
		return true
	})
}

func (d *Dominion) checkServices(ctx context.Context, _ time.Time) {
	dependencies := make(map[string]int)

	d.domains.Range(func(id uuid.UUID, domainGuard *domain.DomainGuard) bool {
		domainGuard.LatchRead(func(domain domain.Domain) error {

			// find service dependencies
			for _, serviceIdent := range domain.Services {
				serviceType := serviceIdent.Type
				for _, dependency := range config.GetServicesConfig().Services[serviceType].Dependencies {
					dependencies[dependency]++
				}
			}
			return nil
		})
		return true
	})

	// Check that dependencies exist
	for dependency := range dependencies {

		if len(d.findService(dependency)) == 0 {
			candidates := d.findServiceCandidates(dependency)
			if len(candidates) == 0 {
				system.Errorf("No candidates available for %v", dependency)
				continue
			}

			// TODO Handle multiples
			candidate := candidates[0]
			domainGuard, ok := d.domains.Load(candidate.ID)
			if !ok {
				system.Errorf("Viable candidate no longer available for %v", dependency)
				continue
			}

			go d.rpcStartService(ctx, domainGuard, dependency)
		}
	}
}

func (d *Dominion) hostDominion(ctx context.Context) error {
	routineName := "hostDominion"
	system.LogRoutinef(routineName, "Starting routine")

	address := fmt.Sprintf("%s:%v", "", d.Address.Port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("hostDominion() Failed to listen: %w", err)
	}

	server := grpc.NewServer()
	pb.RegisterDominionServer(server, d)
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
