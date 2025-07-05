package repository

import "github.com/Caio26Gualberto/gobank/internal/transaction/models"

type TransactionRepository interface {
	Create(t *models.Transaction) (int64, error)
	ListByAccountId(accountId int64) ([]*models.Transaction, error)
}
