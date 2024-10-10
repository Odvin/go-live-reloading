package ports

import "github.com/Odvin/go-live-reloading/internal/application/domain"

type DB interface {
	CreateCustomer(customer *domain.Customer) error
	GetCustomer(id int64) (*domain.Customer, error)
}
