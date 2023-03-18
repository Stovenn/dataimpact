package http

import (
	"net/http"
)

func (s *Server) HandleCreate(w http.ResponseWriter, r *http.Request) {
	s.userStore.Create()
}
