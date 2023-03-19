package http

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
	"github.com/stovenn/dataimpact/internal/model"
	"github.com/stovenn/dataimpact/pkg/bcrypt"
	"github.com/stovenn/dataimpact/pkg/util"
)

func (s *Server) HandleCreate(w http.ResponseWriter, r *http.Request) {
	parentContext := context.Background()

	file, err := os.Open("DataSet")
	if err != nil {
		s.errLogger.Printf("%v", err)
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	defer file.Close()

	d := json.NewDecoder(bufio.NewReader(file))

	d.Token()
	for d.More() {
		user := &model.User{}
		d.Decode(user)

		go func(ctx context.Context, u *model.User) {
			found, _ := s.store.FindOne(ctx, u.ID)
			if found != nil {
				s.infoLogger.Printf("user with id %s exists in database, skipping entry", u.ID)
				return
			}

			hash, err := bcrypt.HashPassword(*u.Password)
			if err != nil {
				s.errLogger.Panicf("%v\n", err)
				handleError(w, http.StatusInternalServerError, err)
				return
			}
			*user.Password = string(hash)

			err = s.store.Create(ctx, u)
			if err != nil {
				s.errLogger.Panicf("%v\n", err)
				handleError(w, http.StatusInternalServerError, err)
				return
			}

			err = writeData(u.ID, *u.Data)
			if err != nil {
				s.errLogger.Panicf("%v\n", err)
				handleError(w, http.StatusInternalServerError, err)
				return
			}

			s.infoLogger.Printf("successfully created user %s\n", u.ID)
		}(parentContext, user)
	}
	d.Token()
}

func (s *Server) HandleGet(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	user, err := s.store.FindOne(context.Background(), userID)
	if err != nil {
		s.errLogger.Panicf("%v\n", err)
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(model.ToResponse(user))
}

func (s *Server) HandleList(w http.ResponseWriter, r *http.Request) {
	users, err := s.store.Find(context.Background())
	if err != nil {
		s.errLogger.Panicf("%v\n", err)
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	var userResponse []*model.UserResponse
	for _, u := range users {
		fmt.Printf("%+v", u)
		userResponse = append(userResponse, model.ToResponse(u))
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(userResponse)
}

func (s *Server) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	err := checkAccessRight(userID, r)
	if err != nil {
		s.errLogger.Printf("%v\n", err)
		handleError(w, http.StatusUnauthorized, err)
		return
	}

	user, err := s.store.FindOne(context.Background(), userID)
	if err != nil {
		s.errLogger.Printf("%v\n", err)
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	var req *model.User
	json.NewDecoder(r.Body).Decode(&req)

	update := buildUpdate(user, req)

	err = s.store.Update(context.Background(), userID, update)
	if err != nil {
		s.errLogger.Panicf("%v\n", err)
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	if update.Data != user.Data {
		writeData(userID, *update.Data)
	}

	updatedUser, _ := s.store.FindOne(context.Background(), userID)

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(model.ToResponse(updatedUser))
}

func buildUpdate(old, new *model.User) *model.User {
	update := *old
	if new.Password != nil {
		update.Password = new.Password
	}
	if new.IsActive != nil {
		update.IsActive = new.IsActive
	}
	if new.Balance != nil {
		update.Balance = new.Balance
	}
	if new.Age != nil {
		update.Age = new.Age
	}
	if new.Name != nil {
		update.Name = new.Name
	}
	if new.Gender != nil {
		update.Gender = new.Gender
	}
	if new.Company != nil {
		update.Company = new.Company
	}
	if new.Email != nil {
		update.Email = new.Email
	}
	if new.Phone != nil {
		update.Phone = new.Phone
	}
	if new.Address != nil {
		update.Address = new.Address
	}
	if new.About != nil {
		update.About = new.About
	}
	if new.Registered != nil {
		update.Registered = new.Registered
	}
	if new.Latitude != nil {
		update.Latitude = new.Latitude
	}
	if new.Longitude != nil {
		update.Longitude = new.Longitude
	}
	if new.Tags != nil {
		update.Tags = new.Tags
	}
	if new.Friends != nil {
		update.Friends = new.Friends
	}
	if new.Data != nil {
		update.Data = new.Data
	}

	return &update
}

func (s *Server) HandleDelete(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	err := checkAccessRight(userID, r)
	if err != nil {
		s.errLogger.Printf("%v\n", err)
		handleError(w, http.StatusUnauthorized, err)
		return
	}

	err = s.store.DeleteOne(context.Background(), userID)
	if err != nil {
		s.errLogger.Panicf("%v\n", err)
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	err = deleteData(userID)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			s.infoLogger.Printf("no file found for id %s", userID)
			return
		}
		s.errLogger.Panicf("%v\n", err)
		handleError(w, http.StatusInternalServerError, err)
		return
	}
}

func writeData(id, data string) error {
	util.CreateDataDirIfNotExists()
	workdirPath, _ := os.Getwd()
	f, err := os.Create(path.Join(workdirPath, "data", id))
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write([]byte(data))

	return nil
}

func deleteData(id string) error {
	workdirPath, _ := os.Getwd()
	filePath := path.Join(workdirPath, "data", id)
	_, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
