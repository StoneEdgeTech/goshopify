package goshopify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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

		g.Describe("serialization", func() {
			g.It("should unmarshal a single fulfillment response without error", func() {
				var fulfillmentResponse SingleFulfillmentResponse
				err := json.Unmarshal([]byte(SingleFulfillmentsJson), &fulfillmentResponse)
				Expect(err).NotTo(HaveOccurred())
				Expect(fulfillmentResponse.Fulfillment.Id).To(Equal(int64(1533523203)))
				Expect(fulfillmentResponse.Fulfillment.LineItems).To(HaveLen(4))
			})
		})

		g.Describe("CreateFulfillment", func() {
			g.It("should send back a single fulfillment after creating a fulfillment", func() {
				mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte(SingleFulfillmentsJson))
				}))
				s := &Shopify{mock.URL}
				c := &Credentials{"some-cart-id", "oauthom"}

				fulfillmentJson := `{"fulfillment": {"tracking_number": "9405510200882805665013" }}`
				orderId := `2017986627`
				fulfillment, err := s.CreateFulfillment(fulfillmentJson, orderId, c, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(fulfillment.Id).To(Equal(int64(1533523203)))
				Expect(fulfillment.LineItems).To(HaveLen(4))
			})
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
	SingleFulfillmentsJson = `{
    "fulfillment": {
        "created_at": "2015-11-20T17:39:18-05:00",
        "id": 1533523203,
        "line_items": [
            {
                "fulfillable_quantity": 0,
                "fulfillment_service": "manual",
                "fulfillment_status": "fulfilled",
                "gift_card": false,
                "grams": 0,
                "id": 3438143235,
                "name": "A Thing",
                "price": "3.00",
                "product_exists": true,
                "product_id": 2973534595,
                "properties": [],
                "quantity": 4,
                "requires_shipping": true,
                "sku": "123458",
                "tax_lines": [],
                "taxable": true,
                "title": "A Thing",
                "total_discount": "0.00",
                "variant_id": 8683025987,
                "variant_inventory_management": "shopify",
                "variant_title": "",
                "vendor": "service tier development store"
            },
            {
                "fulfillable_quantity": 0,
                "fulfillment_service": "manual",
                "fulfillment_status": "fulfilled",
                "gift_card": false,
                "grams": 454,
                "id": 3438143299,
                "name": "Joseph's Amazing Technicolor Dreamcoat - medium / lilac",
                "price": "20.00",
                "product_exists": true,
                "product_id": 1783591875,
                "properties": [],
                "quantity": 2,
                "requires_shipping": true,
                "sku": "346723476062",
                "tax_lines": [],
                "taxable": true,
                "title": "Joseph's Amazing Technicolor Dreamcoat",
                "total_discount": "0.00",
                "variant_id": 5161007363,
                "variant_inventory_management": "shopify",
                "variant_title": "medium / lilac",
                "vendor": "service tier development store"
            },
            {
                "fulfillable_quantity": 0,
                "fulfillment_service": "manual",
                "fulfillment_status": "fulfilled",
                "gift_card": false,
                "grams": 241,
                "id": 3438143363,
                "name": "Incredibly Epic Hoodie - Blue / Large",
                "price": "12.34",
                "product_exists": true,
                "product_id": 2191471363,
                "properties": [],
                "quantity": 2,
                "requires_shipping": false,
                "sku": "hoodie006",
                "tax_lines": [],
                "taxable": false,
                "title": "Incredibly Epic Hoodie",
                "total_discount": "0.00",
                "variant_id": 6284702403,
                "variant_inventory_management": "shopify",
                "variant_title": "Blue / Large",
                "vendor": "service tier development store"
            },
            {
                "fulfillable_quantity": 0,
                "fulfillment_service": "manual",
                "fulfillment_status": "fulfilled",
                "gift_card": false,
                "grams": 3629,
                "id": 3438143427,
                "name": "Example T-Shirt - Lithograph - Height: 9\" x Width: 12\"",
                "price": "25.00",
                "product_exists": true,
                "product_id": 1783688963,
                "properties": [],
                "quantity": 1,
                "requires_shipping": true,
                "sku": "346723476084",
                "tax_lines": [],
                "taxable": true,
                "title": "Example T-Shirt",
                "total_discount": "0.00",
                "variant_id": 5161249027,
                "variant_inventory_management": "shopify",
                "variant_title": "Lithograph - Height: 9\" x Width: 12\"",
                "vendor": "Acme"
            }
        ],
        "order_id": 2017986627,
        "receipt": {},
        "service": "manual",
        "status": "success",
        "tracking_company": "USPS",
        "tracking_number": "9405510200882805665013",
        "tracking_numbers": [
            "9405510200882805665013"
        ],
        "tracking_url": "https://tools.usps.com/go/TrackConfirmAction_input?qtc_tLabels1=9405510200882805665013",
        "tracking_urls": [
            "https://tools.usps.com/go/TrackConfirmAction_input?qtc_tLabels1=9405510200882805665013"
        ],
        "updated_at": "2015-11-20T17:39:18-05:00"
    }
}`
)
