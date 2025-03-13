package controllers

import (
	"eventdriven/src/internal/application/useCases"
	"eventdriven/src/internal/application/services" 
	"eventdriven/src/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreatePedidoController struct {
	useCase             *useCases.CreateOrderUseCase
	notificationService *services.NotificationService
}

func NewCreatePedidoController(useCase *useCases.CreateOrderUseCase, notificationService *services.NotificationService) *CreatePedidoController {
	return &CreatePedidoController{
		useCase:             useCase,
		notificationService: notificationService,
	}
}

type PedidoRequest struct {
	ProductoID  int     `json:"producto_id" binding:"required"` 
	Cliente     string  `json:"cliente" binding:"required"`     
	Direccion   string  `json:"direccion" binding:"required"`
	MetodoPago  string  `json:"metodo_pago" binding:"required"` 
	Monto       float64 `json:"monto_pagado" binding:"required"`
	Estado      string  `json:"estado" binding:"required"` 
}

func (c *CreatePedidoController) Execute(g *gin.Context) {
	var newPedido PedidoRequest

	
	if err := g.ShouldBindJSON(&newPedido); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	pedido := &entities.Pedido{
		ProductoID: newPedido.ProductoID,
		Cliente:    newPedido.Cliente,
		Direccion:  newPedido.Direccion,
		MetodoPago: newPedido.MetodoPago,
		Monto:      newPedido.Monto,
		Estado:     newPedido.Estado,
	}

	
	c.useCase.Execute(pedido)
	
	err := c.notificationService.NotifyPedidoCreation(pedido)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar la notificación"})
		return
	}

	g.JSON(http.StatusCreated, gin.H{
		"message":   "Pedido creado con éxito",
		"pedido_id": pedido.ID,
	})
}
