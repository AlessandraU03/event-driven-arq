package routes

import (
	"eventdriven/src/internal/application/useCases"
	"eventdriven/src/internal/application/useCases/products"
	"eventdriven/src/internal/application/services"
	"eventdriven/src/internal/infrastructure/adapters"
	"eventdriven/src/internal/infrastructure/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"github.com/gin-contrib/cors"

)

func RegisterPedidosRoutes(router *gin.Engine) {

	dbPedidos := adapters.NewMySQLPedidos()
	dpFoods := adapters.NewMySQLFoods()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://13.217.156.220"}, 
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	ByIdFoodUseCase := products.NewByIdFoodUseCase(dpFoods)
	CreateFoodUseCase := products.NewCreateFoodUseCase(dpFoods)
	ListFoodUseCase := products.NewListFoodUseCase(dpFoods)

	CreateOrderUseCase := useCases.NewCreateOrderUseCase(dbPedidos)

	rabbitAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Fatalf("Error al configurar RabbitMQ: %v", err)
	}

	notificationService := services.NewNotificationService(rabbitAdapter)

	CreatePedidoController := controllers.NewCreatePedidoController(CreateOrderUseCase, notificationService)
	CreateFoodController := controllers.NewCreateFoodController(CreateFoodUseCase)
	byIdFoodController := controllers.NewByIdFoodController(ByIdFoodUseCase)
   ListFoodController := controllers.NewListFoodController(ListFoodUseCase)

	router.POST("/pedidos", CreatePedidoController.Execute)
	
	router.POST("/foods", CreateFoodController.Execute)
	router.GET("/foods/:producto_id", byIdFoodController.Execute)
	router.GET("/foods", ListFoodController.Execute)
}

