package secretify

import "time"

type Secret struct {
	Base
	Identifier    string `gorm:"unique"`
	Cipher        string
	Email         string
	WebhookAddr   string
	ExpiresAt     time.Time
	HasPassphrase bool
	RevealOnce    bool
	FileID        int
	File          File `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
