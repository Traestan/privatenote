package api

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path"

	"github.com/traestan/privatenote/internal/model"
)

type ResultResponse struct {
	Data  map[string]interface{} `json:"data"`
	Error ApiError               `json:"error"`
}
type ApiError struct {
	Status int
	Title  string
}

// Writes the response as a standard JSON response with StatusOK
func (s *Server) WriteOKResponse(w http.ResponseWriter, m map[string]interface{}) {
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&ResultResponse{Data: m}); err != nil {
		s.WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// Writes the error response as a Standard API JSON response with a response code
func (s *Server) WriteErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.WriteHeader(errorCode)
	json.
		NewEncoder(w).
		Encode(&ResultResponse{Error: ApiError{Status: errorCode, Title: errorMsg}})
}

func (s *Server) WriteTextResponse(w http.ResponseWriter, note model.Note) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fp := path.Join("../web/templates", "note.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
