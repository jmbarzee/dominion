package config

import (
	"fmt"
	"path"
	"time"

	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/system/config"
)

type (
	// DominionConfig holds all the information required to start a Dominion
	DominionConfig struct {
		// DominionIdentity holds identifying information
		ident.DominionIdentity

		// LogFile is where logs are sent to
		LogFile string

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

func (c *DominionConfig) Patch() error {
	patchedIdent, err := config.Patch(c.Identity)
	if err != nil {
		return err
	}
	c.Identity = patchedIdent

	if c.LogFile == "" {
		c.LogFile = path.Join(".", "logs", c.ID.String(), "dominion.log")
		fmt.Printf("LogFile not specified, defaulting to '%v'.\n", c.LogFile)
	}

	if c.DialTimeout == 0 {
		c.DialTimeout = time.Duration(2000000000)
		fmt.Printf("DialTimeout not specified, defaulting to %v.\n", c.DialTimeout)
	}

	if c.DomainCheck == 0 {
		c.DomainCheck = time.Duration(1000000000)
		fmt.Printf("DomainCheck not specified, defaulting to %v.\n", c.DomainCheck)
	}

	if c.ServiceCheck == 0 {
		c.ServiceCheck = time.Duration(2000000000)
		fmt.Printf("ServiceCheck not specified, defaulting to %v.\n", c.ServiceCheck)
	}
	return nil

}

func (c DominionConfig) Check() error {
	err := config.Check(c.Identity)
	if err != nil {
		return err
	}

	if c.LogFile == "" {
		return fmt.Errorf("configuration variable 'LogFile' was not set")
	}

	if c.DialTimeout == 0 {
		return fmt.Errorf("configuration variable 'DialTimeout' was not set")
	}
	if c.DomainCheck == 0 {
		return fmt.Errorf("configuration variable 'HeartbeatCheck' was not set")
	}
	if c.ServiceCheck == 0 {
		return fmt.Errorf("configuration variable 'DependencyCheck' was not set")
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
