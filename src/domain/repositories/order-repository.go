package repositories

type IOrder interface {
	Save(product string, quantity int, price float64)
}