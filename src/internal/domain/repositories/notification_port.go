package repositories

import "eventdriven/src/internal/domain/entities"

type NotificationPort interface {
	NotifyPedidoCreation(pedido *entities.Pedido) error
}
