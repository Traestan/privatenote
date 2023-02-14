package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/privatenote/internal/model"
)

func (s *Server) register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var response map[string]interface{}
	userExternal := s.factory.NewUser(&model.User{})
	err := json.NewDecoder(r.Body).Decode(&userExternal)
	switch {
	case err == io.EOF:
		s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 500, "Internal Server Error")
	case err != nil:
		s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 500, "Internal Server Error")
	}

	user, err := userExternal.Registration()

	if err != nil {
		s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 500, "Internal Server Error")
		return
	}
	if user == nil {
		//s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 400, "user not found")
		return
	}

	if user.Email == "" {
		response = map[string]interface{}{
			"login": "user not found",
		}
		s.WriteOKResponse(w, response)
		return
	}

	response = map[string]interface{}{
		"user": "success registration",
	}
	s.WriteOKResponse(w, response)
	return
}
