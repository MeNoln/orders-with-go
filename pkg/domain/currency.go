package domain

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
