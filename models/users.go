package models

import (
	"io"
)

type UserModel struct{
	UserName string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	PetsRes int `json:"res"`
	PetsAdop int `json:"adopt"`
}

type UserInterface interface{
	RegisterUser(*io.ReadCloser) error
	LoginUser(*io.ReadCloser) (string , error)
	ForgotPassword(*io.ReadCloser) error
}