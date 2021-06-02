package config

import (
	"encoding/json"
	"time"

	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/system/config"
)

// ServiceConfig contains all information needed to start a service
type ServiceConfig struct {
	// ServiceIdentity holds identifying information
	ident.ServiceIdentity

	// Identity holds identifying information for the domain
	DomainIdentity ident.Identity

	// DominionIdentity holds identifying information for the dominion
	DominionIdentity ident.Identity

	// LogFile is where logs are sent to
	LogFile string
}

// DefaultServiceDialTimeout is the default time out for grpc dial operations
const DefaultServiceDialTimeout = time.Millisecond * 100

// FromEnv builds a ServiceConfig from the environment and arguments
func FromEnv(serviceType string) (ServiceConfig, error) {
	serviceIdent := ident.ServiceIdentity{}
	serviceIdentString := config.RequireEnvString("SERVICE_IDENTITY")
	if err := json.Unmarshal([]byte(serviceIdentString), &serviceIdent); err != nil {
		return ServiceConfig{}, err
	}

	domainIdent := ident.Identity{}
	domainIdentString := config.RequireEnvString("DOMAIN_IDENTITY")
	if err := json.Unmarshal([]byte(domainIdentString), &domainIdent); err != nil {
		return ServiceConfig{}, err
	}

	dominionIdent := ident.Identity{}
	dominionIdentString := config.RequireEnvString("DOMINION_IDENTITY")
	if err := json.Unmarshal([]byte(dominionIdentString), &dominionIdent); err != nil {
		return ServiceConfig{}, err
	}

	logFile := config.RequireEnvString("SERVICE_LOG_FILE")

	return ServiceConfig{
		ServiceIdentity:  serviceIdent,
		DomainIdentity:   domainIdent,
		DominionIdentity: dominionIdent,
		LogFile:          logFile,
	}, nil
}
