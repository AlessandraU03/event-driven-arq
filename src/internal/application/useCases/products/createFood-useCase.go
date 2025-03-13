package products

import (
	"eventdriven/src/internal/domain/entities"
	"eventdriven/src/internal/domain/repositories"
)



type CreateFoodUseCase struct {
	fr repositories.IFood
} 

func NewCreateFoodUseCase(fr repositories.IFood) *CreateFoodUseCase {
	return &CreateFoodUseCase{fr: fr}
}

func (useCase *CreateFoodUseCase) Execute(food *entities.Food) {
	useCase.fr.Save(food)
}