package main

import (
	"fmt"
	"github.com/feader02/online-shop/internal/api"
	"github.com/feader02/online-shop/internal/db"
	"github.com/feader02/online-shop/internal/utils"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

var (
	Version = ""
)

func main() {
	jwt, err := utils.GenerateJWT("vania")
	fmt.Println(jwt)
	fmt.Printf("App version: %s\n", Version)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		fmt.Println("Port not found")
		port = "8000"
	} else {
		fmt.Printf("Running the server on port %s\n", port)
	}

	mysql := db.NewStorage()
	defer mysql.Close()

	app := api.NewApp(mysql)
	router := app.GetHandle()

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		fmt.Println("Server run error", err)
	}
}
