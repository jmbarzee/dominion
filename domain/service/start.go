package service

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"

	"github.com/google/uuid"
	"github.com/jmbarzee/dominion/system"
)

// Start calls make from the services directory to start a new service
func Start(serviceType string, dockerImage string, ip net.IP, dominionPort int, domainID uuid.UUID, servicePort int) error {
	system.Logf("Starting %v!", serviceType)
	system.Logf("Using image: %v", dockerImage)

	dockerPath, err := exec.LookPath("docker")
	if err != nil {
		return fmt.Errorf("docker was not found in path: %w", err)
	}

	servicePortString := strconv.Itoa(servicePort)

	cmd := exec.Command(
		dockerPath,
		"run",
		"--pull",
		"--publish "+servicePortString+":"+servicePortString)
	// pgid is same as parents by default
	cmd.Env = []string{
		"DOMINION_IP=" + ip.String(),
		"DOMINION_PORT=" + strconv.Itoa(dominionPort),
		"DOMAIN_ID=" + domainID.String(),
		"SERVICE_PORT=" + servicePortString,
	}

	err = cmd.Start()
	if err != nil {
		system.Logf(err.Error())
		return err
	}

	return nil
}
