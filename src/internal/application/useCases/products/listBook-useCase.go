package products

import (
	"eventdriven/src/internal/domain/entities"
	"eventdriven/src/internal/domain/repositories"
	"errors"
)

type ListFoodUseCase struct {
	lf repositories.IFood
}

func NewListFoodUseCase(lf repositories.IFood) *ListFoodUseCase {
	return &ListFoodUseCase{lf: lf}
}

// ✅ Ahora devuelve un slice de Food y un error
func (useCase *ListFoodUseCase) Execute() ([]*entities.Food, error) {
	foods, err := useCase.lf.GetAll()
	if err != nil {
		return nil, errors.New("failed to retrieve foods")
	}
	return foods, nil
}
