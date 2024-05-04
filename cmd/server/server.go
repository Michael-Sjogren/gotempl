package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Michael-Sjogren/gotempl/db"
	"github.com/Michael-Sjogren/gotempl/handler"
	"github.com/Michael-Sjogren/gotempl/model"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(dir)
	mux := http.NewServeMux()
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

	userModel := model.UserModel{Con: con}
	_, err = userModel.CreateUser(model.User{
		Username: "Michael",
		Access:   -1,
	}, bytes)

	if err != nil {
		log.Println(err)
	}

	defer con.Close()

	home := handler.HomeHandler{}
	users := handler.UserHandler{}
	// Replace "." with the actual path of the directory you want to expose.
	directoryPath := "./static"

	// Check if the directory exists
	_, err = os.Stat(directoryPath)
	if os.IsNotExist(err) {
		fmt.Printf("Directory '%s' not found.\n", directoryPath)
		return
	}
	// handle main page routes
	mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(directoryPath))))
	mux.HandleFunc("/", home.HandlerHomePageView)
	mux.HandleFunc("/users", users.HandleUsersPageView)
	mux.HandleFunc("/login", users.HandleUsersPageView)

	// htmx handlers

	ip := "127.0.0.1:3333"
	log.Printf("starting server on: http://%s\n", ip)
	if err := http.ListenAndServe(ip, mux); err != nil {
		log.Fatal(err)
	}
}
