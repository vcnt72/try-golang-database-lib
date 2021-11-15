package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vcnt72/try-golang-database-lib/usecase/order_summary"
	"github.com/vcnt72/try-golang-database-lib/usecase/product"
	"github.com/vcnt72/try-golang-database-lib/usecase/user"
	"go.uber.org/fx"
)

type HandlerParams struct {
	fx.In
	orderSummaryService order_summary.Usecase
	userService         user.Usecase
	productService      product.Usecase
}

func NewHandler(g *gin.Engine, h HandlerParams) {

	baseGroup := g.Group("/api")
	{
		NewOrderSummaryHandler(baseGroup, h.orderSummaryService)
		NewProductHandler(baseGroup, h.productService)
		NewUserHandler(baseGroup, h.userService)
	}
}
