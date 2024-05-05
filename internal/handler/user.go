package handler

import (
	"log"
	"net/http"
	"slices"
	"strconv"

	"github.com/Michael-Sjogren/gotempl/internal/model"
	"github.com/Michael-Sjogren/gotempl/internal/views/components"
	"github.com/Michael-Sjogren/gotempl/internal/views/pages"
)

type UserHandler struct {
	UserModel *model.UserRepo
}

func (h *UserHandler) HandleUsersPageView(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserModel.GetAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = pages.UsersPage(users).Render(r.Context(), w)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *UserHandler) HandleLoginView(w http.ResponseWriter, r *http.Request) {
	if err := pages.LoginPage().Render(r.Context(), w); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	if err := r.ParseForm(); err != nil {
		log.Panicln(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	confirmPassword := r.Form.Get("confirm-password")
	access := r.Form.Get("access")

	errorList := []string{}

	if len(username) == 0 {
		errorList = append(errorList, "Username must be defined")
	}
	if slices.Equal([]byte(password), []byte(confirmPassword)) {
		errorList = append(errorList, "confirm password and password was not equal")
	}

	if len(password) < 8 {
		errorList = append(errorList, "password must be atleast 8 in length")
	}

	var err error
	if newUser.Access, err = strconv.Atoi(access); err != nil {
		errorList = append(errorList, "Invalid access value, must be a numeric value.")
	}

	hash, err := model.GeneratePassword(password)
	if err != nil {
		errorList = append(errorList, err.Error())
	}
	hasErrors := len(errorList) > 0

	if !hasErrors {

		newUser, err = h.UserModel.CreateUser(model.User{
			Username: username,
			Access:   0,
		}, hash)

		if err != nil {
			errorList = append(errorList, err.Error())
			hasErrors = true
		} else {
			password = ""
			confirmPassword = ""
			username = ""
			access = ""
		}
	}

	log.Println(errorList)

	err = components.CreateUserForm("", 0, errorList).Render(r.Context(), w)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *UserHandler) HandleUserFormView(w http.ResponseWriter, r *http.Request) {
	err := components.CreateUserForm("", 0, make([]string, 0)).Render(r.Context(), w)
	if err != nil {
		log.Fatal(err)
	}
}
