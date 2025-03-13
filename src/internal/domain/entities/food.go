package entities

type Food struct {
	ID       int     `json:"producto_id"`
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
	Descripcion string     `json:"descripcion"`
}

func NewFood(FoodID int, nombre string, precio float64, descripcion string) *Food {
	return &Food{
		ID:       FoodID,
		Nombre:   nombre,
		Precio:   precio,
		Descripcion: descripcion,
	}
}
