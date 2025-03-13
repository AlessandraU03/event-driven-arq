package repositories

import "eventdriven/src/internal/domain/entities"

type IPedido interface {
	Save(pedido *entities.Pedido) error

}