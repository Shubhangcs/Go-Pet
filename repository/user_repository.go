package repository

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/VeereshAkki/Pet_App_Backend/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func(ur *UserRepository) RegisterUser(r *io.ReadCloser) error {
	var newUser models.UserModel
	err := json.NewDecoder(*r).Decode(&newUser)
	if err != nil {
		return err
	}
	pass , err := bcrypt.GenerateFromPassword([]byte(newUser.Password) , 2)
	if err != nil {
		return err
	}
	res , err :=  ur.db.Exec("INSERT INTO users VALUES($1 , $2 , $3 , $4 , $5)" , newUser.UserName , newUser.Email , string(pass) , newUser.PetsRes , newUser.PetsAdop)
	if err != nil {
		return err
	}
	log.Println(res)
	return nil
}

func(ur *UserRepository) LoginUser(r *io.ReadCloser) (string , error) {
	var usr models.UserModel
	var tst models.UserModel
	var secretKey = []byte("secret-key")
	err := json.NewDecoder(*r).Decode(&usr)
	if err !=  nil {
		return "" , err
	}
	 err = ur.db.QueryRow("SELECT email , password FROM users WHERE email = $1" , usr.Email).Scan(&tst.Email , &tst.Password)
	if err != nil {
		return "" , err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(tst.Password) , []byte(usr.Password)); err != nil {
		return "" , err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , jwt.MapClaims{
		"email": tst.Email,
		"password": tst.Password,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString , err := token.SignedString(secretKey)
	if err != nil {
		return "" , err
	}
	return tokenString , nil
}

func(ur *UserRepository) ForgotPassword(r *io.ReadCloser) error {
	var usr models.UserModel
	err := json.NewDecoder(*r).Decode(&usr)
	if err != nil {
		return err
	}
	pass , err := bcrypt.GenerateFromPassword([]byte(usr.Password) , 2)
	if err != nil {
		return err
	}
	res , err := ur.db.Exec("UPDATE users SET password=$1 WHERE email=$2" , pass , usr.Email);
	if err != nil {
		return err
	}
	log.Println(res)
	return nil
}