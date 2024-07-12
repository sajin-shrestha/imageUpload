package routers

import (
	"github.com/gorilla/mux"
	"github.com/sajin-shrestha/imageUpload/controllers"
	"github.com/sajin-shrestha/imageUpload/utils"
)

func InitRouter() *mux.Router {
	// router := mux.NewRouter()

	// router.Use(utils.CORSMiddleware) // Apply the CORS middleware

	// router.HandleFunc("/testimonials", controllers.GetTestimonials).Methods("GET")
	// router.HandleFunc("/testimonials/{id}", controllers.GetTestimonial).Methods("GET")
	// router.HandleFunc("/testimonials", controllers.CreateTestimonial).Methods("POST")
	// router.HandleFunc("/testimonials/{id}", controllers.UpdateTestimonial).Methods("PUT")
	// router.HandleFunc("/testimonials/{id}", controllers.DeleteTestimonial).Methods("DELETE")

	// return router

	router := mux.NewRouter()

	router.Use(utils.CORSMiddleware) // Apply the CORS middleware

	router.HandleFunc("/testimonials", controllers.GetTestimonials).Methods("GET")
	router.HandleFunc("/testimonials/{id}", controllers.GetTestimonial).Methods("GET")
	router.HandleFunc("/testimonials", controllers.CreateTestimonial).Methods("POST")
	router.HandleFunc("/testimonials/{id}", controllers.UpdateTestimonial).Methods("PUT")
	router.HandleFunc("/testimonials/{id}", controllers.DeleteTestimonial).Methods("DELETE")

	return router
}
