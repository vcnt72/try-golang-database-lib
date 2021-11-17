package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vcnt72/try-golang-database-lib/usecase/order_summary"
	"github.com/vcnt72/try-golang-database-lib/usecase/product"
	"github.com/vcnt72/try-golang-database-lib/usecase/user"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HandlerParams struct {
	fx.In
	OrderSummaryService order_summary.Usecase
	UserService         user.Usecase
	ProductService      product.Usecase
}

func NewGin(lc fx.Lifecycle, logger *zap.Logger, h HandlerParams) *gin.Engine {

	g := gin.Default()

	port := viper.GetString("server.port")

	srv := http.Server{
		Addr:    ":" + port,
		Handler: g,
	}

	baseGroup := g.Group("/api")
	{
		NewOrderSummaryHandler(baseGroup, h.OrderSummaryService)
		NewProductHandler(baseGroup, h.ProductService)
		NewUserHandler(baseGroup, h.UserService)
	}

	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {

			// Initializing the server in a goroutine so that
			// it won't block the graceful shutdown handling below
			go func() {

				logger.Sugar().Infof("Listen: %s", port)

				if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
					logger.Sugar().Infof("Listen: %s", err)
				}

			}()

			return nil
		},
		OnStop: func(c context.Context) error {
			logger.Info("Shutting down server...")

			return srv.Shutdown(c)
		},
	})

	return g
}
