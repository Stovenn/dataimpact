package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/stovenn/dataimpact/internal"
)

var (
	ErrServerClosed error = http.ErrServerClosed
)

type Server struct {
	infoLogger *log.Logger
	errLogger  *log.Logger
	store      internal.Store

	*http.Server
}

func NewServer(us internal.Store, infoLogger, errLogger *log.Logger) *Server {
	s := &Server{
		infoLogger: infoLogger,
		errLogger:  errLogger,

		store: us,
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
	userRouter.HandleFunc("/{id}", server.HandleGet).Methods(http.MethodGet)
	userRouter.HandleFunc("/", server.HandleList).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", server.HandleUpdate).Methods(http.MethodPatch)
	userRouter.HandleFunc("/{id}", server.HandleDelete).Methods(http.MethodDelete)

	server.Handler = r
}
