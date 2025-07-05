package repository

import (
	"database/sql"

	"github.com/Caio26Gualberto/gobank/internal/account/models"
)

type SQLAccountRepository struct {
	DB *sql.DB
}

func NewSQLAccountRepository(db *sql.DB) *SQLAccountRepository {
	return &SQLAccountRepository{DB: db}
}

func (r *SQLAccountRepository) Create(account *models.Account) (int64, error) {
	query := `INSERT INTO accounts (owner, balance, currency, created_at)
	VALUES (@p1, @p2, @p3, GETDATE()); SELECT SCOPE_IDENTITY();`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(account.Owner, account.Balance, account.Currency).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SQLAccountRepository) GetById(id int64) (*models.Account, error) {
	query := "SELECT id, owner, balance FROM accounts WHERE id = @p1"
	row := r.DB.QueryRow(query, id)

	var account models.Account
	err := row.Scan(&account.ID, &account.Owner, &account.Balance)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *SQLAccountRepository) List() ([]*models.Account, error) {
	rows, err := r.DB.Query("SELECT id, owner, balance FROM accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*models.Account
	for rows.Next() {
		var account models.Account
		if err := rows.Scan(&account.ID, &account.Owner, &account.Balance); err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *SQLAccountRepository) Update(account *models.Account) error {
	query := "UPDATE accounts SET owner = @p1, balance = @p2 WHERE id = @p3"
	_, err := r.DB.Exec(query, account.Owner, account.Balance, account.ID)
	return err
}

func (r *SQLAccountRepository) Delete(id int64) error {
	query := "DELETE FROM accounts WHERE id = @p1"
	_, err := r.DB.Exec(query, id)
	return err
}
