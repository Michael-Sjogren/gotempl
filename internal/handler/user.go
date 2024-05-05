package handler

import (
	"log"
	"slices"
	"strconv"

	"github.com/Michael-Sjogren/gotempl/internal/model"
	"github.com/Michael-Sjogren/gotempl/internal/views/components"
	"github.com/Michael-Sjogren/gotempl/internal/views/pages"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserModel *model.UserRepo
}

func (h *UserHandler) HandleUsersPageView(c *fiber.Ctx) error {
	users, err := h.UserModel.GetAll()
	if err != nil {
		log.Println(err)
		return err
	}

	return Render(c, pages.UsersPage(users))
}

func (h *UserHandler) HandleLoginView(c *fiber.Ctx) error {
	return Render(c, pages.LoginPage())
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", -1)

	if err != nil {
		c.Status(400)
		return err
	}
	err = h.UserModel.Delete(id)

	log.Println(err)
	return err
}

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var newUser model.User

	username := c.FormValue("username")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirm-password")
	access := c.FormValue("access")

	errorList := []string{}

	if len(username) == 0 {
		errorList = append(errorList, "Username must be defined")
	}
	newUser.Username = username

	if slices.Equal([]byte(password), []byte(confirmPassword)) {
		errorList = append(errorList, "confirm password and password was not equal")
	}

	if len(password) < 8 {
		errorList = append(errorList, "password must be atleast 8 in length")
	}

	var err error
	if n, err := strconv.Atoi(access); err == nil {
		newUser.Access = n
	} else {
		errorList = append(errorList, "Invalid access value, must be a numeric value "+err.Error())
	}

	hash, err := model.GeneratePassword(password)
	if err != nil {
		errorList = append(errorList, err.Error())
	}
	hasErrors := len(errorList) > 0

	if !hasErrors {

		newUser, err = h.UserModel.CreateUser(newUser, hash)

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
	return Render(c, components.CreateUserForm("", 0, errorList))
}

func (h *UserHandler) HandleUserFormView(c *fiber.Ctx) error {
	return Render(c, components.CreateUserForm("", 0, make([]string, 0)))
}
