package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Santiagozh1998/PokedexAPI/database"
	"github.com/Santiagozh1998/PokedexAPI/routes"
	_ "github.com/lib/pq"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	err := database.ConnectDatabase()
	if err != nil {
		fmt.Println("DATABASE CONNECTION FAILED.")
	}

	//test.TestDatabase()

	router := routes.AppRouter()
	fmt.Println("Server running in 	port: http://localhost:" + port)
	http.ListenAndServe(":"+port, router)
}
