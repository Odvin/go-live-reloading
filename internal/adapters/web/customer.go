package web

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Odvin/go-live-reloading/internal/application/domain"
)

func (web *WebService) createCustomer(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}

	err := readJSON(w, r, &input)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	customer, err := web.api.CreateCustomer(input.Title)
	if err != nil {
		serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/customers/%d", customer.ID))

	err = writeJSON(w, http.StatusCreated, envelope{"customer": customer}, headers)
	if err != nil {
		serverErrorResponse(w, r, err)
	}
}

func (web *WebService) getCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r)
	if err != nil {
		notFoundResponse(w, r)
		return
	}

	customer, err := web.api.GetCustomer(id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrRecordNotFound):
			notFoundResponse(w, r)
		default:
			serverErrorResponse(w, r, err)
		}
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"movie": customer}, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}
}
