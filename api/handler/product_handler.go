package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vcnt72/try-golang-database-lib/api/presenter"
	"github.com/vcnt72/try-golang-database-lib/api/request"
	"github.com/vcnt72/try-golang-database-lib/entity"
	"github.com/vcnt72/try-golang-database-lib/usecase/product"
)

func createProduct(productService product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()
		var createProductRequest request.CreateProductRequest
		if err := c.ShouldBindJSON(&createProductRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "bad request",
			})
			return
		}

		product, err := productService.Store(ctx, product.CreateProductDTO{
			Name:     createProductRequest.Name,
			Price:    createProductRequest.Price,
			Quantity: createProductRequest.Quantity,
			UserID:   createProductRequest.UserID,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success",
			"data": gin.H{
				"product": product,
			},
		})
	}
}

func updateProduct(productService product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

		defer cancel()

		var updateProductDTO request.UpdateProductRequest

		if err := c.ShouldBindJSON(&updateProductDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Bad Request",
			})
			return
		}

		productID := c.Param("productID")
		productObj, err := productService.GetByID(ctx, productID)

		if errors.Is(err, entity.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "Product not found",
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Unknown error",
			})
			return
		}

		productObj, err = productService.Update(ctx, product.UpdateProductDTO{
			ID:       productID,
			Name:     updateProductDTO.Name,
			Price:    updateProductDTO.Price,
			Quantity: updateProductDTO.Quantity,
		})

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "Unknown error",
			})
			return
		}

		productResp := presenter.ProductResponse{
			ID:       productObj.ID,
			Name:     productObj.Name,
			Price:    productObj.Price,
			Quantity: productObj.Quantity,
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Success",
			"data": gin.H{
				"product": productResp,
			},
		})
	}
}

func getProductByID(productService product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

		defer cancel()
		productID := c.Param("productID")

		productObj, err := productService.GetByID(ctx, productID)

		if errors.Is(err, entity.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "Product not found",
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Unknown error",
			})
			return
		}

		productResp := presenter.ProductResponse{
			ID:       productObj.ID,
			Name:     productObj.Name,
			Price:    productObj.Price,
			Quantity: productObj.Quantity,
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Success",
			"data": gin.H{
				"product": productResp,
			},
		})
	}
}

func paginateProduct(productService product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

		defer cancel()

		var paginateDTO request.PaginateProductRequest

		if err := c.ShouldBindJSON(&paginateDTO); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Bad Request",
			})

			return
		}

		paginate := &entity.Paginate{
			Page:    paginateDTO.Page,
			PerPage: paginateDTO.PerPage,
		}

		products, err := productService.Paginate(ctx, product.FilterDTO{
			ProductName: paginateDTO.ProductName,
			UserID:      paginateDTO.UserID,
			ProductIDs:  paginateDTO.ProductIDs,
		}, paginate)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Unknown error",
			})
			return
		}

		var productResponses []presenter.ProductResponse

		for _, v := range products {
			productResponse := presenter.ProductResponse{
				ID:       v.ID,
				Name:     v.Name,
				Price:    v.Price,
				Quantity: v.Quantity,
			}

			productResponses = append(productResponses, productResponse)
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Success",
			"date": gin.H{
				"products": productResponses,
				"paginate": presenter.PaginatePresenter{
					Page:    paginate.Page,
					PerPage: paginate.PerPage,
				},
			},
		})
	}
}

func NewProductHandler(base *gin.RouterGroup, productService product.Usecase) {
	productGroup := base.Group("products")
	{
		productGroup.POST("", createProduct(productService))
		productGroup.GET("", paginateProduct(productService))
		productGroup.GET(":productID", getProductByID(productService))
		productGroup.PUT(":productID", updateProduct(productService))
	}
}
