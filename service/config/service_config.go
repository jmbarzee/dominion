package config

import (
	"net"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/dominion/system"
)

// ServiceConfig contains all information needed to start a service
type ServiceConfig struct {
	DominionIP   net.IP
	DominionPort int
	DomainID     uuid.UUID
	ServiceType  string
	ServicePort  int
	ServiceID    uuid.UUID
}

// DefaultServiceDialTimeout is the default time out for grpc dial operations
const DefaultServiceDialTimeout = time.Millisecond * 100

// FromEnv builds a ServiceConfig from the environment and arguments
func FromEnv(serviceType string) (config ServiceConfig, err error) {
	dominionIPString := system.RequireEnv("DOMINION_IP")
	dominionIP := net.ParseIP(dominionIPString)

	dominionPortString := system.RequireEnv("DOMINION_PORT")
	dominionPort64, err := strconv.ParseInt(dominionPortString, 0, 32)
	if err != nil {
		return ServiceConfig{}, err
	}
	dominionPort := int(dominionPort64)

	domainIDString := system.RequireEnv("DOMAIN_ID")
	domainID, err := uuid.Parse(domainIDString)
	if err != nil {
		return ServiceConfig{}, err
	}

	servicePortString := system.RequireEnv("SERVICE_PORT")
	servicePort64, err := strconv.ParseInt(servicePortString, 0, 32)
	if err != nil {
		return ServiceConfig{}, err
	}
	servicePort := int(servicePort64)

	serviceID := uuid.New()

	return ServiceConfig{
		DominionIP:   dominionIP,
		DominionPort: dominionPort,
		DomainID:     domainID,
		ServiceType:  serviceType,
		ServicePort:  servicePort,
		ServiceID:    serviceID,
	}, nil
}
