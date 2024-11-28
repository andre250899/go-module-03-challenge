package api

import (
	"encoding/json"
	"net/http"

	"github.com/andre250899/go-module-03-challenge/models"
	"github.com/andre250899/go-module-03-challenge/repository"
	"github.com/go-chi/chi/v5"
)

type PostBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Biography string `json:"bio"`
}

func handleCreateUser(db *map[models.Id]models.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			sendJSON(
				w,
				Response{Error: "Please provide FirstName LastName and bio for the user"},
				http.StatusBadRequest,
			)
			return
		}

		user, err := repository.Insert(db, body.FirstName, body.LastName, body.Biography)
		if err != nil {
			sendJSON(
				w,
				Response{Error: "There was an error while saving the user to the database"},
				http.StatusInternalServerError,
			)
			return
		}

		sendJSON(
			w,
			Response{Data: user},
			http.StatusCreated,
		)
	}
}

func handleFetchUsers(db *map[models.Id]models.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := repository.FindAllUsers(db)
		sendJSON(
			w,
			Response{Data: users},
			http.StatusOK,
		)
	}
}

func handleGetUser(db *map[models.Id]models.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paramId := chi.URLParam(r, "id")
		user, err := repository.FindById(db, models.Id(paramId))
		if err != nil {
			sendJSON(
				w,
				Response{Error: "The user with the specified ID does not exist"},
				http.StatusNotFound,
			)
			return
		}

		if user.Data == nil {
			sendJSON(
				w,
				Response{Error: "The user information could not be retrieved"},
				http.StatusInternalServerError,
			)
			return
		}

		sendJSON(
			w,
			Response{Data: user},
			http.StatusOK,
		)
	}
}

func handleDeleteUser(db *map[models.Id]models.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paramId := chi.URLParam(r, "id")
		_, err := repository.FindById(db, models.Id(paramId))
		if err != nil {
			sendJSON(
				w,
				Response{Error: "The user with the specified ID does not exist"},
				http.StatusNotFound,
			)
			return
		}

		_, err = repository.Delete(db, models.Id(paramId))
		if err != nil {
			sendJSON(
				w,
				Response{Error: "The user could not be removed"},
				http.StatusInternalServerError,
			)
			return
		}

		sendJSON(
			w,
			Response{Data: "User deleted successfully"},
			http.StatusOK,
		)

	}
}

func handleUpdateUser(db *map[models.Id]models.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paramId := chi.URLParam(r, "id")
		var body PostBody
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			sendJSON(
				w,
				Response{Error: "Please provide FirstName LastName and bio for the user"},
				http.StatusBadRequest,
			)
			return
		}

		_, err = repository.FindById(db, models.Id(paramId))
		if err != nil {
			sendJSON(
				w,
				Response{Error: "The user with the specified ID does not exist"},
				http.StatusNotFound,
			)
			return
		}

		user, err := repository.Update(db, models.Id(paramId), body.FirstName, body.LastName, body.Biography)
		if err != nil {
			sendJSON(
				w,
				Response{Error: "The user information could not be modified"},
				http.StatusInternalServerError,
			)
		}

		sendJSON(
			w,
			Response{Data: user},
			http.StatusOK,
		)
	}
}
