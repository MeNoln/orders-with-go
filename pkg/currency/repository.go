package currency

import (
	"github.com/MeNoln/orders-with-go/pkg/database"
	"github.com/MeNoln/orders-with-go/pkg/domain"
	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	Create(currency *domain.Currency) error
	GetAll() ([]domain.Currency, error)
	Get(id int) (domain.Currency, error)
	Dispose()
}

type repository struct {
	db *sqlx.DB
}

// CreateRepository ...
func CreateRepository() Repository {
	db, _ := database.GetDbContext()
	return repository{db}
}

func (r repository) Create(currency *domain.Currency) error {
	query := `insert into currency(name, title) values($1, $2)`
	queryArgs := []interface{}{
		currency.Name,
		currency.Title,
	}

	_, err := r.db.Exec(query, queryArgs...)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) GetAll() ([]domain.Currency, error) {
	var currencies []domain.Currency

	query := `select id, name, title from currency`

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var currency domain.Currency

		err = rows.StructScan(&currency)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (r repository) Get(id int) (domain.Currency, error) {
	query := `select id, name, title from currency where id = $1`

	var currency domain.Currency
	err := r.db.QueryRowx(query, id).StructScan(&currency)
	if err != nil {
		return domain.Currency{}, err
	}

	return currency, nil
}

func (r repository) Dispose() {
	r.db.Close()
}
