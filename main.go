package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	todo := Todo{}
	routes := makeRoutes(todo)
	http.ListenAndServe(":"+port, routes)
}
