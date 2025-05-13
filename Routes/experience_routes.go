package routes

import (
	"net/http"
	controllers "spoBackend/Controlers"

	"github.com/go-chi/chi/v5"
)

func ExperienceRoutes() http.Handler {
    r := chi.NewRouter()

    r.Post("/", controllers.SaveExperience )       
    r.Get("/", controllers.GetAllExperiences)     
    r.Get("/stats", controllers.GetStats)         

    return r
}
