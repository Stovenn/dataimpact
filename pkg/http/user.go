package http

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
	"github.com/stovenn/dataimpact/internal/model"
	"github.com/stovenn/dataimpact/pkg/bcrypt"
)

const NB_WORKER = 4

func (s *Server) HandleCreate(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("DataSet")
	if err != nil {
		s.errLogger.Printf("HandleCreate: %v", err)
		return
	}
	defer file.Close()

	d := json.NewDecoder(bufio.NewReader(file))

	d.Token()
	for d.More() {
		user := &model.User{}
		d.Decode(user)

		go func(u *model.User) {
			found, _ := s.userStore.FindOne(u.ID)
			if found != nil {
				s.infoLogger.Printf("user with id %s exists in database, skipping entry", u.ID)
				return
			}

			hash, err := bcrypt.HashPassword(u.Password)
			if err != nil {
				s.errLogger.Panicf("%v\n", err)
			}
			user.Password = string(hash)
			s.userStore.Create(u)

			workdirPath, _ := os.Getwd()
			f, err := os.Create(path.Join(workdirPath, "data", u.ID))
			if err != nil {
				s.errLogger.Panicf("%v\n", err)
			}
			defer f.Close()
			f.Write([]byte(u.Data))

			s.infoLogger.Printf("created new user with id %s\n", u.ID)
		}(user)
	}
	d.Token()
}

func (s *Server) HandlerGet(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	user, err := s.userStore.FindOne(userID)
	if err != nil {
		s.errLogger.Panicf("%v\n", err)
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}
