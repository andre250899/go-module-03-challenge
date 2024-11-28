package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/andre250899/go-module-03-challenge/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Error string `json:"message,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal json data", "error", err)
		sendJSON(
			w,
			Response{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}
}

func NewHandler(db map[models.Id]models.User) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/users", handleCreateUser(&db))
	r.Get("/api/users", handleFetchUsers(&db))
	r.Get(("/api/users/{id}"), handleGetUser(&db))
	r.Delete(("/api/users/{id}"), handleDeleteUser(&db))
	r.Put(("/api/users/{id}"), handleUpdateUser(&db))

	return r
}
