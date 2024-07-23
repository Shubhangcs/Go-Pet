package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/VeereshAkki/Pet_App_Backend/routes"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)


func main(){
	// db , err := sql.Open("postgres" , "host=localhost port=5432 user=root dbname=pet sslmode=disable password=password")
	db , err := sql.Open("postgres" , "postgresql://root:xuc87BFMfwzq2xqFJHXf3I1AfNbwwTOz@dpg-cqg05256l47c73bn7gig-a.oregon-postgres.render.com/pet_69y5")
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
	var router = mux.NewRouter()

	log.Println("Server Is Running At Port 8000")

	routes.PetCardRoutes(router , db)
	routes.UserRouter(router , db)

	log.Fatal(http.ListenAndServe(":8000" , router))
}