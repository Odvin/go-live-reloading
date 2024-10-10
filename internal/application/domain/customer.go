package domain

import "time"

type Customer struct {
	ID      int64     `json:"id"`
	Title   string    `json:"title"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Active  bool      `json:"active"`
	Version int32     `json:"version"`
}

func ValidateCustomer(v *validator, customer *Customer) {
	v.check(customer.Title != "", "title", "must be provided")
	v.check(len(customer.Title) <= 200, "title", "must not be more then 200 bytes long")
}
