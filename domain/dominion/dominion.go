package dominion

import (
	"time"

	"github.com/jmbarzee/dominion/ident"
)

type Dominion struct {
	//Identity holds the identifying information of the service
	ident.Identity

	// LastContact is the last time a service replied to a rpc
	LastContact time.Time
}
