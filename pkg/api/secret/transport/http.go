package transport

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/DarioCalovic/secretify"
	"github.com/DarioCalovic/secretify/pkg/api/secret"
	"github.com/gorilla/mux"
)

const (
	route = "/secret"
)

// HTTP represents user http service
type HTTP struct {
	svc secret.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc secret.Service, r *mux.Router) {
	h := HTTP{svc}

	s := r.PathPrefix(route).Subrouter()

	s.HandleFunc("", h.create).
		Methods("POST")

	s.HandleFunc("/{id}", h.view).
		Methods("GET")

	s.HandleFunc("/{id}/_cipher", h.cipher).
		Methods("GET")

	s.HandleFunc("/{id}", h.delete).
		Methods("DELETE")

}

type createReq struct {
	Cipher         string `json:"cipher"`
	ExpiresAt      string `json:"expires_at"`
	RevealOnce     bool   `json:"reveal_once"`
	DestroyManual  bool   `json:"destroy_manual"`
	HasPassphrase  bool   `json:"has_passphrase"`
	FileIdentifier string `json:"file_identifier"`
	Email          string `json:"email"`
	WebhookAddr    string `json:"webhook_addr"`
}

type createRes struct {
	CreatedAt  time.Time `json:"created_at"`
	Identifier string    `json:"identifier"`
}

func (h *HTTP) create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	var cr createReq
	err := json.NewDecoder(r.Body).Decode(&cr)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// Validation
	if cr.Cipher == "" && cr.FileIdentifier == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: errors.New("no cipher nor file provided").Error(),
		})
		return
	}

	if h.svc.ServiceConfig().Policy().Passphrase.Required && !cr.HasPassphrase {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: errors.New("cipher needs to be encrypted with passphrase as well").Error(),
		})
		return
	}
	var expires time.Time
	{
		et, err := time.ParseDuration(cr.ExpiresAt)
		if err == nil {
			expires = time.Now().Add(et)
		} else {
			i, err := strconv.ParseInt(cr.ExpiresAt, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusNotAcceptable)
				json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
					Error: err.Error(),
				})
				return
			}
			expires = time.Unix(i, 0)
		}
		// check if date is in future
		if expires.Before(time.Now()) {
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
				Error: errors.New("expires_at needs to be in future").Error(),
			})
			return
		}
	}

	// Create secret
	var (
		scr secretify.Secret
	)
	if cr.FileIdentifier == "" {
		scr, err = h.svc.Create(cr.Cipher, cr.HasPassphrase, expires, cr.RevealOnce, cr.DestroyManual, 0, cr.Email, cr.WebhookAddr)
	} else {
		scr, err = h.svc.CreateWithFile(cr.Cipher, cr.HasPassphrase, expires, cr.RevealOnce, cr.DestroyManual, cr.FileIdentifier, cr.Email, cr.WebhookAddr)
	}
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// Response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(secretify.HTTPOKResponse{
		Data: createRes{
			CreatedAt:  scr.CreatedAt,
			Identifier: scr.Identifier,
		},
	})
}

type viewRes struct {
	Cipher        string    `json:"cipher"`
	ExpiresAt     time.Time `json:"expires_at"`
	HasPassphrase bool      `json:"has_passphrase"`
	RevealOnce    bool      `json:"reveal_once"`
	DestroyManual bool      `json:"destroy_manual"`
	Deleted       bool      `json:"deleted"`
	File          struct {
		Identifier string `json:"identifier"`
		Filename   string `json:"filename"`
		Type       string `json:"type"`
		Size       uint   `json:"size"`
	} `json:"file"`
}

func (h *HTTP) view(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	vars := mux.Vars(r)
	identifier := vars["id"]

	if identifier == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: errors.New("no identifier provided").Error(),
		})
		return
	}

	secret, deleted, err := h.svc.View(identifier, true)
	if err != nil {
		if err == secretify.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)

		} else {
			w.WriteHeader(http.StatusNotAcceptable)
		}
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(secretify.HTTPOKResponse{
		Data: viewRes{
			ExpiresAt:     secret.ExpiresAt,
			HasPassphrase: secret.HasPassphrase,
			RevealOnce:    secret.RevealOnce,
			DestroyManual: secret.DestroyManual,
			Deleted:       deleted,
			File: struct {
				Identifier string `json:"identifier"`
				Filename   string `json:"filename"`
				Type       string `json:"type"`
				Size       uint   `json:"size"`
			}{
				Identifier: secret.File.Identifier,
				Filename:   secret.File.Name,
				Type:       secret.File.Type,
				Size:       secret.File.Size,
			},
		},
	})
}

func (h *HTTP) cipher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	vars := mux.Vars(r)
	identifier := vars["id"]

	if identifier == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: errors.New("no identifier provided").Error(),
		})
		return
	}

	secret, deleted, err := h.svc.View(identifier, false)
	if err != nil {
		if err == secretify.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)

		} else {
			w.WriteHeader(http.StatusNotAcceptable)
		}
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(secretify.HTTPOKResponse{
		Data: viewRes{
			Cipher:        secret.Cipher,
			ExpiresAt:     secret.ExpiresAt,
			HasPassphrase: secret.HasPassphrase,
			RevealOnce:    secret.RevealOnce,
			Deleted:       deleted,
		},
	})
}

func (h *HTTP) delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	vars := mux.Vars(r)
	identifier := vars["id"]

	if identifier == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: errors.New("no identifier provided").Error(),
		})
		return
	}

	err := h.svc.Delete(identifier)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
}
