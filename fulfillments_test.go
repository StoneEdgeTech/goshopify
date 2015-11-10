package goshopify

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestShopifyFullfillments(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Fullfillments", func() {
		var mockShopify *bogus.Bogus

		g.BeforeEach(func() {
			mockShopify = bogus.New()
			mockShopify.Start()
		})

		g.AfterEach(func() {
			mockShopify.Close()
		})

		g.Describe("Get All", func() {
			g.It("should get a single variant by id", func() {
				mockShopify.SetPayload([]byte(SampleFullfillmentsJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-id", "oauthom"}

				fulfillments, err := s.GetFulfillments("808950810", c, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(fulfillments).NotTo(BeNil())
				Expect(len(fulfillments)).To(Equal(1))

				f := fulfillments[0]
				Expect(f).NotTo(BeNil())
				Expect(f.Id).To(Equal(int64(255858046)))
				Expect(f.OrderId).To(Equal(int64(450789469)))
				Expect(f.Status).To(Equal("failure"))
				Expect(f.CreatedAt).To(Equal("2015-09-02T14:48:56-04:00"))
				Expect(f.Service).To(Equal("manual"))
				Expect(f.UpdatedAt).To(Equal("2015-09-02T14:48:56-04:00"))
				Expect(f.TrackingCompany).To(Equal(""))
				Expect(f.TrackingNumber).To(Equal("1Z2345"))
				Expect(len(f.TrackingNumbers)).To(Equal(1))
				Expect(f.TrackingNumbers[0]).To(Equal("1Z2345"))
				Expect(f.TrackingUrl).To(Equal("http://wwwapps.ups.com/etracking/tracking.cgi?InquiryNumber1=1Z2345&TypeOfInquiryNumber=T&AcceptUPSLicenseAgreement=yes&submit=Track"))
				Expect(len(f.TrackingUrls)).To(Equal(1))
				Expect(f.TrackingUrls[0]).To(Equal("http://wwwapps.ups.com/etracking/tracking.cgi?InquiryNumber1=1Z2345&TypeOfInquiryNumber=T&AcceptUPSLicenseAgreement=yes&submit=Track"))
				Expect(f.Receipt).NotTo(BeNil())
				Expect(f.Receipt.TestCase).To(Equal(true))
				Expect(f.Receipt.Authorization).To(Equal("123456"))
				Expect(len(f.LineItems)).To(Equal(1))
			})
		})
	})
}

const (
	SampleFullfillmentsJson = `{
  "fulfillments": [
    {
      "id": 255858046,
      "order_id": 450789469,
      "status": "failure",
      "created_at": "2015-09-02T14:48:56-04:00",
      "service": "manual",
      "updated_at": "2015-09-02T14:48:56-04:00",
      "tracking_company": null,
      "tracking_number": "1Z2345",
      "tracking_numbers": [
        "1Z2345"
      ],
      "tracking_url": "http://wwwapps.ups.com/etracking/tracking.cgi?InquiryNumber1=1Z2345&TypeOfInquiryNumber=T&AcceptUPSLicenseAgreement=yes&submit=Track",
      "tracking_urls": [
        "http://wwwapps.ups.com/etracking/tracking.cgi?InquiryNumber1=1Z2345&TypeOfInquiryNumber=T&AcceptUPSLicenseAgreement=yes&submit=Track"
      ],
      "receipt": {
        "testcase": true,
        "authorization": "123456"
      },
      "line_items": [
        {
          "id": 466157049,
          "variant_id": 39072856,
          "title": "IPod Nano - 8gb",
          "quantity": 1,
          "grams": 200,
          "sku": "IPOD2008GREEN",
          "variant_title": "green",
          "vendor": null,
          "fulfillment_service": "manual",
          "price": "199.00",
          "product_id": 632910392,
          "taxable": true,
          "requires_shipping": true,
          "gift_card": false,
          "name": "IPod Nano - 8gb - green",
          "variant_inventory_management": "shopify",
          "properties": [
            {
              "name": "Custom Engraving Front",
              "value": "Happy Birthday"
            },
            {
              "name": "Custom Engraving Back",
              "value": "Merry Christmas"
            }
          ],
          "product_exists": true,
          "fulfillable_quantity": 1,
          "total_discount": "0.00",
          "fulfillment_status": null,
          "tax_lines": [
          ]
        }
      ]
    }
  ]
}`
)
