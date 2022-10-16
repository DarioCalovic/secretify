package outlook

import (
	"github.com/DarioCalovic/secretify/pkg/api/setting"
)

// Service represents outlook application interface
type Service interface {
	Manifest() []byte
}

// Outlok represents outlook application service
type Outlook struct {
	cfgSvc setting.Service
}

// New creates new outlook application service
func New(cfgSvc setting.Service) *Outlook {
	return &Outlook{cfgSvc}
}

// Initialize initalizes outlook application service with defaults
func Initialize(cfgSvc setting.Service) *Outlook {
	return New(cfgSvc)
}
