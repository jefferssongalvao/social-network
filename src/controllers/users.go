package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"social-network/src/authentication"
	"social-network/src/database"
	"social-network/src/models"
	"social-network/src/repositories"
	"social-network/src/responses"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if error := user.Prepare("create"); error != nil {
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
	filter := strings.ToLower(r.URL.Query().Get("user"))

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	users, error := repository.SearchUsers(filter)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, error := strconv.ParseUint(params["id"], 10, 64)
	if error != nil {
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
	user, error := repository.GetUser(ID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, error := strconv.ParseUint(params["id"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	userIdToken, error := authentication.GetUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if userIdToken != ID {
		responses.Error(w, http.StatusForbidden, errors.New("update forbidden for this user"))
		return
	}

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

	if error := user.Prepare("update"); error != nil {
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
	if error := repository.UpdateUser(ID, user); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, error := strconv.ParseUint(params["id"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	userIdToken, error := authentication.GetUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if userIdToken != ID {
		responses.Error(w, http.StatusForbidden, errors.New("delete forbidden for this user"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	if error := repository.DeleteUser(ID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, error := authentication.GetUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)
	userId, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if followerId == userId {
		responses.Error(w, http.StatusForbidden, errors.New("not allowed to follow yourself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	if error := repository.FollowUser(userId, followerId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, error := authentication.GetUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)
	userId, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if userId == followerId {
		responses.Error(w, http.StatusForbidden, errors.New("not allowed to unfollow yourself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	if error := repository.UnfollowUser(userId, followerId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
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
	followers, error := repository.GetFollowers(userId)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, followers)
}

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
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
	users, error := repository.GetFollowing(userId)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}
