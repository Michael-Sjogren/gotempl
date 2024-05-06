package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Michael-Sjogren/gotempl/internal/db"
	"github.com/Michael-Sjogren/gotempl/internal/handler"
	"github.com/Michael-Sjogren/gotempl/internal/model"
	"github.com/Michael-Sjogren/gotempl/internal/session"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	con, err := db.Connect(dir + "/mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	err = db.InitializeDB(con)

	if err != nil {
		log.Fatal(err)
	}
	bytes, err := model.GeneratePassword("root")
	if err != nil {
		log.Fatal(err)
	}

	userModel := model.NewUserRepo(con)
	sessionManager := session.NewSessionManager()
	_, err = userModel.CreateUser(model.User{
		Username: "Michael",
		Access:   -1,
	}, bytes)

	if err != nil {
		log.Println(err)
	}

	defer con.Close()

	ip := "127.0.0.1:3333"
	app := fiber.New()

	home := handler.HomeHandler{}
	users := handler.UserHandler{UserModel: &userModel, SessionManager: sessionManager}
	directoryPath := "./static"
	_, err = os.Stat(directoryPath)
	if os.IsNotExist(err) {
		fmt.Printf("Directory '%s' not found.\n", directoryPath)
		return
	}

	hxRouter := app.Group("/hx")
	app.Use(logger.New())
	app.Static("/static/", directoryPath)
	app.Get("/", home.HandlerHomePageView)
	app.Get("/users", users.HandleUsersPageView)

	app.Get("/login", users.HandleLoginView)
	app.Post("/hx/login", users.HandleLogin)

	// htmx handlers
	hxRouter.Post("/users", users.HandleCreateUser)
	hxRouter.Get("/users", users.HandleUserFormView)
	hxRouter.Delete("/users/:id", users.HandleDeleteUser)

	log.Printf("starting server on: http://%s\n", ip)
	if err := app.Listen(ip); err != nil {
		log.Fatal(err)
	}

}
