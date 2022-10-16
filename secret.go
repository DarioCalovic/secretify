package secretify

import "time"

type Secret struct {
	Base
	Identifier string `gorm:"unique"`

	// Cipher attributes
	Cipher string

	// Policy attributes
	ExpiresAt     time.Time
	HasPassphrase bool
	RevealOnce    bool
	DestroyManual bool

	// File attributes
	FileID int
	File   File `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
