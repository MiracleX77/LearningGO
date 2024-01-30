package usecases

import "mix/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
