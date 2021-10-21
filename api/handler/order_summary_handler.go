package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vcnt72/try-golang-database-lib/api/presenter"
	"github.com/vcnt72/try-golang-database-lib/api/request"
	"github.com/vcnt72/try-golang-database-lib/usecase/order_summary"
)

func searchOrder(orderService order_summary.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

		defer cancel()

		var filterDTO request.OrderSummaryFilterDTO

		if err := c.ShouldBindJSON(&filterDTO); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Bad Request",
			})
			return
		}

		orderSummary, err := orderService.Search(ctx, order_summary.FilterDTO{
			StartDate: filterDTO.StartDate,
			EndDate:   filterDTO.EndDate,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Unknown error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Success",
			"data": gin.H{
				"order_summary": orderSummary,
			},
		})
	}
}

func createOrder(orderService order_summary.Usecase) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

		defer cancel()

		var createOrderRequest request.CreateOrderSummaryDTO

		if err := c.ShouldBindJSON(&createOrderRequest); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "bad request",
			})

			return
		}

		var orderItemDTOs []order_summary.CreateOrderItemDTO

		for _, v := range createOrderRequest.OrderItems {
			orderItemDTO := order_summary.CreateOrderItemDTO{
				ProductID: v.ProductID,
				Quantity:  v.Quantity,
			}

			orderItemDTOs = append(orderItemDTOs, orderItemDTO)
		}

		createOrderSummaryDTO := order_summary.CreateOrderSummaryDTO{
			PaymentMethodID: createOrderRequest.PaymentMethodID,
			UserID:          createOrderRequest.UserID,
			OrderItems:      orderItemDTOs,
		}

		orderSummary, err := orderService.Store(ctx, createOrderSummaryDTO)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Unknown error",
			})
			return
		}

		var orderItemResponses []presenter.OrderItemResponse

		for _, v := range orderSummary.OrderItems {
			orderItemResponse :=
				presenter.OrderItemResponse{
					ID: v.ID,
					Product: presenter.ProductResponse{
						ID:       v.Product.ID,
						Name:     v.Product.Name,
						Price:    v.Product.Price,
						Quantity: v.Product.Quantity,
					},
					Quantity: v.Quantity,
					Total:    float64(v.Quantity) * v.ProductPrice,
				}
			orderItemResponses = append(orderItemResponses, orderItemResponse)
		}

		orderSummaryResp := presenter.OrderSummaryResponse{
			PaymentMethod: presenter.PaymentMethodResponse{
				ID:   orderSummary.PaymentMethod.ID,
				Name: orderSummary.PaymentMethod.Name,
			},
			User: presenter.UserResponse{
				ID:    orderSummary.User.ID,
				Name:  orderSummary.User.Name,
				Email: orderSummary.User.Email,
			},
			Total: orderSummary.Total,
			ID:    orderSummary.ID,

			OrderItems: orderItemResponses,
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Success",
			"data": gin.H{
				"order_summary": orderSummaryResp,
			},
		})
	}
}

func NewOrderSummaryHandler(base *gin.RouterGroup, orderSummaryService order_summary.Usecase) {
	orderSummaryGroup := base.Group("order-summaries")
	{
		orderSummaryGroup.GET("", searchOrder(orderSummaryService))
		orderSummaryGroup.POST("", createOrder(orderSummaryService))
	}
}
