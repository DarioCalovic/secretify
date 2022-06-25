package transport

import (
	"encoding/json"
	"net/http"

	"github.com/DarioCalovic/secretify"
	"github.com/DarioCalovic/secretify/pkg/api/setting"
	"github.com/gorilla/mux"
)

const (
	route = "/setting"
)

// HTTP represents setting http service
type HTTP struct {
	svc setting.Service
}

// NewHTTP creates new setting http service
func NewHTTP(svc setting.Service, r *mux.Router) {
	h := HTTP{svc}

	s := r.PathPrefix(route).Subrouter()

	s.HandleFunc("/_meta", h.meta).
		Methods("GET")
	s.HandleFunc("/_policy", h.policy).
		Methods("GET")

}

func (h *HTTP) meta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(secretify.HTTPOKResponse{
		Data: h.svc.Meta(),
	})
}

func (h *HTTP) policy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(secretify.HTTPOKResponse{
		Data: h.svc.Policy(),
	})
}
