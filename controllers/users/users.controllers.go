package userController

import (
	"net/http"

	"github.com/rishavmngo/todo/models/users"
	"github.com/rishavmngo/todo/responses"
	"github.com/rishavmngo/todo/token"
	"github.com/rishavmngo/todo/utils"
)

type ErrorMessage struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	newUser := &userModel.User{}
	utils.ParseBody(r, newUser)
	user, err := newUser.Register()
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
	}
	response.JSON(w, http.StatusOK, user)

}

func Login(w http.ResponseWriter, r *http.Request) {
	user := &userModel.User{}
	utils.ParseBody(r, user)
	err := user.Validate("login")

	resp, err := user.GetByEmail()

	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tokenString, err := token.Createtoken(resp.ID)

	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, tokenString)
}
