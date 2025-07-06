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
			middlewares.WriteError(w, http.StatusBadRequest, "INVALID_PAYLOAD", "Some error with de informations for create acount")
			return
		}

		id, err := h.Repo.Create(&acc)
		if err != nil {
			middlewares.WriteError(w, http.StatusInternalServerError, "ERROR_CREATING_ACCOUNT", "Some error has been found for save the entity in database")
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
		middlewares.WriteError(w, http.StatusBadRequest, "INVALID_PARAMETER_ID", "Parameter Id is invalid, check with support")
		return
	}

	account, err := h.Repo.GetById(id)
	if err != nil {
		middlewares.WriteError(w, http.StatusNotFound, "ACCOUNT_NOT_FOUND", "Some error has been found for search the entity in database")
		return
	}

	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.Repo.List()
	if err != nil {
		middlewares.WriteError(w, http.StatusNotFound, "ACCOUNTS_NOT_FOUND", "Some error has been found for search the entities in database")
		return
	}

	json.NewEncoder(w).Encode(accounts)
}

func (h *AccountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		middlewares.WriteError(w, http.StatusBadRequest, "INVALID_PARAMETER_ID", "Parameter Id is invalid, check with support")
		return
	}

	var acc *models.Account
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		middlewares.WriteError(w, http.StatusBadRequest, "INVALID_JSON", "The struct or some value in JSON is not valid")
		return
	}

	acc.ID = id

	if err := h.Repo.Update(acc); err != nil {
		middlewares.WriteError(w, http.StatusInternalServerError, "UPDATING_ENTITY_ERROR", "Some error in the process UPDATE in the database")
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
		middlewares.WriteError(w, http.StatusBadRequest, "INVALID_PARAMETER_ID", "Parameter Id is invalid, check with support")
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		middlewares.WriteError(w, http.StatusInternalServerError, "DELETE_ENTITY_ERROR", "Some error in the process DELETE in the database")
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Account deleted successfully",
	})
}
