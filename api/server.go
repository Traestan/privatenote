package api

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"github.com/traestan/privatenote/internal/model"
)

// NewRouter returns default httprouter.Router
func NewRouter() *httprouter.Router {
	r := httprouter.New()
	r.HandleOPTIONS = true
	r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Content-Type", "application/json; charset=UTF-8")
			header.Set("Access-Control-Allow-Origin", "*")
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			header.Set("Access-Control-Allow-Credentials", "false")
		}
		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
	r.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write([]byte("Method is not allowed"))
	})

	return r
}

// Server is struct for http server
type Server struct {
	http.Server
	logger  *zerolog.Logger
	engine  *httprouter.Router
	config  *viper.Viper
	factory *model.Factory
}

func NewServer(logger *zerolog.Logger, engine *httprouter.Router, config *viper.Viper, factory *model.Factory) *Server {
	s := &Server{
		Server: http.Server{
			Addr:         config.GetString("host"),
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		},
		logger:  logger,
		engine:  engine,
		config:  config,
		factory: factory,
	}
	s.handlerInit()
	s.Handler = engine
	logger.Debug().Msg("Server http start")
	return s
}

func (s *Server) handlerInit() {
	s.logger.Debug().Msg("Server http routers init")
	//service
	s.engine.Handle("GET", "/health", s.Logger(s.health))
	s.engine.Handle("GET", "/service/ttl", s.CORS(s.Logger(s.ttl)))
	// user
	s.engine.Handle("POST", "/user/login", s.CORS(s.login))
	s.engine.Handle("POST", "/user/register", s.CORS(s.register))
	s.engine.Handle("POST", "/user/changepass", s.CORS(s.AuthMiddleware(s.changepass)))
	// note
	s.engine.Handle("POST", "/note/create", s.CORS(s.AuthMiddleware(s.create)))
	s.engine.Handle("GET", "/note/list", s.CORS(s.AuthMiddleware(s.list)))
	s.engine.Handle("GET", "/note/get/:shorturl", s.CORS(s.AuthMiddleware(s.get)))
	s.engine.Handle("GET", "/url/:shorturl", s.view)
	s.engine.Handle("PUT", "/note/edit/:shorturl", s.CORS(s.AuthMiddleware(s.edit)))

}
