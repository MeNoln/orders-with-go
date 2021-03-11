package database

import (
	"log"

	"github.com/MeNoln/orders-with-go/internal/config"
	_ "github.com/jackc/pgx/stdlib" //sql lib
	"github.com/jmoiron/sqlx"
)

// ValidateConnectivity checks connection to database at start
func ValidateConnectivity() error {
	db, err := GetDbContext()
	defer db.Close()

	if err != nil {
		log.Fatalln("Failed to connect to rdbms")
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Failed to ping database")
		return err
	}

	log.Println("Database connection valid")
	return nil
}

// GetDbContext gets database context
func GetDbContext() (*sqlx.DB, error) {
	return sqlx.Connect("pgx", config.Cfg.DB)
}
