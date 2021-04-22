package config

import (
	"net"
	"os"
	"strconv"
	"time"
)

// ServiceConfig contains all information needed to start a service
type ServiceConfig struct {
	DominionIP   net.IP
	DominionPort int
	DomainUUID   string
	ServiceType  string
	ServicePort  int
}

// DefaultServiceDialTimeout is the default time out for grpc dial operations
const DefaultServiceDialTimeout = time.Millisecond * 100

// FromEnv builds a ServiceConfig from the environment and arguments
func FromEnv(serviceType string) (config ServiceConfig, err error) {
	dominionIPString := os.Getenv("DOMINION_IP")
	dominionIP := net.ParseIP(dominionIPString)

	dominionPortString := os.Getenv("DOMINION_PORT")
	dominionPort64, err := strconv.ParseInt(dominionPortString, 0, 32)
	if err != nil {
		return ServiceConfig{}, err
	}
	dominionPort := int(dominionPort64)

	domainUUID := os.Getenv("DOMAIN_UUID")

	servicePortString := os.Getenv("SERVICE_PORT")
	servicePort64, err := strconv.ParseInt(servicePortString, 0, 32)
	if err != nil {
		return ServiceConfig{}, err
	}
	servicePort := int(servicePort64)

	return ServiceConfig{
		DominionIP:   dominionIP,
		DominionPort: dominionPort,
		DomainUUID:   domainUUID,
		ServiceType:  serviceType,
		ServicePort:  servicePort,
	}, nil
}
