package controllers

import (
	"eventdriven/src/internal/application/useCases"
	"eventdriven/src/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreatePedidoController struct {
	useCase *useCases.CreateOrderUseCase
}

func NewCreatePedidoController(useCase *useCases.CreateOrderUseCase) *CreatePedidoController {
	return &CreatePedidoController{
		useCase: useCase,
	}
}

func (c *CreatePedidoController) Execute(g *gin.Context) {
	var newPedido entities.Pedido

	if err := g.ShouldBindJSON(&newPedido); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Ejecutar el caso de uso
	_, err := c.useCase.Execute(&newPedido)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el pedido"})
		return
	}

	g.JSON(http.StatusCreated, gin.H{
		"message":   "Pedido creado con éxito",
		"pedido_id": newPedido.ID,
	})
}
