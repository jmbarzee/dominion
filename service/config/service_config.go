package config

import (
	"encoding/json"
	"time"

	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/system"
)

// ServiceConfig contains all information needed to start a service
type ServiceConfig struct {
	// ServiceIdentity holds identifying information
	ident.ServiceIdentity

	// DomainIdentity holds identifying information for the domain
	DomainIdentity ident.DomainIdentity

	// DominionIdentity holds identifying information for the dominion
	DominionIdentity ident.DominionIdentity
}

// DefaultServiceDialTimeout is the default time out for grpc dial operations
const DefaultServiceDialTimeout = time.Millisecond * 100

// FromEnv builds a ServiceConfig from the environment and arguments
func FromEnv(serviceType string) (config ServiceConfig, err error) {
	serviceIdent := ident.ServiceIdentity{}
	serviceIdentString := system.RequireEnv("SERVICE_IDENTITY")
	if err := json.Unmarshal([]byte(serviceIdentString), &serviceIdent); err != nil {
		return ServiceConfig{}, err
	}

	domainIdent := ident.DomainIdentity{}
	domainIdentString := system.RequireEnv("DOMAIN_IDENTITY")
	if err := json.Unmarshal([]byte(domainIdentString), &domainIdent); err != nil {
		return ServiceConfig{}, err
	}

	dominionIdent := ident.DominionIdentity{}
	dominionIdentString := system.RequireEnv("DOMINION_IDENTITY")
	if err := json.Unmarshal([]byte(dominionIdentString), &dominionIdent); err != nil {
		return ServiceConfig{}, err
	}

	return ServiceConfig{
		ServiceIdentity:  serviceIdent,
		DomainIdentity:   domainIdent,
		DominionIdentity: dominionIdent,
	}, nil
}
