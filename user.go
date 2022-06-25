package secretify

import "time"

// User represents user domain model
type User struct {
	Base
	Username string `json:"username"`

	Active    bool      `json:"active"`
	LastLogin time.Time `json:"last_login,omitempty"`

	Token string `json:"-"`
}

// UpdateLastLogin updates last login field
func (u *User) UpdateLastLogin(token string) {
	u.Token = token
	u.LastLogin = time.Now()
}

// AuthUser represents data stored in JWT token for user
type AuthUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
