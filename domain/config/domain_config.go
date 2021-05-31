package config

import (
	"fmt"
	"path"
	"strconv"
	"time"

	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/system/config"
)

type (
	// DomainConfig holds all the information required to start a Domain
	DomainConfig struct {
		// DomainIdentity holds identifying information
		ident.DomainIdentity

		// LogFile is where logs are sent to
		LogFile string

		// DialTimeout is how long a domain will wait for a grpc.ClientConn to establish
		DialTimeout time.Duration
		// IsolationCheck is the duration between isolation checks
		IsolationCheck time.Duration
		// ServiceCheck is the length of time after which a domain will send a hearbeat
		ServiceCheck time.Duration
	}
)

func (c *DomainConfig) Patch() error {
	patchedIdent, err := config.Patch(c.Identity)
	if err != nil {
		return err
	}
	c.Identity = patchedIdent

	if c.LogFile == "" {
		c.LogFile = path.Join(".", "logs", c.ID.String(), "domain.log")
		fmt.Printf("LogFile not specified, defaulting to '%v'.\n", c.LogFile)
	}

	if c.DialTimeout == 0 {
		c.DialTimeout = time.Duration(100000000)
		fmt.Printf("DialTimeout not specified, defaulting to %v.\n", c.DialTimeout)
	}

	if c.IsolationCheck == 0 {
		c.IsolationCheck = time.Duration(5000000000)
		fmt.Printf("IsolationCheck not specified, defaulting to %v.\n", c.IsolationCheck)
	}

	if c.ServiceCheck == 0 {
		c.ServiceCheck = time.Duration(2000000000)
		fmt.Printf("ServiceCheck not specified, defaulting to %v.\n", c.ServiceCheck)
	}
	return nil
}

func (c DomainConfig) Check() error {
	err := config.Check(c.Identity)
	if err != nil {
		return err
	}

	if len(c.Traits) == 0 {
		fmt.Println("Warning: Traits were not defined.")
	}

	if c.LogFile == "" {
		return fmt.Errorf("configuration variable 'LogFile' was not set")
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
