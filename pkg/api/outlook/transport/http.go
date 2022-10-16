package transport

import (
	"net/http"

	"github.com/DarioCalovic/secretify/pkg/api/outlook"
	"github.com/gorilla/mux"
)

const (
	route = "/outlook"
)

// HTTP represents outlook http service
type HTTP struct {
	svc outlook.Service
}

// NewHTTP creates new outlook http service
func NewHTTP(svc outlook.Service, r *mux.Router) {
	h := HTTP{svc}

	s := r.PathPrefix(route).Subrouter()

	s.HandleFunc("/_manifest", h.manifest).
		Methods("GET")

}

func (h *HTTP) manifest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(h.svc.Manifest())
	//xml.NewEncoder(w).Encode(h.svc.Manifest())
}
