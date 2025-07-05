package account

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Caio26Gualberto/gobank/internal/account/models"
	"github.com/Caio26Gualberto/gobank/internal/account/repository"
	"github.com/Caio26Gualberto/gobank/internal/middlewares"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	Repo repository.AccountRepository
}

func NewAccountHandler(repo repository.AccountRepository) *AccountHandler {
	return &AccountHandler{Repo: repo}
}

func (h *AccountHandler) CreateAccount() http.HandlerFunc {
	return middlewares.ValidateJSON(func(w http.ResponseWriter, r *http.Request, payload *models.Account) {
		var acc models.Account
		if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
			http.Error(w, "Invalid payload", http.StatusBadRequest)
			return
		}

		id, err := h.Repo.Create(&acc)
		if err != nil {
			http.Error(w, "Error creating account", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Account created successfully",
			"id":      id,
		})
	})
}

func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	account, err := h.Repo.GetById(id)
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.Repo.List()
	if err != nil {
		http.Error(w, "Accounts not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(accounts)
}

func (h *AccountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var acc *models.Account
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	acc.ID = id

	if err := h.Repo.Update(acc); err != nil {
		http.Error(w, "Error updating account", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Account updated successfully",
	})
}

func (h *AccountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, "Error deleting account", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Account deleted successfully",
	})
}
