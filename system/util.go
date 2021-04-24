package system

import (
	"context"
	"fmt"
	"os"
	"time"
)

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
