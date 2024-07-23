package repository

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"

	"github.com/VeereshAkki/Pet_App_Backend/models"
)

type PetRepository struct {
	db *sql.DB
}

func NewPetRepository(db *sql.DB) *PetRepository {
	return &PetRepository{
		db,
	}
}

func (pr *PetRepository) GetPetDetails() ([]models.PetCard, error) {
	res, err := pr.db.Query("SELECT * FROM adoption")

	if err != nil {
		return nil, err
	}
	defer res.Close()

	var pets []models.PetCard
	var pet models.PetCard

	for res.Next() {
		err := res.Scan(&pet.Name, &pet.Type, &pet.Age, &pet.Phone, &pet.Image, &pet.Details)

		if err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}

	if res.Err() != nil {
		return nil, res.Err()
	}

	return pets, nil
}

func (pr *PetRepository) AddPets(r *io.ReadCloser) error {
	var newPet models.PetCard
	err := json.NewDecoder(*r).Decode(&newPet)
	if err != nil {
		return err
	}
	res, err := pr.db.Exec("INSERT INTO adoption VALUES($1 , $2 , $3 , $4 , $5 , $6)", newPet.Name, newPet.Type, newPet.Age, newPet.Phone, newPet.Image, newPet.Details)
	if err != nil {
		return err
	}
	log.Println(res)
	return nil
}

func (pr *PetRepository) GetPerticularDetail(r *io.ReadCloser) (*models.PetCard, error) {
	var detPet models.PetCard
	err := json.NewDecoder(*r).Decode(&detPet)
	if err != nil {
		return nil , err
	}
	return &detPet , nil
}
