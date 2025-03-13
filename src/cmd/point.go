package cmd

import (
	"eventdriven/src/internal/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func Point() {
	router := gin.Default()

	routes.RegisterPedidosRoutes(router)

	router.Run(":8080")
}