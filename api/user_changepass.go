package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/traestan/privatenote/internal/model"
)

func (s *Server) changepass(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var response map[string]interface{}
	userModel := s.factory.NewUser(&model.User{})
	err := json.NewDecoder(r.Body).Decode(&userModel)
	switch {
	case err == io.EOF:
		s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 500, "Internal Server Error")
		return
	case err != nil:
		s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 500, "Internal Server Error")
		return
	}

	user := r.Context().Value("Email").(jwt.MapClaims)
	userEmail := user["email"].(string)
	userModel.Email = userEmail

	if userModel.NewPassword != userModel.NewPassCompare {
		s.WriteErrorResponse(w, 200, "newpass != newpasscompare")
		return
	}

	userNewPass, err := userModel.Changepass()
	if err != nil {
		s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 500, "Internal Server Error")
		return
	}
	if userNewPass == nil {
		s.WriteErrorResponse(w, 400, "user not found")
		return
	}
	response = map[string]interface{}{
		"user": "success change password",
	}
	s.WriteOKResponse(w, response)
	return
}
