package rest

import (
	"context"
	"github.com/NuttayotSukkum/purchase/internal/handlers"
	"github.com/NuttayotSukkum/purchase/internal/repositories/db"
	"github.com/NuttayotSukkum/purchase/internal/services"
	"github.com/labstack/echo/v4"
	logger "github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func ProductRouter(ctx context.Context, e *echo.Echo, database *gorm.DB) {
	productRepo := db.NewProductRepositoryImpl(database)
	paymentRepo := db.NewPaymentRepositoryImpl(database)
	productSvc := services.NewProductServiceImpl(productRepo)
	paymentSvc := services.NewPaymentServiceImpl(paymentRepo)
	productHandler := handlers.NewProductHandler(productSvc, paymentSvc)
	logger.Infof("%v Check data", ctx)
	productMn := e.Group("/product-manage")
	api := productMn.Group("/api")
	api.POST("/create-product", productHandler.CreateProductHandler)
	api.POST("/edit-product", productHandler.EditProductHandler)
	api.POST("/get-product", productHandler.GetProductHandler)
	api.POST("/search-product", productHandler.PartialSearchProduct)
	api.POST("/submit-purchase", productHandler.CreatePayment)
	api.POST("/confirm-purchase", productHandler.ConfirmPurchaseHandler)

}
