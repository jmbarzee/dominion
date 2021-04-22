package config

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
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

func GetServicesConfig() ServicesConfig {
	if servicesConfig == nil {
		system.Panic(errors.New("servicesConfig has not been intialized"))
	}
	return *servicesConfig
}

func setupServicesConfigFromTOML(configFilePath string) error {
	if servicesConfig != nil {
		return errors.New("servicesConfig has already been intialized")
	}

	configFile, err := os.OpenFile(configFilePath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		return err
	}

	config := &ServicesConfig{}
	if err = toml.Unmarshal(bytes, config); err != nil {
		return err
	}

	servicesConfig = config
	return nil
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
