package routes

import (
	"database/sql"

	"github.com/VeereshAkki/Pet_App_Backend/controllers"
	"github.com/VeereshAkki/Pet_App_Backend/repository"
	"github.com/gorilla/mux"
)


func PetCardRoutes(router *mux.Router , db *sql.DB){
	repo := repository.NewPetRepository(db)
	cont := controllers.NewPetController(repo)

	router.HandleFunc("/getpets" , cont.PetDetailsGetController).Methods("GET")
	router.HandleFunc("/addpets" , cont.PetDetailsAddController).Methods("POST")
	router.HandleFunc("/getdet" , cont.PetDetailController).Methods("POST")
}