package main

import (
	"database/sql"
	"fmt"
	"go-web/conf"
	"go-web/internal/service"
	"go-web/pkg/db"
	"go-web/pkg/helper"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("*** Admin panel ***")

	database, err := sql.Open(conf.DbDriver, conf.DbConnectionString)
	helper.HandlePanic(err)
	db.SetDB(database)
	defer database.Close()

	for {
		var option uint8

		fmt.Println("Select option\n 1) create user\n 0) exit")
		fmt.Scan(&option)

		switch option {
		case 1:
			var login string
			var password string

			fmt.Print("Login: ")
			fmt.Scanf("%s", &login)

			fmt.Print("Password: ")
			fmt.Scanf("%s", &password)
			service.NewUserService().CreateUser(login, password)
		case 0:
			return
		default:
			fmt.Println("Unknown option.")
		}
	}
}
