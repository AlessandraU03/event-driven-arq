package routes

import (
	"eventdriven/src/internal/application/useCases"
	"eventdriven/src/internal/application/useCases/products"
	"eventdriven/src/internal/application/services"
	"eventdriven/src/internal/infrastructure/adapters"
	"eventdriven/src/internal/infrastructure/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/gin-contrib/cors"

)

func RegisterPedidosRoutes(router *gin.Engine) {

	dbPedidos := adapters.NewMySQLPedidos()
	dpFoods := adapters.NewMySQLFoods()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	ByIdFoodUseCase := products.NewByIdFoodUseCase(dpFoods)
	CreateFoodUseCase := products.NewCreateFoodUseCase(dpFoods)
	ListFoodUseCase := products.NewListFoodUseCase(dpFoods)
	DeleteFoodUseCase := products.NewDeleteFoodUseCase(dpFoods)


	rabbitAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Fatalf("Error al configurar RabbitMQ: %v", err)
	}

	notificationService := services.NewNotificationService(rabbitAdapter)
	CreateOrderUseCase := useCases.NewCreateOrderUseCase(dbPedidos, notificationService)
	CreatePedidoController := controllers.NewCreatePedidoController(CreateOrderUseCase)


	CreateFoodController := controllers.NewCreateFoodController(CreateFoodUseCase)
	byIdFoodController := controllers.NewByIdFoodController(ByIdFoodUseCase)
   ListFoodController := controllers.NewListFoodController(ListFoodUseCase)
   DeleteFoodController := controllers.NewDeleteFoodController(DeleteFoodUseCase)

   router.POST("/pedidos", CreatePedidoController.Execute)
	router.DELETE("/foods/:producto_id", DeleteFoodController.Execute)
	router.POST("/foods", CreateFoodController.Execute)
	router.GET("/foods/:producto_id", byIdFoodController.Execute)
	router.GET("/foods", ListFoodController.Execute)
}

