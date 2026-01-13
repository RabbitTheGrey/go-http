package main

import (
	"database/sql"
	"fmt"
	"go-web/conf"
	"go-web/internal/handler/auth_handler"
	"go-web/internal/handler/post_handler"
	"go-web/pkg/db"
	"go-web/pkg/helper"
	"go-web/pkg/router"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := sql.Open(conf.DbDriver, conf.DbConnectionString)
	helper.HandlePanic(err)
	db.SetDB(database)
	defer database.Close()

	router := router.Router{}
	router.Post("/auth/login", auth_handler.Login)

	router.Prefix("/api/post")
	router.Get("", post_handler.List)
	router.Get("/show", post_handler.Read)
	router.Post("", post_handler.Create)
	router.Patch("", post_handler.Update)
	router.Delete("", post_handler.Delete)

	fmt.Println("Listening...")
	err = http.ListenAndServe(conf.HttpHost+":"+conf.HttpPort, &router)

	if err != nil {
		panic(err)
	}
}
