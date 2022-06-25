package secretify

import (
	"time"
)

// Base contains common fields for all tables
type Base struct {
	ID        int       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty" gorm:"index"`
}

// ListQuery holds company/location data used for list db queries
type ListQuery struct {
	Query string
	ID    int
}

type HTTPErrorResponse struct {
	Error string `json:"error"`
}

type HTTPOKResponse struct {
	Data interface{} `json:"data"`
}
