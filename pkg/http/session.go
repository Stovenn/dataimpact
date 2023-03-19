package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/stovenn/dataimpact/internal/model"
	"github.com/stovenn/dataimpact/pkg/bcrypt"
	"github.com/stovenn/dataimpact/pkg/token"
)

func (server *Server) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var credentials model.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	foundUser, err := server.store.FindOne(context.Background(), credentials.ID)
	if err != nil {
		handleError(w, http.StatusNotFound, err)
		return
	}

	err = bcrypt.CheckPassword(credentials.Password, *foundUser.Password)
	if err != nil {
		handleError(w, http.StatusUnauthorized, err)
		return
	}

	token, err := server.tokenMaker.CreateToken(credentials.ID, server.config.TokenDuration)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	response := struct {
		AccessToken string
		User        *model.User
	}{
		AccessToken: token,
		User:        foundUser,
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func checkAccessRight(userID string, r *http.Request) error {
	authPayload := r.Context().Value(authPayloadKey).(*token.Payload)
	if authPayload.UserID != userID {
		return errors.New("cannot access ressource")
	}
	return nil
}
