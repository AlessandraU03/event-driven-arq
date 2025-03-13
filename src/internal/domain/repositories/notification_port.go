package repositories

import "eventdriven/src/internal/domain/entities"

// NotificationPort es el puerto que define el método para enviar notificaciones
type NotificationPort interface {
	NotifyPedidoCreation(pedido *entities.Pedido) error
}
