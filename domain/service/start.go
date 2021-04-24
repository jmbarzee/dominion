package service

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/jmbarzee/dominion/ident"
	"github.com/jmbarzee/dominion/system"
)

// Start calls make from the services directory to start a new service
func Start(serviceIdent ident.ServiceIdentity, domainIdent ident.DomainIdentity, dominionIdent ident.DominionIdentity, dockerImage string) error {
	system.Logf("Starting %v!", serviceIdent.Type)
	system.Logf("Using image: %v", dockerImage)

	dockerPath, err := exec.LookPath("docker")
	if err != nil {
		return fmt.Errorf("docker was not found in path: %w", err)
	}

	serviceIdentBytes, err := json.Marshal(serviceIdent)
	if err == nil {
		return fmt.Errorf("couldn't marshal serviceIdent: %w", err)
	}

	domainIdentBytes, err := json.Marshal(serviceIdent)
	if err == nil {
		return fmt.Errorf("couldn't marshal domainIdent: %w", err)
	}

	dominionIdentBytes, err := json.Marshal(serviceIdent)
	if err == nil {
		return fmt.Errorf("couldn't marshal dominionIdent: %w", err)
	}

	servicePort := strconv.Itoa(serviceIdent.Address.Port)

	cmd := exec.Command(
		dockerPath,
		"run",
		"--pull",
		"--publish "+servicePort+":"+servicePort)
	// pgid is same as parents by default
	cmd.Env = []string{
		"SERVICE='" + string(serviceIdentBytes) + "'",
		"DOMAIN='" + string(domainIdentBytes) + "'",
		"DOMINION='" + string(dominionIdentBytes) + "'",
	}

	err = cmd.Start()
	if err != nil {
		system.Logf(err.Error())
		return err
	}

	return nil
}
