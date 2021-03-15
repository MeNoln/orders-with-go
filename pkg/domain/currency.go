package domain

import (
	"encoding/json"
)

// Currency domain model which mapped to db
type Currency struct {
	ID    int
	Name  string
	Title string
}

// CreateCurrency creates new domain record
func CreateCurrency(name string, title string) *Currency {
	return &Currency{
		Name:  name,
		Title: title,
	}
}

type CurrencyCreatedEvent struct {
	MsgId string
	Title string
}

func (c CurrencyCreatedEvent) ToMessage() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return ""
	}

	return string(bytes)
}
