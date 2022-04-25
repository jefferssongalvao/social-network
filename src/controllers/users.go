package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"social-network/src/database"
	"social-network/src/models"
	"social-network/src/repositories"
	"social-network/src/responses"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if error := user.Prepare(); error != nil {
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
	user.ID, error = repository.Create(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}

	responses.JSON(w, http.StatusCreated, user)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Users!"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User!"))
}
