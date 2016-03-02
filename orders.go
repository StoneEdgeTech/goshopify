package goshopify

import (
	"errors"
	"fmt"
	"net/url"
)

type OrdersCount struct {
	Count int64 `json:"count"`
}

type OrderResponse struct {
	Order *Order `json:"order"`
}

type OrdersResponse struct {
	Orders []*Order `json:"orders"`
}

type Order struct {
	BuyerAcceptsMarketing bool             `json:"buyer_accepts_marketing"`
	CancelReason          string           `json:"cancel_reason"`
	CancelledAt           string           `json:"cancelled_at"`
	CartToken             string           `json:"cart_token"`
	CheckoutToken         string           `json:"checkout_token"`
	ClosedAt              string           `json:"closed_at"`
	Confirmed             bool             `json:"confirmed"`
	CreatedAt             string           `json:"created_at"`
	Currency              string           `json:"currency"`
	DevicdId              int64            `json:"device_id"`
	Email                 string           `json:"email"`
	FinancialStatus       string           `json:"financial_status"`
	FulfillmentStatus     string           `json:"fulfillment_status"`
	Gateway               string           `json:"gateway"`
	Id                    int64            `json:"id"`
	LandingSite           string           `json:"landing_site"`
	LocationId            int64            `json:"location_id"`
	Name                  string           `json:"name"`
	Note                  string           `json:"note"`
	Number                int64            `json:"number"`
	ProcessedAt           string           `json:"processed_at"`
	Reference             string           `json:"reference"`
	ReferringSite         string           `json:"referring_site"`
	SourceIdentifier      string           `json:"source_identifier"`
	SourceUrl             string           `json:"source_url"`
	SubtotalPrice         string           `json:"subtotal_price"`
	TaxesIncluded         bool             `json:"taxes_included"`
	Test                  bool             `json:"test"`
	Token                 string           `json:"token"`
	TotalDiscount         string           `json:"total_discounts"`
	TotalLineItemsPrice   string           `json:"total_line_items_price"`
	TotalPrice            string           `json:"total_price"`
	TotalPriceUsd         string           `json:"total_price_usd"`
	TotalTax              string           `json:"total_tax"`
	TotalWeight           int64            `json:"total_weight"`
	UpdatedAt             string           `json:"updated_at"`
	UserId                int64            `json:"user_id"`
	BrowserIp             string           `json:"browser_ip"`
	LandingSiteRef        string           `json:"landing_site_ref"`
	OrderNumber           int64            `json:"order_number"`
	DiscountCodes         []*DiscountCode  `json:"discount_codes"`
	NoteAttributes        []*NoteAttribute `json:"note_attributes"`
	PaymentGatewayNames   []*string        `json:"payment_gateway_names"`
	ProcessingMethod      string           `json:"processing_method"`
	Source                string           `json:"source"`
	CheckoutId            int64            `json:"checkout_id"`
	SourceName            string           `json:"source_name"`
	TaxLines              []*TaxLine       `json:"tax_lines"`
	Tags                  string           `json:"tags"`
	LineItems             []*LineItem      `json:"line_items"`
	ShippingLines         []*ShippingLine  `json:"shipping_lines"`
	BillingAddress        *Address         `json:"billing_address"`
	ShippingAddress       *Address         `json:"shipping_address"`
	Fulfillments          []*Fulfillment   `json:"fulfillments"`
	ClientDetails         *ClientDetails   `json:"client_details"`
	Refunds               []*Refund        `json:"refunds"`
	Customer              *Customer        `json:"customer"`
}

type DiscountCode struct {
	Code   string `json:"code"`
	Amount string `json:"amount"`
	Type   string `json:"type"`
}

type NoteAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TaxLine struct {
	Price string  `json:"price"`
	Rate  float64 `json:"rate"`
	Title string  `json:"title"`
}

