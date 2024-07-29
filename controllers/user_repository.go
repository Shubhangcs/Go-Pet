package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/VeereshAkki/Pet_App_Backend/models"
)

type UserController struct{
	userRepositoryIns  models.UserInterface
}

func NewUserController(uri models.UserInterface) *UserController {
	return &UserController{
		userRepositoryIns: uri,
	}
}

func(uc *UserController) RegisterUserController(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	err := uc.userRepositoryIns.RegisterUser(&r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorPayload{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SuccessPayload{Message: "Success"})
}

func(uc *UserController) LoginUserController(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	res , err := uc.userRepositoryIns.LoginUser(&r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorPayload{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SuccessPayload{Message: "Success" , Data: res})
}

func(uc *UserController) ForgotPasswordController(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	 err := uc.userRepositoryIns.ForgotPassword(&r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorPayload{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SuccessPayload{Message: "Success" , Data: nil})
}