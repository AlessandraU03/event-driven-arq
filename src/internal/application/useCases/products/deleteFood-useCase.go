package products

import (
	"eventdriven/src/internal/domain/repositories"
)


type DeleteFoodUseCase struct {
	df repositories.IFood
} 

func NewDeleteFoodUseCase(df repositories.IFood) *DeleteFoodUseCase {
	return &DeleteFoodUseCase{df: df}
}

func (useCase *DeleteFoodUseCase) Execute(FoodID int32) error {
	return useCase.df.Delete(FoodID) // Devuelve el error correctamente
}