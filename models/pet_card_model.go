package models

import "io"

type PetCard struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Age     string `json:"age"`
	Phone   string `json:"phone"`
	Image   string `json:"image"`
	Details string `json:"description"`
	Adopted bool `json:"adopted"`
}

type PetInterface interface {
	GetPetDetails() ([]PetCard, error)
	AddPets(*io.ReadCloser) error
	GetPerticularDetail(*io.ReadCloser) (*PetCard, error)
}
