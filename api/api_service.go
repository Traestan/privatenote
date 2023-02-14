package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

//ttl storage time of records in redis
func (s *Server) ttl(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"300s":   "5 min",
		"600s":   "10 min",
		"3600s":  "1 hour",
		"86400s": "1 day",
	}
	s.WriteOKResponse(w, response)
}
