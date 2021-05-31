package system

import (
	"context"
	"time"
)

// RoutineOperation offers repeatedly runs op and then waits for wait.
// canceling ctx will end the cycle
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
