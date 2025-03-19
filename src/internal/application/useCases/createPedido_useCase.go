package useCases

import (
	"eventdriven/src/internal/domain/entities"
	"eventdriven/src/internal/domain/repositories"
)

type CreateOrderUseCase struct {
	repo               repositories.IPedido
	notificationService repositories.NotificationPort
}

func NewCreateOrderUseCase(repo repositories.IPedido, notificationService repositories.NotificationPort) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		repo:               repo,
		notificationService: notificationService,
	}
}

func (uc *CreateOrderUseCase) Execute(pedido *entities.Pedido) (int, error) {
	err := uc.repo.Save(pedido)
	if err != nil {
		return 0, err
	}

	err = uc.notificationService.NotifyPedidoCreation(pedido)
	if err != nil {
		return 0, err
	}

	return pedido.ID, nil
}
