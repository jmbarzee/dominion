package dominion

import (
	"sync"

	"github.com/jmbarzee/dominion/ident"
)

// DominionGuard protects a dominion for concurrent access
type DominionGuard struct {
	// RWMutex controls gates to the dominion's data
	rwmutex sync.RWMutex
	// data is the Dominions
	dominion Dominion
}

// NewDominionGuard returns a new DominionGuard with the passed identity
func NewDominionGuard(identity ident.DominionIdentity) *DominionGuard {
	return &DominionGuard{
		dominion: Dominion{
			DominionIdentity: identity,
		},
	}
}

func (d *DominionGuard) LatchWrite(operation func(*Dominion) error) error {
	d.rwmutex.Lock()
	err := operation(&d.dominion)
	d.rwmutex.Unlock()
	return err
}

func (d *DominionGuard) LatchRead(operation func(*Dominion) error) error {
	d.rwmutex.RLock()
	err := operation(&d.dominion)
	d.rwmutex.RUnlock()
	return err
}
