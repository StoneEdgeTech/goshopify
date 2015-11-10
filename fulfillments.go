package shopify

import (
	"errors"
	"fmt"
	"net/url"
)

type FulfillmentsResponse struct {
	Fulfillments []*Fulfillment `json:"fulfillments"`
}

type Fulfillment struct {
	CreatedAt       string      `json:"created_at"`
	Id              int64       `json:"id"`
	OrderId         int64       `json:"order_id"`
	Service         string      `json:"service"`
	Status          string      `json:"status"`
	TrackingCompany string      `json:"tracking_company"`
	UpdatedAt       string      `json:"updated_at"`
	TrackingNumber  string      `json:"tracking_number"`
	TrackingNumbers []string    `json:"tracking_numbers"`
	TrackingUrl     string      `json:"tracking_url"`
	TrackingUrls    []string    `json:"tracking_urls"`
	Receipt         *Receipt    `json:"receipt"`
	LineItems       []*LineItem `json:"line_items"`
}

// call shopify to create fulfillment
func (s *Shopify) CreateFulfillment(fulfillmentJson, orderId string, creds *Credentials, params url.Values) (*Fulfillment, error) {
	uri, err := s.getUri(fmt.Sprintf(FulfillmentsEndpoint, orderId), creds, params)
	if err != nil {
		return nil, err
	}

	var fulfillmentsResponse *FulfillmentsResponse
	err = s.DoResponse("POST", uri, creds, []byte(fulfillmentJson), &fulfillmentsResponse)
	if err != nil {
		return nil, fmt.Errorf("Request to Shopify Create Fulfillment failed: %s", err.Error())
	}

	if len(fulfillmentsResponse.Fulfillments) < 1 {
		return nil, errors.New("CreateFulfillment returned with no fulfillment response.")
	}
	return fulfillmentsResponse.Fulfillments[0], nil
}

// call shopify to ship fulfillment (mark as complete)
func (s *Shopify) CompleteFulfillment(orderId, shipmentId string, creds *Credentials, params url.Values) (*Fulfillment, error) {
	uri, err := s.getUri(fmt.Sprintf(FulfillmentCompleteEndpoint, orderId, shipmentId), creds, params)
	if err != nil {
		return nil, err
	}

	var fulfillmentResponse *Fulfillment
	err = s.DoResponse("POST", uri, creds, []byte(""), &fulfillmentResponse)
	if err != nil {
		return nil, fmt.Errorf("Request to Shopify Complete Fulfillment failed: %s", err.Error())
	}

	return fulfillmentResponse, nil
}

// get all the fulfillments for a specific order
func (s *Shopify) GetFulfillments(orderId string, creds *Credentials, params url.Values) ([]*Fulfillment, error) {
	uri, err := s.getUri(fmt.Sprintf(FulfillmentsEndpoint, orderId), creds, params)
	if err != nil {
		return nil, err
	}

	var fulfillmentsResponse *FulfillmentsResponse
	err = s.DoResponse("GET", uri, creds, nil, &fulfillmentsResponse)
	if err != nil {
		return nil, fmt.Errorf("Request to Shopify Orders failed: %s", err.Error())
	}

	if fulfillmentsResponse == nil {
		return nil, nil
	}

	return fulfillmentsResponse.Fulfillments, nil
}
