package core

import "errors"

type OrderService interface {
	CreateOrder(order Order) error
}

type orderServiceImpl struct {
	repo OrderRepository
}

func newOrderService(repo OrderRepository) OrderService {
	return &orderServiceImpl{repo: repo}
}

func (s *orderServiceImpl) CreateOrder(order Order) error {
	if order.Total < 0 {
		return errors.New("total cannot be negative")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}
	return nil
}
