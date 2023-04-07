package todoroutes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getNotes(w http.ResponseWriter, r *http.Request) {

	res, _ := json.Marshal("todos")
	w.Write(res)
}

func TodoRoutes() *mux.Router {
	router := mux.NewRouter().PathPrefix("/todo").Subrouter()

	router.HandleFunc("/get",getNotes)
	return router
}
