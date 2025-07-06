package transaction

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Caio26Gualberto/gobank/internal/middlewares"
	"github.com/Caio26Gualberto/gobank/internal/transaction/models"
	"github.com/Caio26Gualberto/gobank/internal/transaction/repository"
	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	Repo repository.TransactionRepository
}

func NewTransactionHandler(repo repository.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{Repo: repo}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var t *models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		middlewares.WriteError(w, http.StatusBadRequest, "INVALID_PAYLOAD", "Some error with de informations for create transaction")
		return
	}

	id, err := h.Repo.Create(t)
	if err != nil {
		middlewares.WriteError(w, http.StatusInternalServerError, "ERROR_CREATING_TRANSACTION", "Some error has been found for save the entity in database")
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Transaction created successfully",
		"id":      id,
	})
}

func (h *TransactionHandler) GetTransactionsByAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountIDStr := vars["id"]

	accountID, err := strconv.ParseInt(accountIDStr, 10, 64)
	if err != nil {
		middlewares.WriteError(w, http.StatusBadRequest, "INVALID_PARAMETER_ID", "Parameter Id is invalid, check with support")
		return
	}

	transactions, err := h.Repo.ListByAccountId(accountID)
	if err != nil {
		middlewares.WriteError(w, http.StatusNotFound, "TRANSACTION_NOT_FOUND", "Some error has been found for search the entities in database")
		return
	}

	json.NewEncoder(w).Encode(transactions)
}

func (h *TransactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountIdStr := vars["id"]

	accountId, err := strconv.ParseInt(accountIdStr, 10, 64)
	if err != nil {
		middlewares.WriteError(w, http.StatusBadRequest, "INVALID_PARAMETER_ID", "Parameter Id is invalid, check with support")
		return
	}

	if err := h.Repo.DeleteById(accountId); err != nil {
		middlewares.WriteError(w, http.StatusInternalServerError, "DELETE_ENTITY_ERROR", "Some error in the process DELETE in the database")
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Account deleted successfully",
	})
}
