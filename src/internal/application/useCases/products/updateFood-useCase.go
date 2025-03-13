package products

import (
	"eventdriven/src/internal/domain/entities"
	"eventdriven/src/internal/domain/repositories"
)



type UpdateFoodUseCase struct {
	df repositories.IFood
} 

func NewUpdateFoodUseCase(df repositories.IFood) *UpdateFoodUseCase {
	return &UpdateFoodUseCase{df: df}
}

func (useCase *UpdateFoodUseCase) Execute(food *entities.Food) {
	useCase.df.Update(food)
}