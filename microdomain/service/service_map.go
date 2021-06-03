package service

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/jmbarzee/dominion/system"
)

// ServiceMap is a wrapper for sync.map to ensure better type safety
type ServiceMap struct {
	sMap *sync.Map
}

// NewServiceMap returns a new ServiceMap
func NewServiceMap() ServiceMap {
	return ServiceMap{
		sMap: &sync.Map{},
	}
}

// Delete removes a Service from the ServiceMap
func (m ServiceMap) Delete(uuid string) {
	m.sMap.Delete(uuid)
}

// Load offers access to a Service
func (m ServiceMap) Load(uuid string) (*ServiceGuard, bool) {
	v, ok := m.sMap.Load(uuid)
	if v == nil {
		return nil, ok
	}
	return v.(*ServiceGuard), ok
}

// LoadOrStore offers access to a Service or stores a new one
func (m ServiceMap) LoadOrStore(uuid string, mem *ServiceGuard) (*Service, bool) {
	v, loaded := m.sMap.LoadOrStore(uuid, mem)
	if !loaded {
		return nil, loaded
	}
	return v.(*Service), loaded
}

// Range iterates across all Services in the ServiceMap
func (m ServiceMap) Range(f func(uuid string, mem *ServiceGuard) bool) {
	m.sMap.Range(func(k, v interface{}) bool {
		uuid := k.(string)
		mem := v.(*ServiceGuard)
		return f(uuid, mem)
	})
}

// SizeEstimate only garuntees that the number of all existing keys
// in some length of time is equal to the result
func (m *ServiceMap) SizeEstimate() int {
	size := int32(0)
	m.sMap.Range(func(k, v interface{}) bool {
		atomic.AddInt32(&size, 1)
		return true
	})
	return int(size)
}

// Store stores a new Service in the ServiceMap
func (m *ServiceMap) Store(uuid string, mem *ServiceGuard) {
	if mem == nil {
		system.Panic(fmt.Errorf("Store() mem was nil"))
	}
	m.sMap.Store(uuid, mem)
}
