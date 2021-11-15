package di

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewGin(lc fx.Lifecycle, logger *zap.Logger) *gin.Engine {

	g := gin.Default()

	port := viper.GetString("server.port")

	srv := http.Server{
		Addr:    ":" + port,
		Handler: g,
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
