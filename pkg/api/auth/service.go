package auth

import (
	root "github.com/DarioCalovic/secretify"
	"github.com/DarioCalovic/secretify/pkg/api/auth/platform"
	utildb "github.com/DarioCalovic/secretify/pkg/util/db"
)

// New creates new iam service
func New(db utildb.DB, j TokenGenerator, sec Securer, aur Repository) *Auth {
	return &Auth{
		db:  db,
		tg:  j,
		sec: sec,
		rep: aur,
	}
}

// Initialize initializes auth application service
func Initialize(db utildb.DB, j TokenGenerator, sec Securer) *Auth {
	switch db.(type) {
	case *utildb.SQLiteDB:
		return New(db, j, sec, platform.NewSQLiteAuthUserRepository())
	}
	return New(db, j, sec, platform.NewSQLiteAuthUserRepository())
}

// Service represents auth service interface
type Service interface {
	Authenticate(string, string) (root.AuthToken, error)
	Me() (root.AuthUser, error)
}

// Auth represents auth application service
type Auth struct {
	tg  TokenGenerator
	rep Repository
	db  utildb.DB
	sec Securer
}

// Repository represents user repository interface
type Repository interface {
	Create(utildb.DB, root.User) (root.User, error)
	View(utildb.DB, int) (root.User, error)
	FindByUsername(utildb.DB, string) (root.User, error)
	FindByToken(utildb.DB, string) (root.User, error)
	Update(utildb.DB, root.User) error
}

// TokenGenerator represents token generator (jwt) interface
type TokenGenerator interface {
	GenerateToken(root.User) (string, error)
}

// Securer represents security interface
type Securer interface {
	HashMatchesPassword(string, string) bool
	Token(string) string
}
