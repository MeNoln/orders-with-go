package currency

import (
	"net/http"
	"strconv"

	"github.com/MeNoln/orders-with-go/internal/common"
	"github.com/gin-gonic/gin"
)

// RegisterCurrencyRoutes ...
func RegisterCurrencyRoutes(router *gin.Engine) {
	router.POST("v1/currency", createCurrency)
	router.GET("v1/currency", getAll)
	router.GET("v1/currency/:id", getByID)
}

func createCurrency(c *gin.Context) {
	var body CreateCurrencyCommand

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.ErrorResponse{
			Message: "Failed to read body",
		})
		return
	}

	currencyService := CreateService()
	err := currencyService.CreateCurrency(&body)
	if err != nil {
		c.JSON(400, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.Status(200)
}

func getAll(c *gin.Context) {
	currencyService := CreateService()

	currencies, err := currencyService.GetAll()
	if err != nil {
		c.JSON(400, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, &currencies)
}

func getByID(c *gin.Context) {
	currencyService := CreateService()

	queryID, _ := strconv.Atoi(c.Param("id"))

	dto, err := currencyService.GetByID(queryID)
	if err != nil {
		c.JSON(400, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, &dto)
}
