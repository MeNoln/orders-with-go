package currency

import (
	"log"

	"github.com/MeNoln/orders-with-go/internal/domain"
)

// Service ...
type Service interface {
	CreateCurrency(command *CreateCurrencyCommand) error
	GetAll() (*GetAllCurrenciesResponse, error)
	GetByID(id int) (*CurrencyDto, error)
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
		log.Printf("Failed to create new currency: %s, %s\n", command.Name, command.Title)
		log.Println(err.Error())
		return err
	}

	log.Printf("Currency %s added", command.Name)
	return nil
}

// GetAllCurrenciesResponse ...
type GetAllCurrenciesResponse struct {
	Currencies []CurrencyDto `json:"currencies"`
}

//CurrencyDto ...
type CurrencyDto struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

func (s service) GetAll() (*GetAllCurrenciesResponse, error) {
	defer s.repo.Dispose()

	var currencyDtos []CurrencyDto
	currencies, err := s.repo.GetAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for _, currency := range currencies {
		currencyDtos = append(currencyDtos, CurrencyDto{
			ID:    currency.ID,
			Name:  currency.Name,
			Title: currency.Title,
		})
	}

	return &GetAllCurrenciesResponse{Currencies: currencyDtos}, nil
}

func (s service) GetByID(id int) (*CurrencyDto, error) {
	defer s.repo.Dispose()

	currency, err := s.repo.Get(id)
	if err != nil {
		log.Println(err.Error())
		return &CurrencyDto{}, err
	}

	dto := &CurrencyDto{
		ID:    currency.ID,
		Name:  currency.Name,
		Title: currency.Title,
	}

	return dto, nil
}
