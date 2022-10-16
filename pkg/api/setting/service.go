package setting

import (
	utilconfig "github.com/DarioCalovic/secretify/pkg/util/config"
)

// Service represents setting application interface
type Service interface {
	Meta() utilconfig.Meta
	Policy() utilconfig.Policy
	Storage() utilconfig.Storage
	Outlook() utilconfig.Outlook
}

// User represents setting application service
type Setting struct {
	cfg *utilconfig.Configuration
}

// New creates new setting application service
func New(cfg *utilconfig.Configuration) *Setting {
	return &Setting{cfg}
}

// Initialize initalizes setting application service with defaults
func Initialize(cfg *utilconfig.Configuration) *Setting {
	return New(cfg)
}
