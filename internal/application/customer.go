package application

import "github.com/Odvin/go-live-reloading/internal/application/domain"

func (app *Application) CreateCustomer(title string) (*domain.Customer, error) {
	v := domain.NewValidator()

	customer := &domain.Customer{Title: title, Active: true}

	if domain.ValidateCustomer(v, customer); !v.Valid() {
		return nil, nil
	}

	err := app.db.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (app *Application) GetCustomer(id int64) (*domain.Customer, error) {
	customer, err := app.db.GetCustomer(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
