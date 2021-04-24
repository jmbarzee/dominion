package config

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/google/uuid"
	"github.com/jmbarzee/dominion"
	"github.com/jmbarzee/dominion/ident"
)

func Patch(i ident.Identity) (ident.Identity, error) {
	if i.ID == uuid.Nil {
		i.ID = uuid.New()
		fmt.Println("Warning: No id was specified, so one was generated.")
		fmt.Println("Warning: All reconnecting and persistence features depend")
		fmt.Println("Warning: on ids persisting beyond start and stop.")
	}

	if i.Version.Major == 0 && i.Version.Minor == 0 && i.Version.Patch == 0 {
		i.Version = dominion.Version()
		fmt.Println("Warning: No Version was specified, so the Dominion Version was used.")
	}

	var err error
	i.Address, err = patch(i.Address)
	if err != nil {
		return ident.Identity{}, err
	}
	return i, nil
}

func Check(i ident.Identity) error {
	if i.ID == uuid.Nil {
		return fmt.Errorf("configuration variable 'Identity.ID' was not set")
	}

	if i.Version.Major == 0 && i.Version.Minor == 0 && i.Version.Patch == 0 {
		return fmt.Errorf("configuration variable 'Identity.Version' was not set")
	}

	if err := check(i.Address); err != nil {
		return err
	}
	return nil
}

func patch(a ident.Address) (ident.Address, error) {
	if a.IP == nil {
		var err error
		a.IP, err = getOutboundIP()
		fmt.Println("Warning: No IP was specified, so it was deduced.")
		if err != nil {
			return ident.Address{}, fmt.Errorf("couldn't patch missing IP: %w", err)
		}
	}

	// We do not patch the port as it will either generate a collision
	// or not be a port exposed by the encapsulating docker container.
	return a, nil
}

func check(a ident.Address) error {
	if a.Port == 0 {
		return fmt.Errorf("configuration variable 'Identity.Address.Port' was not set")
	}

	if a.IP == nil {
		return fmt.Errorf("configuration variable 'Identity.Address.IP' was not set")
	}
	return nil
}

// func getOutboundIP() (net.IP, error) {
// 	ifaces, err := net.Interfaces()
// 	if err != nil {
// 		return nil, err
// 	}
// 	for _, i := range ifaces {
// 		Logf("interface [%s]: %v %v", i.Name, i.Index, i.HardwareAddr)
// 		addrs, err := i.Addrs()
// 		if err != nil {
// 			return nil, err
// 		}
// 		for _, addr := range addrs {
// 			var ip net.IP
// 			switch v := addr.(type) {
// 			case *net.IPNet:
// 				ip = v.IP
// 				v.IP
// 				Logf("IPNet: %s", ip.String())
// 			case *net.IPAddr:
// 				ip = v.IP
// 				Logf("IPAddr: %s", ip.String())
// 			}
// 			// process IP address
// 		}
// 	}
// 	return nil, fmt.Errorf("FAKE ERROR")
// }

func getOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return net.IP{}, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}

func ReadWholeConfigFile(configFilePath string) ([]byte, error) {
	configFile, err := os.OpenFile(configFilePath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(configFile)
}
