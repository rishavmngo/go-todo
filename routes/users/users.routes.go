package usersroutes

import (
	"github.com/gorilla/mux"
	userController "github.com/rishavmngo/todo/controllers/users"
)

func UserRoute() *mux.Router {

	router := mux.NewRouter().PathPrefix("/users").Subrouter()
	router.HandleFunc("/register", userController.Register).Methods("POST")
	return router

}
