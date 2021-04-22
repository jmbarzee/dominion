package system

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"
)

// func GetOutboundIP() (net.IP, error) {
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

func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return net.IP{}, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}

// RoutineOperation offers repeatedly runs op and then waits for wait.
// cancleing ctx will end the cycle
func RoutineOperation(ctx context.Context, routineName string, wait time.Duration, op func(context.Context, time.Time)) {
	LogRoutinef(routineName, "Starting routine")
	ticker := time.NewTicker(wait)

Loop:
	for {
		select {
		case t := <-ticker.C:
			op(ctx, t)
		case <-ctx.Done():
			break Loop
		}
	}
	LogRoutinef(routineName, "Stopping routine")
}

// RequireEnv finds the value of the required variable or panics
func RequireEnv(varName string) string {
	value := os.Getenv(varName)
	if value == "" {
		panic(fmt.Errorf("could not find required environment variable %s", varName))
	}
	return value
}
