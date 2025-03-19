package services

import (
	"eventdriven/src/internal/domain/entities"
	"eventdriven/src/internal/domain/repositories"
)

type NotificationService struct {
	notificationRepo repositories.NotificationPort
}

func NewNotificationService(notificationRepo repositories.NotificationPort) *NotificationService {
	return &NotificationService{
		notificationRepo: notificationRepo,
	}
}

func (s *NotificationService) NotifyPedidoCreation(pedido *entities.Pedido) error {
	return s.notificationRepo.NotifyPedidoCreation(pedido)
}
