package models

type Account struct {
	ID        int64   `db:"id"`
	Owner     string  `db:"owner"`
	Balance   float64 `db:"balance"`
	Currency  string  `db:"currency"`
	CreatedAt string  `db:"created_at"`
}
