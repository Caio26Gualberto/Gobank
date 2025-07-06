package repository

import (
	"database/sql"

	"github.com/Caio26Gualberto/gobank/internal/transaction/models"
)

type SQLTransactionRepository struct {
	DB *sql.DB
}

func NewSQLTransactionRepository(db *sql.DB) *SQLTransactionRepository {
	return &SQLTransactionRepository{DB: db}
}

func (r *SQLTransactionRepository) Create(t *models.Transaction) (int64, error) {
	var insertedID int64
	query := `
        INSERT INTO Transactions (AccountID, Amount, Type)
        OUTPUT INSERTED.ID
        VALUES (@accountID, @amount, @type)
    `
	err := r.DB.QueryRow(query,
		sql.Named("accountID", t.AccountID),
		sql.Named("amount", t.Amount),
		sql.Named("type", t.Type),
	).Scan(&insertedID)

	if err != nil {
		return 0, err
	}
	return insertedID, nil
}

func (r *SQLTransactionRepository) ListByAccountId(accountID int64) ([]*models.Transaction, error) {
	query := "SELECT ID, AccountID, Amount, Type FROM Transactions WHERE AccountID = @accountID"
	rows, err := r.DB.Query(query, sql.Named("accountID", accountID))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var txs []*models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.AccountID, &t.Amount, &t.Type); err != nil {
			return nil, err
		}
		tCopy := t
		txs = append(txs, &tCopy)
	}
	return txs, nil
}

func (r *SQLTransactionRepository) DeleteById(accountId int64) error {
	query := "DELETE Transactions WHERE Id = @id"
	rows, err := r.DB.Query(query, sql.Named("id", accountId))

	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
