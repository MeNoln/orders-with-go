package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MeNoln/orders-with-go/internal/config"
	"github.com/MeNoln/orders-with-go/internal/currency"
	"github.com/MeNoln/orders-with-go/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := config.Load(getConfigEnv())
	if err != nil {
		log.Fatalln("Failed to setup config")
		os.Exit(-1)
	}

	if err = database.ValidateConnectivity(); err != nil {
		os.Exit(-1)
	}

	router := gin.Default()
	currency.RegisterCurrencyRoutes(router)

	router.Run(fmt.Sprintf(":%s", config.Cfg.Port))
}

func getConfigEnv() string {
	const localCfg string = "local"
	if cfgEnv := os.Getenv("APP_ENV"); len(cfgEnv) != 0 {
		return cfgEnv
	}

	return localCfg
}
