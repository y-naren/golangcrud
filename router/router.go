/*
Router for the web application
*/
package router

import (
	"log"
	"net/http"
	"webapp/controller"

	"github.com/gorilla/mux"
)

func InitializeRouter() {
	//declared new mux router
	router := mux.NewRouter()
	//registering endpoints to router
	router.HandleFunc("/", controller.HomeController).Methods(http.MethodGet)                 //GET
	router.HandleFunc("/user", controller.GetUserController).Methods(http.MethodGet)          //GET
	router.HandleFunc("/user/{id:int}", controller.GetUserController).Methods(http.MethodGet) //GET
	router.HandleFunc("/user", controller.CreateUserController).Methods(http.MethodPost)      //POST
	router.HandleFunc("/user", controller.UpdateUserController).Methods(http.MethodPatch)     //PATCH
	router.HandleFunc("/user", controller.DeleteUserController).Methods(http.MethodDelete)    //DELETE

	log.Println("Server Started on port localhost:1112")
	log.Fatal(http.ListenAndServe(":1112", router))
}
