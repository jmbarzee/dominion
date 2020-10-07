package service

import (
	"sync"

	"github.com/jmbarzee/dominion/identity"
)

// ServiceGuard protects a service for concurrent access
type ServiceGuard struct {
	// RWMutex controls gates to the service's data
	rwmutex sync.RWMutex
	// service is the Service
	service Service
}

// NewServiceGuard returns a new ServiceGuard with the passed identity
func NewServiceGuard(identity identity.ServiceIdentity) *ServiceGuard {
	return &ServiceGuard{
		service: Service{
			ServiceIdentity: identity,
		},
	}
}

// LatchWrite offers write access to the Service
func (d *ServiceGuard) LatchWrite(operation func(*Service) error) error {
	d.rwmutex.Lock()
	err := operation(&d.service)
	d.rwmutex.Unlock()
	return err
}

// LatchRead offers read access to the Service
func (d *ServiceGuard) LatchRead(operation func(Service) error) error {
	d.rwmutex.RLock()
	err := operation(d.service)
	d.rwmutex.RUnlock()
	return err
}
