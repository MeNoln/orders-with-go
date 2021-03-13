package database

import (
	log "github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/stdlib" //sql lib
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// ValidateConnectivity checks connection to database at start
func ValidateConnectivity() error {
	db, err := GetDbContext()
	defer db.Close()

	if err != nil {
		log.Fatal("Failed to connect to rdbms")
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database")
		return err
	}

	log.Info("Database connection valid")
	return nil
}

// GetDbContext gets database context
func GetDbContext() (*sqlx.DB, error) {
	return sqlx.Connect("pgx", viper.GetString("dbString"))
}
