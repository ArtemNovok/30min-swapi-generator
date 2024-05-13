package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
}

const webPort = "8000"

func main() {
	log.Println("Starting server...")
	app := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
