package goshopify

import (
	"fmt"
	"net/url"
)

type TransactionsResponse struct {
	Transactions []*Transaction `json:"transactions"`
}

type Transaction struct {
	Amount         string          `json:"amount"`
	Authorization  string          `json:"authorization"`
	CreatedAt      string          `json:"created_at"`
	Currency       string          `json:"currency"`
	Gateway        string          `json:"gateway"`
	Id             int64           `json:"id"`
	Kind           string          `json:"kind"`
	LocationId     string          `json:"location_id"`
	Message        string          `json:"message"`
	OrderId        int64           `json:"order_id"`
	ParentId       int64           `json:"parent_id"`
	Status         string          `json:"status"`
	Test           bool            `json:"test"`
	UserId         int64           `json:"user_id"`
	DeviceId       int64           `json:"device_id"`
	Receipt        *Receipt        `json:"receipt"`
	ErrorCode      string          `json:"error_code"`
	SourceName     string          `json:"source_name"`
	PaymentDetails *PaymentDetails `json:"payment_details"`
}

type PaymentDetails struct {
	CreditCardBin     string `json:"credit_card_bin"`
	AvsResultCode     string `json:"avs_result_code"`
	CvvResultCode     string `json:"cvv_result_code"`
	CreditCardNumber  string `json:"credit_card_number"`
	CreditCardCompany string `json:"credit_card_company"`
}
type Receipt struct {
	TestCase      bool   `json:"testcase"`
	Authorization string `json:"authorization"`
}

func (s *Shopify) GetTransactions(orderId string, creds *Credentials, params url.Values) ([]*Transaction, error) {
	uri, err := s.getUri(fmt.Sprintf(TransactionsEndpoint, orderId), creds, params)
	if err != nil {
		return nil, err
	}

	var transactionsResponse *TransactionsResponse
	err = s.DoResponse("GET", uri, creds, nil, &transactionsResponse)
	if err != nil {
		return nil, fmt.Errorf("Request to Shopify Transactions failed: %s", err.Error())
	}

	return transactionsResponse.Transactions, nil
}
