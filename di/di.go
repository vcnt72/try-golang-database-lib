package di

import (
	"github.com/vcnt72/try-golang-database-lib/api/handler"
	"github.com/vcnt72/try-golang-database-lib/config"
	"github.com/vcnt72/try-golang-database-lib/infrastructure/repository"
	"github.com/vcnt72/try-golang-database-lib/usecase/order_summary"
	"github.com/vcnt72/try-golang-database-lib/usecase/payment_method"
	"github.com/vcnt72/try-golang-database-lib/usecase/product"
	"github.com/vcnt72/try-golang-database-lib/usecase/user"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var diRepository = fx.Provide(
	repository.NewOrderSummaryInMemRepository,
	repository.NewPaymentMethodInMemRepository,
	repository.NewProductInMemRepository,
	repository.NewUserInMemRepository,
)

var diService = fx.Provide(user.NewService, product.NewService, payment_method.NewService, order_summary.NewService)

func NewApp() *fx.App {
	app := fx.New(fx.Invoke(func() {
		config.Init()
	}),
		// Logger
		fx.Provide(zap.NewProduction),
		diRepository,
		diService,
		fx.Invoke(handler.NewGin),
	)

	return app
}
