package controllers

import (
	"eventdriven/src/internal/application/useCases/products"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListFoodController struct {
	useCase *products.ListFoodUseCase
}

func NewListFoodController(useCase *products.ListFoodUseCase) *ListFoodController {
	return &ListFoodController{useCase: useCase}
}

func (c *ListFoodController) Execute(g *gin.Context) {
	foods, err := c.useCase.Execute()
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
		return
	}

	g.JSON(http.StatusOK, foods)
}
