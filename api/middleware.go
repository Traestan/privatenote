package api

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

// A Logger function which simply wraps the handler function around some log messages
func (s *Server) Logger(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		s.logger.Info().Str("method", r.Method).Str("path", r.URL.Path).Msg("Start request")
		fn(w, r, param)
	}
}

// CORS
func (s *Server) CORS(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		fn(w, r, param)
	}
}

func (s *Server) AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.Method != "OPTIONS" {
			tokenString := r.Header.Get("Authorization")
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				}

				return []byte(s.config.GetString("JwtKey")), nil
			})
			if err != nil {
				s.logger.Error().Str("err", err.Error()).Msg("Parse token")
				return
			}
			if token.Valid {
				s.logger.Debug().Msg("Parse token")
			} else if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					http.Error(w, "ERROR", http.StatusUnauthorized)
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					http.Error(w, "ERROR", http.StatusUnauthorized)
					return
				} else {
					http.Error(w, "ERROR", http.StatusUnauthorized)
					return
				}
			} else {
				http.Error(w, "ERROR", http.StatusUnauthorized)
			}
			ctx := context.WithValue(r.Context(), "Email", token.Claims)
			next(w, r.WithContext(ctx), ps)
		} else {
			next(w, r, ps)
		}
	}
}
