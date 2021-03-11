package currency

import (
	"log"

	"github.com/MeNoln/orders-with-go/internal/domain"
)

// Service ...
type Service interface {
	CreateCurrency(command *CreateCurrencyCommand) error
}

type service struct {
	repo Repository
}

// CreateService ...
func CreateService() Service {
	return service{repo: CreateRepository()}
}

// CreateCurrencyCommand ...
type CreateCurrencyCommand struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

func (s service) CreateCurrency(command *CreateCurrencyCommand) error {
	defer s.repo.Dispose()
	currency := domain.CreateCurrency(command.Name, command.Title)

	err := s.repo.Create(currency)
	if err != nil {
		log.Fatalln("Failed to create new currency: %s, %s", command.Name, command.Title)
		return err
	}

	return nil
}
