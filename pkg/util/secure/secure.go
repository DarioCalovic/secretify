package secure

import (
	"fmt"
	"hash"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// New initializes security service
func New(h hash.Hash) *Service {
	return &Service{h: h}
}

// Service holds security related methods
type Service struct {
	minPWStr int
	h        hash.Hash
}

// Hash hashes the password using bcrypt
func (*Service) Hash(password string) string {
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPW)
}

// HashMatchesPassword matches hash with password. Returns true if hash and password match.
func (*Service) HashMatchesPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// Token generates new unique token
func (s *Service) Token(str string) string {
	s.h.Reset()
	fmt.Fprintf(s.h, "%s%s", str, strconv.Itoa(time.Now().Nanosecond()))
	return fmt.Sprintf("%x", s.h.Sum(nil))
}
