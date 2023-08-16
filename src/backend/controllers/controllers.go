package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VictorMilhomem/GoRestApi/src/backend/database"
	"github.com/VictorMilhomem/GoRestApi/src/backend/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalities

	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	var p models.Personalities

	database.DB.First(&p, id)

	json.NewEncoder(w).Encode(p)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personalities

	json.NewDecoder(r.Body).Decode(&personality)

	database.DB.Create(&personality)

	json.NewEncoder(w).Encode(personality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	var p models.Personalities

	database.DB.Delete(&p, id)

	json.NewEncoder(w).Encode(p)

}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	var p models.Personalities

	database.DB.First(&p, id)

	json.NewDecoder(r.Body).Decode(&p)

	database.DB.Save(&p)

	json.NewEncoder(w).Encode(p)

}