type LineItem struct {
	FulfillmentService         string      `json:"fulfillment_service"`
	FulfillmentStatus          string      `json:"fulfillment_status"`
	GiftCard                   bool        `json:"gift_card"`
	Grams                      float64     `json:"grams"`
	Id                         int64       `json:"id"`
	Price                      string      `json:"price"`
	ProductId                  int64       `json:"product_id"`
	Quantity                   int64       `json:"quantity"`
	RequiresShipping           bool        `json:"requires_shipping"`
	Sku                        string      `json:"sku"`
	Taxable                    bool        `json:"taxable"`
	Title                      string      `json:"title"`
	VariantId                  int64       `json:"variant_id"`
	VariantTitle               string      `json:"variant_title"`
	Vendor                     string      `json:"vendor"`
	Name                       string      `json:"name"`
	VariantInventoryManagement string      `json:"variant_inventory_management"`
	Properties                 []*Property `json:"properties"`
	ProductExists              bool        `json:"product_exists"`
	FulfillableQuantity        int64       `json:"fulfillable_quantity"`
	TotalDiscount              string      `json:"total_discount"`
	TaxLines                   []*TaxLine  `json:"tax_lines"`
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ShippingLine struct {
	Code     string     `json:"code"`
	Price    string     `json:"price"`
	Source   string     `json:"source"`
	Title    string     `json:"title"`
	TaxLines []*TaxLine `json:"tax_lines"`
}

type Address struct {
	Address1     string  `json:"address1"`
	Address2     string  `json:"address2"`
	City         string  `json:"city"`
	Company      string  `json:"company"`
	Country      string  `json:"country"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Phone        string  `json:"phone"`
	Province     string  `json:"province"`
	Zip          string  `json:"zip"`
	Name         string  `json:"name"`
	CountryCode  string  `json:"country_code"`
	ProvinceCode string  `json:"province_code"`
}

type ClientDetails struct {
	AcceptLanguage string `json:"accept_language"`
	BrowserHeight  int64  `json:"browser_height"`
	BrowserIp      string `json:"browser_ip"`
	BrowserWidth   int64  `json:"browser_width"`
	SessionHash    string `json:"session_hash"`
	UserAgent      string `json:"user_agent"`
}

type Refund struct {
	CreatedAt       string            `json:"created_at"`
	Id              int64             `json:"id"`
	Note            string            `json:"note"`
	OrderId         int64             `json:"order_id"`
	Restock         bool              `json:"restock"`
	UserId          int64             `json:"user_id"`
	RefundLineItems []*RefundLineItem `json:"refund_line_items"`
	Transactions    []*Transaction    `json:"transactions"`
}

type RefundLineItem struct {
	Id         int64     `json:"id"`
	LineItemId int64     `json:"line_item_id"`
	Quantity   int64     `json:"quantity"`
	LineItem   *LineItem `json:"line_item"`
}

func (s *Shopify) GetOrdersCount(creds *Credentials, params url.Values) (*OrdersCount, error) {
	uri, err := s.getUri(OrdersCountEndpoint, creds, params)
	if err != nil {
		return nil, err
	}

	var count *OrdersCount
	err = s.DoRequest("GET", uri, creds, nil, &count)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("OrdersCount: %s", err.Error()))
	}

	return count, nil
}

func (s *Shopify) GetOrders(creds *Credentials, params url.Values) ([]*Order, error) {
	uri, err := s.getUri(OrdersEndpoint, creds, params)
	if err != nil {
		return nil, err
	}

	var ordersResponse *OrdersResponse
	err = s.DoRequest("GET", uri, creds, nil, &ordersResponse)
	if err != nil {
		return nil, fmt.Errorf("Request to Shopify Orders failed: %s", err.Error())
	}

	return ordersResponse.Orders, nil
}

func (s *Shopify) GetOrdersSinceId(creds *Credentials, sinceId string, params url.Values) ([]*Order, error) {
	if params == nil {
		params = url.Values{}
	}
	params.Add("since_id", sinceId)
	return s.GetOrders(creds, params)
}

func (s *Shopify) GetOrdersSinceOrderNumber(creds *Credentials, sinceNum string, params url.Values) ([]*Order, error) {
	sinceId, err := s.FindOrderIdFromOrderNumber(creds, sinceNum)
	if err != nil {
		return nil, err
	}
	return s.GetOrdersSinceId(creds, sinceId, params)
}

func (s *Shopify) FindOrderIdFromOrderNumber(creds *Credentials, orderNum string) (string, error) {
	params := url.Values{}
	params.Add("name", orderNum)
	params.Add("status", "any")
	params.Add("fields", "id,order_number")
	params.Add("limit", fmt.Sprintf("%v", MaxLimit))
	uri, err := s.getUri(OrdersEndpoint, creds, params)
	var ordersResponse *OrdersResponse
	err = s.DoRequest("GET", uri, creds, nil, &ordersResponse)
	if err != nil {
		return "", fmt.Errorf("Request to Shopify Orders failed: %s", err.Error())
	}
	if ordersResponse == nil || ordersResponse.Orders == nil || len(ordersResponse.Orders) != 1 {
		return "", fmt.Errorf("invalid response from shopify when resolving orders")
	}
	return fmt.Sprintf("%v", ordersResponse.Orders[0].Id), nil
}

func (s *Shopify) FindOrderNumberFromOrderId(creds *Credentials, orderId string) (int64, error) {
	params := url.Values{}
	params.Add("fields", "order_number")
	order, err := s.GetOrder(orderId, creds, params)
	if err != nil {
		return 0, err
	}
	if order == nil {
		return 0, nil
	}
	return order.OrderNumber, nil
}

func (s *Shopify) GetOrder(orderId string, creds *Credentials, params url.Values) (*Order, error) {
	uri, err := s.getUri(fmt.Sprintf(OrderEndpoint, orderId), creds, params)
	if err != nil {
		return nil, err
	}
	var orderResponse *OrderResponse
	err = s.DoRequest("GET", uri, creds, nil, &orderResponse)
	if err != nil {
		return nil, fmt.Errorf("Request to Shopify Get Order failed: %s", err.Error())
	}
	// if no order returned, then just return nil. Do not error out.
	if orderResponse == nil {
		return nil, nil
	}
	return orderResponse.Order, nil
}
