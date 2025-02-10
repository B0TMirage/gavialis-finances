package main

import (
	"fmt"
	"net/http"

	"github.com/B0TMirage/gavialis-finances/pkg/database"
	"github.com/B0TMirage/gavialis-finances/pkg/routes"
)

func main() {
	database.Connect()
	if err := database.DB.Ping(); err != nil {
		fmt.Println(err)
	}
	database.Migrate()
	defer database.DB.Close()

	mux := routes.Regrouter()

	fmt.Println("Server started on PORT 8000.")
	http.ListenAndServe(":8000", mux)
}
