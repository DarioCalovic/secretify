package setting

import (
	utilconfig "github.com/DarioCalovic/secretify/pkg/util/config"
)

// General returns meta information
func (s *Setting) Meta() utilconfig.Meta {
	return *s.cfg.Meta
}

// Policy returns policy information
func (s *Setting) Policy() utilconfig.Policy {
	return *s.cfg.Policy
}

// Storage returns storage information
func (s *Setting) Storage() utilconfig.Storage {
	return *s.cfg.Storage
}

// Outlook returns storage information
func (s *Setting) Outlook() utilconfig.Outlook {
	return *s.cfg.Outlook
}
