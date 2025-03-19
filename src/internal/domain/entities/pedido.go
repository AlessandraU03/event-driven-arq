package entities

type Pedido struct {
	ID         int     `json:"id"`
	ProductoID int     `json:"producto_id" binding:"required"`
	Cliente    string  `json:"cliente" binding:"required"`
	Direccion  string  `json:"direccion" binding:"required"`
	MetodoPago string  `json:"metodo_pago" binding:"required"`
	Monto      float64 `json:"monto_pagado" binding:"required"`
	Estado     string  `json:"estado" binding:"required"`
}

func NewPedido(pedidoID int, productoID int, cliente string, direccion string, metodoPago string, monto float64, estado string) *Pedido {
	return &Pedido{
		ID:         pedidoID,
		ProductoID: productoID,
		Cliente:    cliente,
		Direccion:  direccion,
		MetodoPago: metodoPago,
		Monto:      monto,
		Estado:     estado,
	}
}
