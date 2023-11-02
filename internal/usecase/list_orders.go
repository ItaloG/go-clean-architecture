package usecase

import "github.com/ItaloG/go-clean-architecture/internal/entity"

type ListOrdersOutputDTO struct {
	Orders []entity.Order `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (l *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {
	orders, err := l.OrderRepository.ListAll()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}

	dto := ListOrdersOutputDTO{Orders: orders}

	return dto, nil
}
