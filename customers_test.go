package goshopify

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestShopifyCustomers(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Customers", func() {
		var mockShopify *bogus.Bogus

		g.BeforeEach(func() {
			mockShopify = bogus.New()
			mockShopify.Start()
		})

		g.AfterEach(func() {
			mockShopify.Close()
		})

		g.Describe("Get All", func() {
			g.It("should download a list of customers", func() {
				mockShopify.SetPayload([]byte(SampleCustomersJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-id", "oauthom"}

				custs, err := s.GetCustomers(c, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(custs).NotTo(BeNil())
				Expect(len(custs)).To(Equal(1))
				Expect(custs[0].Id).To(Equal(int64(207119551)))
			})
		})

		g.Describe("Count", func() {
			g.It("should get an customers count", func() {
				mockShopify.SetPayload([]byte(SampleCustomersCountJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-id", "oauthom"}

				count, err := s.GetCustomersCount(c, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(count.Count).To(Equal(int64(33)))
			})
		})
	})
}

const (
	SampleCustomersJson = `{
  "customers": [
    {
      "id": 207119551,
      "email": "bob.norman@hostmail.com",
      "accepts_marketing": false,
      "created_at": "2015-09-02T14:52:18-04:00",
      "updated_at": "2015-09-02T14:52:18-04:00",
      "first_name": "Bob",
      "last_name": "Norman",
      "orders_count": 1,
      "state": "disabled",
      "total_spent": "41.94",
      "last_order_id": 450789469,
      "note": null,
      "verified_email": true,
      "multipass_identifier": null,
      "tax_exempt": false,
      "tags": "",
      "last_order_name": "#1001",
      "default_address": {
        "id": 207119551,
        "first_name": null,
        "last_name": null,
        "company": null,
        "address1": "Chestnut Street 92",
        "address2": "",
        "city": "Louisville",
        "province": "Kentucky",
        "country": "United States",
        "zip": "40202",
        "phone": "555-625-1199",
        "name": "",
        "province_code": "KY",
        "country_code": "US",
        "country_name": "United States",
        "default": true
      },
      "addresses": [
        {
          "id": 207119551,
          "first_name": null,
          "last_name": null,
          "company": null,
          "address1": "Chestnut Street 92",
          "address2": "",
          "city": "Louisville",
          "province": "Kentucky",
          "country": "United States",
          "zip": "40202",
          "phone": "555-625-1199",
          "name": "",
          "province_code": "KY",
          "country_code": "US",
          "country_name": "United States",
          "default": true
        }
      ]
    }
  ]
}`

	SampleCustomersCountJson = `{"count": 33}`
)
