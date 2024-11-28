package repository

import (
	"errors"
	"fmt"

	"github.com/andre250899/go-module-03-challenge/models"
	"github.com/google/uuid"
)

// create a function that findAll values of my db hash map
func FindAllUsers(db *map[models.Id]models.User) []models.Application {
	var users []models.Application
	for id, user := range *db {
		userMap := map[models.Id]models.User{
			models.Id(id): {
				FirstName: user.FirstName,
				LastName:  user.LastName,
				FullName:  user.FullName,
				Biography: user.Biography,
			},
		}
		users = append(users, models.Application{Data: userMap})
	}
	return users
}

func FindById(db *map[models.Id]models.User, id models.Id) (models.Application, error) {
	user, ok := (*db)[models.Id(id)]
	if !ok {
		return models.Application{}, fmt.Errorf("user not found: %v", id)
	}

	userMap := map[models.Id]models.User{
		models.Id(id): {
			FirstName: user.FirstName,
			LastName:  user.LastName,
			FullName:  user.FullName,
			Biography: user.Biography,
		},
	}

	return models.Application{Data: userMap}, nil
}

func Insert(db *map[models.Id]models.User, firstName, lastName, bio string) (models.Application, error) {
	// add a new user to the db
	newUser := models.User{
		FirstName: firstName,
		LastName:  lastName,
		FullName:  firstName + " " + lastName,
		Biography: bio,
	}

	id := uuid.New().String()
	(*db)[models.Id(id)] = newUser

	user, err := FindById(db, models.Id(id))
	if err != nil {
		return models.Application{}, fmt.Errorf("Insert failed: %v", err)
	}
	return user, nil
}

func Update(db *map[models.Id]models.User, id models.Id, firstName, lastName, bio string) (models.Application, error) {
	if _, ok := (*db)[id]; !ok {
		return models.Application{}, errors.New("user not found")
	}

	// Create a new user instance with the updated fields
	updatedUser := models.User{
		FirstName: firstName,
		LastName:  lastName,
		FullName:  firstName + " " + lastName,
		Biography: bio,
	}

	// Update the user in the map
	(*db)[id] = updatedUser

	// return the updated user by ID
	user, err := FindById(db, id)
	if err != nil {
		return models.Application{}, fmt.Errorf("Update failed: %v", err)
	}
	return user, nil
}

func Delete(db *map[models.Id]models.User, id models.Id) (models.Id, error) {
	if _, err := FindById(db, id); err != nil {
		return "", fmt.Errorf("Delete user failed: %v", err)
	}

	delete(*db, id)
	return id, nil
}
