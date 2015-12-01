package goshopify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	// API Headers
	AccessHeader = "X-Shopify-Access-Token"

	// API Endpoints
	FulfillmentsEndpoint        = "/admin/orders/%s/fulfillments.json" // POST and GET
	FulfillmentCompleteEndpoint = "/admin/orders/%s/fulfillments/%s/complete.json"
	OrdersCountEndpoint         = "/admin/orders/count.json"
	OrdersEndpoint              = "/admin/orders.json"
	OrderEndpoint               = "/admin/orders/%s.json"
	ProductEndpoint             = "/admin/products/%s.json"
	ProductsCountEndpoint       = "/admin/products/count.json"
	ProductsEndpoint            = "/admin/products.json"
	TransactionsEndpoint        = "/admin/orders/%s/transactions.json"
	VariantEndpoint             = "/admin/variants/%s.json" // PUT and GET
	CustomersCountEndpoint      = "/admin/customers/count.json"
	CustomersEndpoint           = "/admin/customers.json"

	MaxIdleConnections = 90
	TimeFormat         = "2007-12-31T19:00:00-05:00"
	MaxLimit           = int64(250)
)

type Shopify struct {
	baseUri string
}

type Credentials struct {
	StoreId    string
	OauthToken string
}

func (c *Credentials) SignBaseUri(uri string) string {
	return strings.Split(fmt.Sprintf(uri, c.StoreId), "%!")[0]
}

func (c *Credentials) SignRequest(r *http.Request) {
	r.Header.Add(AccessHeader, c.OauthToken)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
}

func New(baseUri string) *Shopify {
	return &Shopify{baseUri}
}

func (s *Shopify) GetUri(endpoint string, creds *Credentials, params url.Values) (*url.URL, error) {
	baseUri := creds.SignBaseUri(s.baseUri)

	uri, err := url.Parse(baseUri)
	if err != nil {
		return nil, err
	}

	uri.Path = endpoint
	uri.RawQuery = params.Encode()

	return uri, nil
}

func (s *Shopify) DoRequest(verb string, uri *url.URL, creds *Credentials, payload []byte, jsonStruct interface{}) error {
	// Make sure the verb is uppercased
	verb = strings.ToUpper(verb)

	// Make sure characters are encoded
	uri.RawQuery = uri.Query().Encode()

	req, err := http.NewRequest(verb, uri.String(), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	creds.SignRequest(req)

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Invalid response from Shopify Request : %s", err.Error())
	}
	defer resp.Body.Close()

	// Is status code outside of the 200s?
	if resp.StatusCode < 200 || resp.StatusCode > 299 {

		// check if we are being throttled
		if resp.StatusCode == 429 {

			// read X-Retry-After header for sleep interval
			sleepInterval, err := strconv.Atoi(resp.Header.Get("X-Retry-After"))
			if err != nil {
				return fmt.Errorf("GetResponse cannot read sleep interval: (%d) %s", resp.StatusCode, err.Error())
			}

			// sleep
			time.Sleep(time.Duration(sleepInterval) * time.Second)

			// retry
			return s.DoRequest(verb, uri, creds, payload, jsonStruct)
		}

		// return the error
		return fmt.Errorf("Bad response code : (%d) %s", resp.StatusCode, uri)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Unable to read body from doResponse call: %s", err.Error())
	}

	if len(bodyBytes) > 0 {
		err = json.Unmarshal(bodyBytes, jsonStruct)
		if err != nil {
			return fmt.Errorf("Unmarshal failed: %s", err.Error())
		}
	}

	return nil
}
