package userController

import (
	"encoding/json"
	"net/http"

	"github.com/rishavmngo/todo/models/users"
	"github.com/rishavmngo/todo/utils"
)

type ErrorMessage struct{
	Message string `json:"message"`
	Error string `json:"error"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newUser := &userModel.User{}
	utils.ParseBody(r, newUser)
	u, err := newUser.Register()
	if err != nil {
		errResp := ErrorMessage{Message: "Server Error", Error: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		e, _ := json.Marshal(errResp)
		w.Write(e)
		return
	}
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
