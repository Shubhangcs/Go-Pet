package routes

import (
	"database/sql"

	"github.com/VeereshAkki/Pet_App_Backend/controllers"
	"github.com/VeereshAkki/Pet_App_Backend/repository"
	"github.com/gorilla/mux"
)


func UserRouter(router *mux.Router , db *sql.DB){
	repo := repository.NewUserRepository(db)
	cont := controllers.NewUserController(repo)

	router.HandleFunc("/register" , cont.RegisterUserController).Methods("POST")
	router.HandleFunc("/login" , cont.LoginUserController).Methods("POST")
	router.HandleFunc("/forgotpassword" , cont.ForgotPasswordController).Methods("POST")
}