package main

import (
	"log"
	"net/http"
	config "spoBackend/Config"
	routes "spoBackend/Routes"

	"github.com/go-chi/chi/v5"
)

func main() {
    config.ConnectDB()

    r := chi.NewRouter()

    r.Mount("/experience", routes.ExperienceRoutes())

    log.Println("Server running on http://localhost:8000")
    http.ListenAndServe(":8000", r)
}
