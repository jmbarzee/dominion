package system

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"

	"github.com/google/uuid"
)

const (

	// PathBase is the location of all dominion src files
	PathBase = "/usr/local/dominion"
	// PathLogs is the location of all dominion log files
	PathLogs = PathBase + "/logs"
	// PathServices is the location of all dominion services
	PathServices = PathBase + "/services"
)

// log is where normal & debugging messages are dumped to
var logger *log.Logger
var closeFile func() error

// Setup initializes logging and signal handling
func Setup(id uuid.UUID, fileName string) error {
	if logger != nil {
		return errors.New("system.Setup has already been called")
	}

	pathIDLogs := path.Join(PathLogs, id.String())
	err := os.MkdirAll(pathIDLogs, 0755)
	if err != nil {
		return fmt.Errorf("failed to check ensure log directory \"%v\" exists: %w", pathIDLogs, err)
	}

	pathIDLogFile := path.Join(pathIDLogs, fileName+".log")
	logFile, err := os.OpenFile(pathIDLogFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to check open log file \"%v\" exists: %w", pathIDLogFile, err)
	}
	closeFile := logFile.Close
	logger = log.New(logFile, "", log.LstdFlags)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		if err := closeFile(); err != nil {
			panic(err)
		}
		os.Exit(0)
	}()
	return nil
}

func Panic(err error) {
	if closeFile != nil {
		if err := closeFile(); err != nil {
			panic(err)
		}
	}
	panic(err)
}

func Logf(fmt string, args ...interface{}) {
	if logger == nil {
		panic(errors.New("system.Setup has not been called"))
	}
	logger.Printf(fmt, args...)
}

func LogRoutinef(routineName, fmts string, args ...interface{}) {
	if logger == nil {
		panic(errors.New("system.Setup has not been called"))
	}
	prefix := fmt.Sprintf("Routine [%s]: ", routineName)
	logger.Printf(prefix+fmts, args...)
}

func LogRPCf(rpcName, fmts string, args ...interface{}) {
	if logger == nil {
		panic(errors.New("system.Setup has not been called"))
	}
	prefix := fmt.Sprintf("RPC-%s	: ", rpcName)
	logger.Printf(prefix+fmts, args...)
}

func Errorf(fmts string, args ...interface{}) {
	if logger == nil {
		panic(errors.New("system.Setup has not been called"))
	}
	err := fmt.Errorf(fmts, args...)
	logger.Printf("Error: %v", err)
}
