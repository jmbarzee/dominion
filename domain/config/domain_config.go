package config

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/system"
	"github.com/jmbarzee/dominion/system/config"
)

type (
	// DomainConfig holds all the information required to start a Domain
	DomainConfig struct {
		// DomainIdentity holds identifying information
		ident.DomainIdentity

		// DialTimeout is how long a domain will wait for a grpc.ClientConn to establish
		DialTimeout time.Duration
		// IsolationCheck is the duration between isolation checks
		IsolationCheck time.Duration
		// ServiceCheck is the length of time after which a domain will send a hearbeat
		ServiceCheck time.Duration
	}
)

var domainConfig *DomainConfig

// GetDominionConfig returns the singleton DominionConfig
func GetDomainConfig() DomainConfig {
	if domainConfig == nil {
		system.Panic(errors.New("domainConfig has not been intialized"))
	}
	return *domainConfig
}

// SetupFromTOML produces a default configuration
func SetupFromTOML(configFilePath string) error {
	if domainConfig != nil {
		return errors.New("domainConfig has already been intialized")
	}

	bytes, err := config.ReadWholeConfigFile(configFilePath)
	if err != nil {
		return err
	}

	c := &DomainConfig{}
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

	domainConfig = c
	return nil
}

func (c DomainConfig) check() error {

	err := config.Check(c.Identity)
	if err != nil {
		return err
	}

	if len(c.Traits) == 0 {
		return fmt.Errorf("configuration variable 'Traits' were not set")
	}

	if c.DialTimeout == 0 {
		return fmt.Errorf("configuration variable 'DialTimeout' was not set")
	}
	if c.IsolationCheck == 0 {
		return fmt.Errorf("configuration variable 'IsolationCheck' was not set")
	}
	if c.ServiceCheck == 0 {
		return fmt.Errorf("configuration variable 'ServiceCheck' was not set")
	}

	return nil
}

func (c DomainConfig) String() string {

	dumpMsg := "\tIdentity: " + c.Identity.String() + "\n"

	dumpMsg += "\tTraits: ["
	for _, trait := range c.Traits {
		dumpMsg += trait + ", "
	}
	dumpMsg += "]\n"

	dumpMsg += "\tDialTimeout: \n" + strconv.Itoa(int(c.DialTimeout)) +
		"\tServiceCheck: \n" + strconv.Itoa(int(c.ServiceCheck)) +
		"\tIsolationCheck: \n" + strconv.Itoa(int(c.IsolationCheck))
	return dumpMsg
}
