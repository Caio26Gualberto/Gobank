package transaction

import (
	"encoding/json"
	"net/http"
	"strconv"

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
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	id, err := h.Repo.Create(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	transactions, err := h.Repo.ListByAccountId(accountID)
	if err != nil {
		http.Error(w, "Failed to fetch transactions", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transactions)
}
