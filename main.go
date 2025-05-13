package main

import (
	"log"
	"net/http"
	config "spoBackend/Config"
	routes "spoBackend/Routes"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/handlers"
)

func main() {
    
    config.ConnectDB()

    r := chi.NewRouter()

 
    r.Mount("/experience", routes.ExperienceRoutes())

   
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),           
        handlers.AllowedMethods([]string{"GET", "POST"}),  
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), 
    )(r)

  
    log.Println("Server running on http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", corsHandler))
}
