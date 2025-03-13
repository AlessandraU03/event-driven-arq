package products

import (
	"eventdriven/src/internal/domain/entities"
	"eventdriven/src/internal/domain/repositories"
)



type DeleteFoodUseCase struct {
	df repositories.IFood
} 

func NewDeleteFoodUseCase(df repositories.IFood) *DeleteFoodUseCase {
	return &DeleteFoodUseCase{df: df}
}

func (useCase *DeleteFoodUseCase) Execute(food *entities.Food) {
	useCase.df.Delete(int32(food.ID))
}