package database

import (
	"github.com/andre250899/go-module-03-challenge/models"
	"github.com/google/uuid"
)

var DB *map[models.Id]models.User

func InitDB() {
	tempDB := map[models.Id]models.User{
		models.Id(uuid.NewString()): {FirstName: "John", LastName: "Doe", FullName: "John Doe", Biography: "A simple man"},
		models.Id(uuid.NewString()): {FirstName: "Jane", LastName: "Smith", FullName: "Jane Smith", Biography: "A complex woman"},
	}
	DB = &tempDB
}
