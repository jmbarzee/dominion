package domain

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/google/uuid"
	"github.com/jmbarzee/dominion/system"
)

// DomainMap is a wrapper for sync.map to ensure better type safety
type DomainMap struct {
	sMap *sync.Map
}

// NewDomainMap returns a new DomainMap
func NewDomainMap() DomainMap {
	return DomainMap{
		sMap: &sync.Map{},
	}
}

// Delete removes a Domain from the DomainMap
func (m DomainMap) Delete(id uuid.UUID) {
	m.sMap.Delete(id)
}

// Load offers access to a Domain
func (m DomainMap) Load(id uuid.UUID) (*DomainGuard, bool) {
	v, ok := m.sMap.Load(id)
	if v == nil {
		return nil, ok
	}
	return v.(*DomainGuard), ok
}

// LoadOrStore offers access to a Domain or stores a new one
func (m DomainMap) LoadOrStore(id uuid.UUID, mem *DomainGuard) (*Domain, bool) {
	v, loaded := m.sMap.LoadOrStore(id, mem)
	if !loaded {
		return nil, loaded
	}
	return v.(*Domain), loaded
}

// Range iterates across all Domain in the DomainMap
func (m DomainMap) Range(f func(id uuid.UUID, mem *DomainGuard) bool) {
	m.sMap.Range(func(k, v interface{}) bool {
		id := k.(uuid.UUID)
		mem := v.(*DomainGuard)
		return f(id, mem)
	})
}

// SizeEstimate only garuntees that the number of all existing keys
// in some length of time is equal to the result
func (m *DomainMap) SizeEstimate() int {
	size := int32(0)
	m.sMap.Range(func(k, v interface{}) bool {
		atomic.AddInt32(&size, 1)
		return true
	})
	return int(size)
}

// Store stores a new Domain in the DomainMap
func (m *DomainMap) Store(id uuid.UUID, mem *DomainGuard) {
	if mem == nil {
		system.Panic(fmt.Errorf("Store() mem was nil"))
	}
	m.sMap.Store(id, mem)
}
