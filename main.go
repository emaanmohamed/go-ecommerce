package main

import (
	"github.com/ecommerce/controllers"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication()
}
