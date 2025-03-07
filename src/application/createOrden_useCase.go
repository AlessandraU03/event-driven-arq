package application

import "eventdriven/src/domain/repositories"



type CreateOrderUseCase struct {
	orderRepository repositories.IOrder	
} 

func NewCreateOrderUseCase(orderRepository repositories.IOrder) *CreateOrderUseCase {
	return &CreateOrderUseCase{orderRepository: orderRepository}
}

func (useCase *CreateOrderUseCase) Execute(product string, quantity int, price float64) {	
	useCase.orderRepository.Save(product, quantity, price)
}