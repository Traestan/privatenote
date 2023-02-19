package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/traestan/privatenote/internal/model"
)

func (s *Server) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userExternal := s.factory.NewUser(&model.User{})
	err := json.NewDecoder(r.Body).Decode(&userExternal)
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
	user, err := userExternal.Login()

	if err != nil {
		s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 500, "Internal Server Error")
		return
	}
	if user.Email == "" {
		s.WriteErrorResponse(w, 400, "user not found")
		return
	}

	expirationTime := time.Now().Add(200 * time.Minute)
	claims := &model.Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.GetString("JwtKey")))
	if err != nil {
		s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 500, "Internal Server Error")
		return
	}
	user.Token = tokenString

	// send response
	resp := map[string]interface{}{
		"user": user,
	}

	s.WriteOKResponse(w, resp)
}
