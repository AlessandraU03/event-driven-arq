package controllers

import (
	application "eventdriven/src/internal/application/useCases/products"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteFoodController struct {
	useCase *application.DeleteFoodUseCase
}

func NewDeleteFoodController(useCase *application.DeleteFoodUseCase) *DeleteFoodController {
	return &DeleteFoodController{useCase: useCase}
}

func (c *DeleteFoodController) Execute(g *gin.Context) {
    idStr := g.Param("producto_id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    // Llamamos al caso de uso para eliminar la comida
    err = c.useCase.Execute(int32(id))
    if err != nil {
        g.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
        return
    }

    // Respondemos con un mensaje de éxito
    g.JSON(http.StatusOK, gin.H{"message": "Food deleted successfully"})
}
