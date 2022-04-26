package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"social-network/src/authentication"
	"social-network/src/database"
	"social-network/src/models"
	"social-network/src/repositories"
	"social-network/src/responses"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	authorID, error := authentication.GetUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	request, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var post models.Post
	if error = json.Unmarshal(request, &post); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	post.AuthorID = authorID

	if error = post.Prepare(); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryPosts(db)
	post.ID, error = repository.Create(post)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusCreated, post)
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	userID, error := authentication.GetUserID(r)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	repository := repositories.NewRepositoryPosts(db)
	posts, error := repository.ListPosts(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId, error := strconv.ParseUint(params["id"], 10, 64)
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

	repository := repositories.NewRepositoryPosts(db)
	post, error := repository.GetPost(postId)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
}
