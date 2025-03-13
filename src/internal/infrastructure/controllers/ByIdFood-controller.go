package controllers

import (
	application "eventdriven/src/internal/application/useCases/products"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ByIdFoodController struct {
	useCase *application.ByIdFoodUseCase
}

func NewByIdFoodController(useCase *application.ByIdFoodUseCase) *ByIdFoodController {
	return &ByIdFoodController{useCase: useCase}
}

func (c *ByIdFoodController) Execute(g *gin.Context) {
    idStr := g.Param("producto_id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    // Llamamos al caso de uso para obtener los datos de la comida
    food, err := c.useCase.Execute(int32(id))
    if err != nil {
        g.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
        return
    }

    // Respondemos con los datos de la comida obtenidos
    g.JSON(http.StatusOK, food)
}
