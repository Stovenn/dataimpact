package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/stovenn/dataimpact/internal"
	"github.com/stovenn/dataimpact/pkg/token"
	"github.com/stovenn/dataimpact/pkg/util"
)

var (
	ErrServerClosed error = http.ErrServerClosed
)

type Server struct {
	infoLogger *log.Logger
	errLogger  *log.Logger
	store      internal.Store
	config     util.Config
	tokenMaker token.Maker

	*http.Server
}

func NewServer(us internal.Store, infoLogger, errLogger *log.Logger, config util.Config) *Server {
	tokenMaker, err := token.NewJWTMaker(config.SymmetricKey)
	if err != nil {
		errLogger.Fatalf("%v", err)
	}

	s := &Server{
		infoLogger: infoLogger,
		errLogger:  errLogger,
		tokenMaker: tokenMaker,
		store:      us,
		config:     config,
		Server: &http.Server{
			Addr:         fmt.Sprintf(":%s", config.Port),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  90 * time.Second,
		},
	}
	s.setupRoutes()

	return s
}

func (server *Server) setupRoutes() {
	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	r.HandleFunc("/login", server.HandleLogin).Methods(http.MethodPost)
	r.HandleFunc("/users/", server.HandleCreate).Methods(http.MethodPost)

	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.Use(authMiddleware(server.tokenMaker))
	userRouter.HandleFunc("/{id}", server.HandleGet).Methods(http.MethodGet)
	userRouter.HandleFunc("/", server.HandleList).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", server.HandleUpdate).Methods(http.MethodPatch)
	userRouter.HandleFunc("/{id}", server.HandleDelete).Methods(http.MethodDelete)

	server.Handler = r
}

func handleError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}
