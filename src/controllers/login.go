package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"social-network/src/authentication"
	"social-network/src/database"
	"social-network/src/models"
	"social-network/src/repositories"
	"social-network/src/responses"
	"social-network/src/security"
)

func Login(w http.ResponseWriter, r *http.Request) {
	request, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error := json.Unmarshal(request, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	userDatabase, error := repository.GetUserForEmail(user.Email)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	if error := security.CheckPassword(userDatabase.Password, user.Password); error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	token, error := authentication.GenerateToken(userDatabase.ID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	w.Write([]byte(token))
}
