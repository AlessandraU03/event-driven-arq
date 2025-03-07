package entities

type Order struct {
	id	int32
	product string
	quantity int
	price float64
}

func NewOrder( product string, quantity int, price float64) *Order {
	return &Order{
		id:0, product: product, quantity: quantity, price: price}
}
