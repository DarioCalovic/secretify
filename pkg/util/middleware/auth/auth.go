package auth

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// TokenParser represents JWT token parser
type TokenParser interface {
	ParseToken(string) (*jwt.Token, error)
}

func Middleware(tokenParser TokenParser) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := tokenParser.ParseToken(r.Header.Get("Authorization"))
			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			claims := token.Claims.(jwt.MapClaims)
			_ = claims

			// context.Set(r, "token", "asdf")

			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		})
	}
}

func MiddlewareEmailValidity(tokenParser TokenParser) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := tokenParser.ParseToken(r.Header.Get("Authorization"))
			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			claims := token.Claims.(jwt.MapClaims)
			_ = claims

			// context.Set(r, "token", "asdf")

			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		})
	}
}
