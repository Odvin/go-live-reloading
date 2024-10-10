package application

import (
	"github.com/Odvin/go-live-reloading/internal/application/domain"
	"github.com/Odvin/go-live-reloading/internal/ports"
)

type API interface {
	CreateCustomer(title string) (*domain.Customer, error)
	GetCustomer(id int64) (*domain.Customer, error)
}

type Application struct {
	db ports.DB
}

func Init(db ports.DB) *Application {
	return &Application{
		db: db,
	}
}
