package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sajin-shrestha/imageUpload/models"
	"github.com/sajin-shrestha/imageUpload/utils"
)

// func GetTestimonials(w http.ResponseWriter, r *http.Request) {
// 	var testimonials []models.Testimonial
// 	utils.DB.Find(&testimonials)
// 	json.NewEncoder(w).Encode(testimonials)
// }

// func GetTestimonial(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id, _ := strconv.Atoi(params["id"])
// 	var testimonial models.Testimonial
// 	utils.DB.First(&testimonial, id)
// 	json.NewEncoder(w).Encode(testimonial)
// }

// func CreateTestimonial(w http.ResponseWriter, r *http.Request) {
// 	title := r.FormValue("title")
// 	description := r.FormValue("description")
// 	file, handler, err := r.FormFile("image")
// 	if err != nil {
// 		http.Error(w, "Invalid image file", http.StatusBadRequest)
// 		return
// 	}

// 	imagePath, err := utils.SaveFile(file, handler)
// 	if err != nil {
// 		http.Error(w, "Failed to save image", http.StatusInternalServerError)
// 		return
// 	}

// 	testimonial := models.Testimonial{
// 		Title:       title,
// 		Description: description,
// 		Image:       imagePath,
// 	}

// 	utils.DB.Create(&testimonial)
// 	json.NewEncoder(w).Encode(testimonial)
// }

// func UpdateTestimonial(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id, _ := strconv.Atoi(params["id"])
// 	var testimonial models.Testimonial
// 	utils.DB.First(&testimonial, id)

// 	title := r.FormValue("title")
// 	description := r.FormValue("description")
// 	testimonial.Title = title
// 	testimonial.Description = description

// 	file, handler, err := r.FormFile("image")
// 	if err == nil {
// 		imagePath, err := utils.SaveFile(file, handler)
// 		if err != nil {
// 			http.Error(w, "Failed to save image", http.StatusInternalServerError)
// 			return
// 		}
// 		testimonial.Image = imagePath
// 	}

// 	utils.DB.Save(&testimonial)
// 	json.NewEncoder(w).Encode(testimonial)
// }

// func DeleteTestimonial(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id, _ := strconv.Atoi(params["id"])
// 	var testimonial models.Testimonial
// 	utils.DB.Delete(&testimonial, id)
// 	json.NewEncoder(w).Encode("The testimonial is deleted successfully!")
// }

func GetTestimonials(w http.ResponseWriter, r *http.Request) {
	var testimonials []models.Testimonial
	utils.DB.Find(&testimonials)
	json.NewEncoder(w).Encode(testimonials)
}

func GetTestimonial(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var testimonial models.Testimonial
	utils.DB.First(&testimonial, id)
	json.NewEncoder(w).Encode(testimonial)
}

func CreateTestimonial(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Invalid image file", http.StatusBadRequest)
		return
	}

	imagePath, err := utils.SaveFile(file, handler)
	if err != nil {
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
		return
	}

	testimonial := models.Testimonial{
		Title:       title,
		Description: description,
		Image:       imagePath,
	}

	utils.DB.Create(&testimonial)
	json.NewEncoder(w).Encode(testimonial)
}

func UpdateTestimonial(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var testimonial models.Testimonial
	utils.DB.First(&testimonial, id)

	title := r.FormValue("title")
	description := r.FormValue("description")
	testimonial.Title = title
	testimonial.Description = description

	file, handler, err := r.FormFile("image")
	if err == nil {
		imagePath, err := utils.SaveFile(file, handler)
		if err != nil {
			http.Error(w, "Failed to save image", http.StatusInternalServerError)
			return
		}
		testimonial.Image = imagePath
	}

	utils.DB.Save(&testimonial)
	json.NewEncoder(w).Encode(testimonial)
}

func DeleteTestimonial(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var testimonial models.Testimonial
	utils.DB.Delete(&testimonial, id)
	json.NewEncoder(w).Encode("The testimonial is deleted successfully!")
}
