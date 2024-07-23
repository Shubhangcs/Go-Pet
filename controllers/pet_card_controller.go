package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/VeereshAkki/Pet_App_Backend/models"
)

type PetController struct{
	petRepository models.PetInterface
}

func NewPetController(petRepository models.PetInterface) *PetController {
	return &PetController{
		petRepository,
	}
}

func(pc *PetController) PetDetailsGetController(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	res , err := pc.petRepository.GetPetDetails()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorPayload{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(models.SuccessPayload{Message: "Success" , Data: res})
}

func(pc *PetController) PetDetailsAddController(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	err := pc.petRepository.AddPets(&r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorPayload{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(models.SuccessPayload{Message: "Success" , Data: nil})
}

func(pc *PetController) PetDetailController(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	res , err := pc.petRepository.GetPerticularDetail(&r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorPayload{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(models.SuccessPayload{Message: "Success" , Data: res})
}

