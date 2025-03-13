package repositories

import "eventdriven/src/internal/domain/entities"

type IFood interface {
	Save(Food *entities.Food) error
	GetAll() ([]*entities.Food, error) 
	GetById(FoodID int32) ([]*entities.Food, error)
	Update(Food *entities.Food) error
	Delete(FoodID int32) error

}