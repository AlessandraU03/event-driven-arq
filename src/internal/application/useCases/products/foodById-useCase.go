package products

import (
	"eventdriven/src/internal/domain/entities"
	"eventdriven/src/internal/domain/repositories"
)



type ByIdFoodUseCase struct {
	byf repositories.IFood
} 

func NewByIdFoodUseCase(byf repositories.IFood) *ByIdFoodUseCase {
	return &ByIdFoodUseCase{byf: byf}
}

func (useCase *ByIdFoodUseCase) Execute(FoodID int32) (food []*entities.Food, err error) {
	return useCase.byf.GetById(FoodID)
}