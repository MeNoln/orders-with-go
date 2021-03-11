package currency

import (
	"net/http"

	"github.com/MeNoln/orders-with-go/internal/common"
	"github.com/gin-gonic/gin"
)

// RegisterCurrencyRoutes ...
func RegisterCurrencyRoutes(router *gin.Engine) {
	router.POST("v1/currency", createCurrency)
}

func createCurrency(c *gin.Context) {
	var body CreateCurrencyCommand

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.ErrorResponse{
			Message: "Failed to read body",
		})
	}

	currencyService := CreateService()
	err := currencyService.CreateCurrency(&body)
	if err != nil {
		c.JSON(400, common.ErrorResponse{
			Message: err.Error(),
		})
	}

	c.Status(200)
}
