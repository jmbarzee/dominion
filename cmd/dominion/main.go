package main

import (
	"context"
	"net"
	"runtime"
	"time"

	"github.com/google/uuid"
	dd "github.com/jmbarzee/dominion"
	"github.com/jmbarzee/dominion/dominion"
	"github.com/jmbarzee/dominion/dominion/config"
	"github.com/jmbarzee/dominion/ident"
	syscfg "github.com/jmbarzee/dominion/system/config"
)

func main() {
	runtime.GOMAXPROCS(4)

	id := uuid.Nil
	var err error
	if id, err = uuid.Parse(syscfg.RequestEnvString("DOMINION_ID")); err != nil {
		panic(err)
	}

	c := config.DominionConfig{
		Identity: ident.Identity{
			ID:      id,
			Version: dd.Version(),
			Address: ident.Address{
				IP:   net.ParseIP(syscfg.RequestEnvString("DOMINION_IP")),
				Port: syscfg.RequireEnvInt("DOMINION_PORT"),
			},
		},
		LogFile:      syscfg.RequestEnvString("DOMINION_LOG_FILE"),
		DialTimeout:  time.Duration(syscfg.RequestEnvInt("DIAL_TIMEOUT")),
		DomainCheck:  time.Duration(syscfg.RequestEnvInt("ISOLATION_CHECK")),
		ServiceCheck: time.Duration(syscfg.RequestEnvInt("SERVICE_CHECK")),
	}

	c.Patch()
	c.Check()

	dominion, err := dominion.NewDominion(c)
	if err != nil {
		panic(err)
	}

	if err := dominion.Run(context.Background()); err != nil {
		panic(err)
	}
}
