package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/stovenn/dataimpact/pkg/mongo"
)

var (
	ErrServerClosed error = http.ErrServerClosed
)

type Server struct {
	infoLogger *log.Logger
	errLogger  *log.Logger
	userStore  mongo.UserStore

	*http.Server
}

func NewServer(us mongo.UserStore, infoLogger, errLogger *log.Logger) *Server {
	s := &Server{
		infoLogger: infoLogger,
		errLogger:  errLogger,

		userStore: us,
		Server: &http.Server{
			Addr:         ":8080",
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
	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", server.HandleCreate).Methods(http.MethodPost)
	userRouter.HandleFunc("/", nil).Methods(http.MethodGet)
	userRouter.HandleFunc("/:id", nil).Methods(http.MethodGet)
	userRouter.HandleFunc("/:id", nil).Methods(http.MethodPut)
	userRouter.HandleFunc("/:id", nil).Methods(http.MethodDelete)

	server.Handler = r
}
