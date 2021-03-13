package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/MeNoln/orders-with-go/pkg/currency"
	"github.com/MeNoln/orders-with-go/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	err := loadConfig()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(-1)
	}

	if err = database.ValidateConnectivity(); err != nil {
		log.Fatal(err.Error())
		os.Exit(-1)
	}

	router := gin.Default()

	currency.RegisterCurrencyRoutes(router)

	router.Run(fmt.Sprintf(":5000"))
}

func loadConfig() error {
	viper.SetConfigName(getRunningEnv())
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config/")

	return viper.ReadInConfig()
}

func getRunningEnv() string {
	const localCfg string = "local"
	if cfgEnv := os.Getenv("APP_ENV"); len(cfgEnv) != 0 {
		return cfgEnv
	}

	return localCfg
}
