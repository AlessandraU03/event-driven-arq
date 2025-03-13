package controllers

import (
	"eventdriven/src/internal/application/useCases/products"
	"eventdriven/src/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateFoodController struct {
	useCase *products.CreateFoodUseCase
}

func NewCreateFoodController(useCase *products.CreateFoodUseCase) *CreateFoodController {
	return &CreateFoodController{useCase: useCase}
}

func (c *CreateFoodController) Execute(g *gin.Context) {
	var newFood struct {
		ID      int32   `json:"producto_id"`
		Nombre        string  `json:"nombre"`
		Precio      float64 `json:"precio"`      
		Descripcion string  `json:"descripcion"`
	}

	if err := g.ShouldBindJSON(&newFood); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	food := &entities.Food{
		ID: int(newFood.ID),
		Nombre:        newFood.Nombre,
		Precio:      newFood.Precio,
		Descripcion: newFood.Descripcion,
	}

    c.useCase.Execute(food)
 
	g.JSON(http.StatusCreated, gin.H{
		"message":     "Comida creada con éxito",
		"producto_id": food.ID,
		"name":        food.Nombre,
		"precio":      food.Precio,
		"descripcion": food.Descripcion,
	})
}
