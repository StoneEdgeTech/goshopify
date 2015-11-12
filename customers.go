package goshopify

import (
	"fmt"
	"net/url"
)

type CustomersCount struct {
	Count int64 `json:"count"`
}

type CustomersResponse struct {
	Customers []*Customer `json:"customers"`
}

type Customer struct {
	AcceptsMarketing bool               `json:"accepts_marketing"`
	CreatedAt        string             `json:"created_at"`
	Email            string             `json:"email"`
	FirstName        string             `json:"first_name"`
	Id               int64              `json:"id"`
	LastName         string             `json:"last_name"`
	LastOrderId      int64              `json:"last_order_id"`
	MultipassId      string             `json:"multipass_identifier"`
	Note             string             `json:"note"`
	OrdersCount      int64              `json:"orders_count"`
	State            string             `json:"state"`
	TaxExempt        bool               `json:"tax_exempt"`
	TotalSpent       string             `json:"total_spent"`
	UpdatedAt        string             `json:"updated_at"`
	VerifiedEmail    bool               `json:"verified_email"`
	Tags             string             `json:"tags"`
	LastOrderName    string             `json:"last_order_name"`
	DefaultAddress   *CustomerAddress   `json:"default_address"`
	Addresses        []*CustomerAddress `json:"addresses"`
}

func (s *Shopify) GetCustomers(creds *Credentials, params url.Values) ([]*Customer, error) {
	uri, err := s.getUri(CustomersEndpoint, creds, params)
	if err != nil {
		return nil, err
	}

	var customersResponse *CustomersResponse
	err = s.DoResponse("GET", uri, creds, nil, &customersResponse)
	if err != nil {
		return nil, fmt.Errorf("Customers: %s", err.Error())
	}

	return customersResponse.Customers, nil
}

func (s *Shopify) GetCustomersCount(creds *Credentials, params url.Values) (*CustomersCount, error) {
	uri, err := s.getUri(CustomersCountEndpoint, creds, params)
	if err != nil {
		return nil, err
	}

	var customersCount CustomersCount
	err = s.DoResponse("GET", uri, creds, nil, &customersCount)
	if err != nil {
		return nil, fmt.Errorf("CustomersCount: %s", err.Error())
	}

	return &customersCount, nil
}
