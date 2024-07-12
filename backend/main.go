package main

import (
	"log"
	"net/http"

	"github.com/sajin-shrestha/imageUpload/models"
	"github.com/sajin-shrestha/imageUpload/routers"
	"github.com/sajin-shrestha/imageUpload/utils"
)

func main() {
	utils.InitDB()
	utils.DB.AutoMigrate(&models.Testimonial{})

	router := routers.InitRouter()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
