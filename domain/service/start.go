package service

import (
	"fmt"
	"net"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/jmbarzee/dominion/system"
)

const rootPath = "/usr/local/dominion/services"

// Start calls make from the services directory to start a new service
func Start(serviceType string, ip net.IP, dominionPort int, domainUUID string, servicePort int) error {
	system.Logf("Starting %v!", serviceType)

	makefilePath := path.Join(rootPath, strings.ToLower(serviceType))
	makePath, err := exec.LookPath("make")
	if err != nil {
		return fmt.Errorf("make was not found in path: %w", err)
	}

	cmd := exec.Command(
		makePath,
		"-C", makefilePath,
		"start",
		"DOMINION_IP="+ip.String(),
		"DOMINION_PORT="+strconv.Itoa(dominionPort),
		"DOMAIN_UUID="+domainUUID,
		"SERVICE_PORT="+strconv.Itoa(servicePort))
	// pgid is same as parents by default

	err = cmd.Start()
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	return nil
}

// Build calls make from the services directory to build a new service
func Build(serviceType string) error {
	system.Logf("Building %v!", serviceType)

	makefilePath := path.Join(rootPath, strings.ToLower(serviceType))
	makePath, err := exec.LookPath("make")
	if err != nil {
		return fmt.Errorf("make was not found in path: %w", err)
	}

	cmd := exec.Command(
		makePath,
		"-C", makefilePath,
		"build")

	err = cmd.Start()
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	return nil
}
