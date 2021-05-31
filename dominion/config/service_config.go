package config

import (
	"errors"

	"github.com/jmbarzee/dominion/system"
)

type (
	ServicesConfig struct {
		Services map[string]ServiceDefinition
	}

	// ServiceDefinition defines a single service in the service hiarchy
	ServiceDefinition struct {
		// DockerImage is the image:tag of the service which can be pulled and started
		DockerImage string
		// Dependencies is the list of service types which this service depends on
		Dependencies []string
		// Traits is the list of triats required by a domain to be able to run a service
		Traits []string
	}
)

var servicesConfig *ServicesConfig

// SetServicesConfig gets the singleton servicesConfig
func GetServicesConfig() ServicesConfig {
	if servicesConfig == nil {
		system.Panic(errors.New("servicesConfig has not been intialized"))
	}
	return *servicesConfig
}

// SetServicesConfig sets the singleton servicesConfig
func SetServicesConfig(c ServicesConfig) {
	if servicesConfig != nil {
		system.Panic(errors.New("servicesConfig has already been intialized"))
	}
	servicesConfig = &c
}

func (s ServiceDefinition) String() string {
	msg := "(" + string(s.DockerImage) + ", ["

	for _, dependency := range s.Dependencies {
		msg += dependency + ", "
	}
	msg += "], ["

	for _, trait := range s.Traits {
		msg += trait + ", "
	}
	msg += "])"

	return msg
}
