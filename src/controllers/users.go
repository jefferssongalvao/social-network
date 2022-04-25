package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User!"))
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
