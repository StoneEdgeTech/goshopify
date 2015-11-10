package goshopify

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestShopifyTransactions(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Transactions", func() {
		var mockShopify *bogus.Bogus

		g.BeforeEach(func() {
			mockShopify = bogus.New()
			mockShopify.Start()
		})

		g.AfterEach(func() {
			mockShopify.Close()
		})

		g.Describe("Get All", func() {
			g.It("should download a list of transactions", func() {
				mockShopify.SetPayload([]byte(SampleTransactionsJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-id", "oauthom"}

				transactions, err := s.GetTransactions("some-order-id", c, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(transactions).NotTo(BeNil())

				Expect(len(transactions)).To(Equal(3))
				t := transactions[0]
				Expect(t).NotTo(BeNil())
				Expect(t.Id).To(Equal(int64(179259969)))
				Expect(t.OrderId).To(Equal(int64(450789469)))
				Expect(t.Kind).To(Equal("refund"))
				Expect(t.Gateway).To(Equal("bogus"))
				Expect(t.Message).To(Equal(""))
				Expect(t.CreatedAt).To(Equal("2005-08-05T12:59:12-04:00"))
				Expect(t.Test).To(Equal(false))
				Expect(t.Authorization).To(Equal("authorization-key"))
				Expect(t.Status).To(Equal("success"))
				Expect(t.Amount).To(Equal("209.00"))
				Expect(t.Currency).To(Equal("USD"))
				Expect(t.LocationId).To(Equal(""))
				Expect(t.UserId).To(Equal(int64(0)))
				Expect(t.ParentId).To(Equal(int64(0)))
				Expect(t.DeviceId).To(Equal(int64(0)))
				Expect(t.Receipt.Authorization).To(Equal(""))
				Expect(t.Receipt.TestCase).To(Equal(false))
				Expect(t.ErrorCode).To(Equal(""))
				Expect(t.SourceName).To(Equal("web"))
			})
		})
	})
}

const (
	SampleTransactionsJson = `{
  "transactions": [
    {
      "id": 179259969,
      "order_id": 450789469,
      "kind": "refund",
      "gateway": "bogus",
      "message": null,
      "created_at": "2005-08-05T12:59:12-04:00",
      "test": false,
      "authorization": "authorization-key",
      "status": "success",
      "amount": "209.00",
      "currency": "USD",
      "location_id": null,
      "user_id": null,
      "parent_id": null,
      "device_id": null,
      "receipt": {},
      "error_code": null,
      "source_name": "web"
    },
    {
      "id": 389404469,
      "order_id": 450789469,
      "kind": "authorization",
      "gateway": "bogus",
      "message": null,
      "created_at": "2005-08-01T11:57:11-04:00",
      "test": false,
      "authorization": "authorization-key",
      "status": "success",
      "amount": "409.94",
      "currency": "USD",
      "location_id": null,
      "user_id": null,
      "parent_id": null,
      "device_id": null,
      "receipt": {
        "testcase": true,
        "authorization": "123456"
      },
      "error_code": null,
      "source_name": "web",
      "payment_details": {
        "credit_card_bin": null,
        "avs_result_code": null,
        "cvv_result_code": null,
        "credit_card_number": "•••• •••• •••• 4242",
        "credit_card_company": "Visa"
      }
    },
    {
      "id": 801038806,
      "order_id": 450789469,
      "kind": "capture",
      "gateway": "bogus",
      "message": null,
      "created_at": "2005-08-05T10:22:51-04:00",
      "test": false,
      "authorization": "authorization-key",
      "status": "success",
      "amount": "250.94",
      "currency": "USD",
      "location_id": null,
      "user_id": null,
      "parent_id": null,
      "device_id": null,
      "receipt": {},
      "error_code": null,
      "source_name": "web"
    }
  ]
}`
)
