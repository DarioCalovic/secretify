package transport

import (
	"encoding/json"
	"net/http"

	"github.com/DarioCalovic/secretify"
	"github.com/DarioCalovic/secretify/pkg/api/auth"
	"github.com/gorilla/mux"
)

// HTTP represents auth http service
type HTTP struct {
	svc auth.Service
}

// NewHTTP creates new auth http service
func NewHTTP(svc auth.Service, r *mux.Router, mw mux.MiddlewareFunc) {
	h := HTTP{svc}

	r.HandleFunc("/login", h.login).
		Methods("POST")

	mr := r.Methods(http.MethodPost).Subrouter()
	mr.HandleFunc("/me", h.me).
		Methods("GET")
	mr.Use(mw)
}

type credentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *HTTP) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	var cred credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: err.Error(),
		})
		return
	}
	rr, err := h.svc.Authenticate(cred.Username, cred.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: err.Error(),
		})
		return
	}

	_ = rr

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(secretify.HTTPOKResponse{
		Data: nil,
	})
}

func (h *HTTP) me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	user, err := h.svc.Me()
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(secretify.HTTPOKResponse{
		Data: user,
	})
}
