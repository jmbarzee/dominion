package config

import (
	"errors"

	"github.com/jmbarzee/dominion/system"
)

type (
	ServicesConfig struct {
		Services map[string]ServiceDefinition
	}

	// ServiceDefinition defines a single service in the service hierarchy
	ServiceDefinition struct {
		// DockerImage is the image:tag of the service which can be pulled and started
		DockerImage string
		// Dependencies is the list of service types which this service depends on
		Dependencies []string
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
	msg += "])"

	return msg
}
