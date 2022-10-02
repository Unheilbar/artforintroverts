package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/unheilbar/artforintrovert_entry_task/internal/entities"
)

type service interface {
	GetAll() ([]entities.User, error)
	Delete(ID string) error
	Update(ID string, user entities.User) error
}

type handlers struct {
	service service
}

func NewHandlers(service service) *handlers {
	return &handlers{service: service}
}
func (h *handlers) GetRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/user", h.getAllUsers)
	router.Put("/user/{user_id}", h.updateUser)
	router.Delete("/user/{user_id}", h.deleteUser)

	return router
}

func (h *handlers) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll()

	if err != nil {
		http.Error(w, fmt.Sprintf("error handling request, %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h *handlers) updateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")

	var user entities.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil || userID == "" {
		http.Error(w, fmt.Sprintf("error parsing request body, %v", err), http.StatusBadRequest)
		return
	}

	err = h.service.Update(userID, user)
	if err != nil {
		http.Error(w, fmt.Sprintf("error handling request, %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *handlers) deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")

	err := h.service.Delete(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("error handling request, %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
