
package services

import (
	"eventdriven/src/internal/domain/entities"
	"eventdriven/src/internal/domain/repositories"
	"log"
)

type NotificationService struct {
	notificationPort repositories.NotificationPort
}

func NewNotificationService(notificationPort repositories.NotificationPort) *NotificationService {
	return &NotificationService{notificationPort: notificationPort}
}

func (s *NotificationService) NotifyPedidoCreation(pedido *entities.Pedido) error {
	err := s.notificationPort.NotifyPedidoCreation(pedido)
	if err != nil {
		log.Printf("Error al enviar notificación: %v", err)
		return err
	}
	log.Printf("Notificación enviada para el pedido %d", pedido.ID)
	return nil
}
