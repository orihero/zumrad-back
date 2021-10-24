package main

import (
	"fmt"
	"log"
	"net/http"
	"zumrad/database"
	"zumrad/env"
	"zumrad/router"
)

func StartServer() {
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", env.Router)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	database.InitialMigration()
	router.NewRouter()
	StartServer()
}
