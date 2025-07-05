package models

type Transaction struct {
	ID        int64   `json:"id"`
	AccountID int64   `json:"account_id"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"` // "credit" or "debit"
}
