package domain

import "time"

// Order domain model which mapped to db
type Order struct {
	ID         int
	CurrencyID int
	Rate       float32
	CreatedAt  time.Time
}

// CreateOrder creates new domain record
func CreateOrder(currencyID int, rate float32, createdAt time.Time) *Order {
	return &Order{
		CurrencyID: currencyID,
		Rate:       rate,
		CreatedAt:  createdAt,
	}
}
