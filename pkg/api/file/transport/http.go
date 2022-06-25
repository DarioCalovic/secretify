package transport

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/DarioCalovic/secretify"
	"github.com/DarioCalovic/secretify/pkg/api/file"
	"github.com/gorilla/mux"
)

const (
	route = "/file"
)

// HTTP represents user http service
type HTTP struct {
	svc file.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc file.Service, r *mux.Router) {
	h := HTTP{svc}

	s := r.PathPrefix(route).Subrouter()

	s.HandleFunc("", h.upload).
		Methods("POST")

	s.HandleFunc("/{id}", h.download).
		Methods("GET")

}

type uploadReq struct {
	Filename string `json:"filename"`
	Type     string `json:"type"`
	Size     uint   `json:"size"`
}

type uploadRes struct {
	CreatedAt  time.Time `json:"created_at"`
	Identifier string    `json:"identifier"`
}

func (h *HTTP) upload(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json; charset=utf-8")
	var req = &uploadReq{
		Filename: r.URL.Query().Get("filename"),
		Type:     r.URL.Query().Get("type"),
	}
	if sz, err := strconv.ParseUint(r.URL.Query().Get("size"), 10, 32); err == nil {
		req.Size = uint(sz)
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: errors.New("could not retrieve file").Error(),
		})
		return
	}

	// Validation
	{
		// Meta
		// TODO: handle file type validation
		if req.Filename == "" || handler.Size == 0 {
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
				Error: errors.New("not all file metadata provided").Error(),
			})
			return
		}
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: errors.New("could not read file content").Error(),
		})
	}
	if len(fileBytes) == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: errors.New("no file content").Error(),
		})
		return
	}

	// Create File
	fle, err := h.svc.Create(fileBytes, req.Filename, req.Type, req.Size)
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
		Data: uploadRes{
			CreatedAt:  fle.CreatedAt,
			Identifier: fle.Identifier,
		},
	})
}

type downloadRes struct {
	Cipher        string    `json:"cipher"`
	ExpiresAt     time.Time `json:"expires_at"`
	HasPassphrase bool      `json:"has_passphrase"`
	RevealOnce    bool      `json:"reveal_once"`
	Deleted       bool      `json:"deleted"`
}

func (h *HTTP) download(w http.ResponseWriter, r *http.Request) {
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

	file, err := h.svc.View(identifier)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: err.Error(),
		})
		return
	}

	content, err := h.svc.Read(file.Path)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(secretify.HTTPErrorResponse{
			Error: err.Error(),
		})
		return
	}

	mime := http.DetectContentType(content)

	// Generate the server headers
	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name))
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.FormatUint(uint64(file.Size), 10))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	http.ServeContent(w, r, file.Name, time.Now(), bytes.NewReader(content))
}
