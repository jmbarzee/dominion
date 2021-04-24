package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"

	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/system"
	"github.com/jmbarzee/dominion/system/config"
)

type (
	// DominionConfig holds all the information required to start a Dominion
	DominionConfig struct {
		// DominionIdentity holds identifying information
		ident.DominionIdentity

		// DialTimeout is how long a domain will wait for a grpc.ClientConn to establish
		DialTimeout time.Duration
		// DomainCheck is the length of time after which a dominion will send a heartbeat
		DomainCheck time.Duration
		// ServiceCheck is the length of time after which a dominion check service dependency
		ServiceCheck time.Duration

		// Services is a map of service type to service config
		Services map[string]ServiceDefinition
	}
)

var dominionConfig *DominionConfig

// GetDominionConfig returns the singleton DominionConfig
func GetDominionConfig() DominionConfig {
	if dominionConfig == nil {
		system.Panic(errors.New("dominionConfig has not been intialized"))
	}
	return *dominionConfig
}

func setupDominionConfigFromTOML(configFilePath string) error {
	if dominionConfig != nil {
		return errors.New("dominionConfig has already been intialized")
	}

	bytes, err := config.ReadWholeConfigFile(configFilePath)
	if err != nil {
		return err
	}

	c := &DominionConfig{}
	if err = toml.Unmarshal(bytes, c); err != nil {
		return err
	}

	c.Identity, err = config.Patch(c.Identity)
	if err != nil {
		return err
	}

	if err = c.check(); err != nil {
		return err
	}

	dominionConfig = c
	return nil
}

func (c DominionConfig) check() error {

	err := config.Check(c.Identity)
	if err != nil {
		return err
	}

	if c.DialTimeout == 0 {
		return fmt.Errorf("configuration variable 'ConnectionConfig.DialTimeout' was not set")
	}
	if c.DomainCheck == 0 {
		return fmt.Errorf("configuration variable 'ConnectionConfig.HeartbeatCheck' was not set")
	}
	if c.ServiceCheck == 0 {
		return fmt.Errorf("configuration variable 'ServiceHierarchyConfig.DependencyCheck' was not set")
	}
	return nil
}

func (c DominionConfig) String() string {
	dumpMsg := "\tIdentity: " + c.Identity.String() + "\n" +
		"\tDialTimeout: " + c.DialTimeout.String() + "\n" +
		"\tDomainCheck: " + c.DomainCheck.String() + "\n" +
		"\tServiceCheck: " + c.ServiceCheck.String() + "\n" +
		"\tServices: {\n"
	for serviceType, serviceConfig := range c.Services {
		dumpMsg += "\t\t" + serviceType + ": " + serviceConfig.String() + ",\n"
	}
	dumpMsg += "\t}\n"

	return dumpMsg
}
