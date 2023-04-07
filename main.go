package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rishavmngo/todo/routes/users"
	"github.com/rishavmngo/todo/routes/todo"
)


func main() {
router := mux.NewRouter()

	router.PathPrefix("/users").Handler(usersroutes.UserRoute())
	router.PathPrefix("/todo").Handler(todoroutes.TodoRoutes())

	log.Fatal(http.ListenAndServe("localhost:3000",router))


}
