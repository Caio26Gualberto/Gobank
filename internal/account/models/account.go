package models

type Account struct {
	ID        int64   `db:"id" json:"id"`
	Owner     string  `db:"owner" json:"owner" validate:"required,min=3"`
	Balance   float64 `db:"balance" json:"balance" validate:"gte=0"`
	Currency  string  `db:"currency" json:"currency" validate:"required,oneof=USD BRL EUR"`
	CreatedAt string  `db:"created_at" json:"created_at"`
}
