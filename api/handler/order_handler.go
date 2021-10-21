package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vcnt72/try-golang-database-lib/usecase/order_summary"
)

func createOrder(orderService order_summary.Usecase) gin.HandlerFunc {

	return func(c *gin.Context) {

	}
}
