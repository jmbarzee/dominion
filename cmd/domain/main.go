package main

import (
	"context"
	"encoding/json"
	"net"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/dominion/domain"
	"github.com/jmbarzee/dominion/domain/config"
	"github.com/jmbarzee/dominion/ident"
	syscfg "github.com/jmbarzee/dominion/system/config"
)

func main() {
	runtime.GOMAXPROCS(4)

	id := uuid.Nil
	var err error
	if id, err = uuid.Parse(syscfg.RequestEnvString("DOMAIN_ID")); err != nil {
		panic(err)
	}

	var traits []string
	traitsString := syscfg.RequestEnvString("DOMAIN_TRAITS")
	if traitsString != "" {
		if err := json.Unmarshal([]byte(traitsString), &traits); err != nil {
			panic(err)
		}
	}

	c := config.DomainConfig{
		DomainIdentity: ident.DomainIdentity{
			Identity: ident.Identity{
				ID: id,
				Address: ident.Address{
					IP:   net.ParseIP(syscfg.RequestEnvString("DOMAIN_IP")),
					Port: syscfg.RequireEnvInt("DOMAIN_PORT"),
				},
			},
			Traits: traits,
		},
		LogFile:        syscfg.RequestEnvString("DOMAIN_LOG_FILE"),
		DialTimeout:    time.Duration(syscfg.RequestEnvInt("DOMAIN_DIAL_TIMEOUT")),
		IsolationCheck: time.Duration(syscfg.RequestEnvInt("DOMAIN_ISOLATION_CHECK")),
		ServiceCheck:   time.Duration(syscfg.RequestEnvInt("DOMAIN_SERVICE_CHECK")),
	}

	c.Patch()
	c.Check()

	domain, err := domain.NewDomain(c)
	if err != nil {
		panic(err)
	}

	if err := domain.Run(context.Background()); err != nil {
		panic(err)
	}
}
