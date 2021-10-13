package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"webapp/database"

	"github.com/gorilla/mux"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("<h2>Welcome to HomePage</h2>")
}

func GetUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database.GetUserData())
}
func GetUserByIdController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	json.NewEncoder(w).Encode(database.GetUserById(id))
}
func CreateUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user database.UserData
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Println("Error", err)
	}
	json.NewEncoder(w).Encode(database.CreateUserData(user))
}
func UpdateUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user database.UserData
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Println("Error", err)
	}
	json.NewEncoder(w).Encode(database.UpdateUserData(user))
}
func DeleteUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user database.UserData
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Println("Error", err)
	}
	json.NewEncoder(w).Encode(database.DeleteUserData(user))
}
