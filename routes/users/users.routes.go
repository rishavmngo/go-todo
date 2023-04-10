package usersroutes

import (
	"github.com/gorilla/mux"
	"github.com/rishavmngo/todo/controllers/users"
	"github.com/rishavmngo/todo/middlewares"
)

func UserRoute() *mux.Router {

	router := mux.NewRouter().PathPrefix("/users").Subrouter()
	router.HandleFunc("/register", middlewares.SetMiddlewareJson(userController.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.SetMiddlewareJson(userController.Login)).Methods("POST")
	return router

}
