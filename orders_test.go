package goshopify

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestShopifyOrders(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Orders", func() {
		var mockShopify *bogus.Bogus

		g.BeforeEach(func() {
			mockShopify = bogus.New()
			mockShopify.Start()
		})

		g.AfterEach(func() {
			mockShopify.Close()
		})

		g.Describe("Get All", func() {
			g.It("should download a list of orders", func() {
				mockShopify.SetPayload([]byte(SampleOrdersJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-id", "oauthom"}

				orders, err := s.GetOrders(c, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(orders).NotTo(BeNil())
				Expect(len(orders)).To(Equal(1))

				o := orders[0]
				Expect(o.BuyerAcceptsMarketing).To(Equal(false))
				Expect(o.CancelReason).To(Equal(""))
				Expect(o.CancelledAt).To(Equal(""))
				Expect(o.CartToken).To(Equal("68778783ad298f1c80c3bafcddeea02f"))
				Expect(o.CheckoutToken).To(Equal(""))
				Expect(o.ClosedAt).To(Equal(""))
				Expect(o.Confirmed).To(Equal(false))
				Expect(o.CreatedAt).To(Equal("2008-01-10T11:00:00-05:00"))
				Expect(o.Currency).To(Equal("USD"))
				Expect(o.DevicdId).To(Equal(int64(0)))
				Expect(o.Email).To(Equal("bob.norman@hostmail.com"))
				Expect(o.FinancialStatus).To(Equal("authorized"))
				Expect(o.FulfillmentStatus).To(Equal(""))
				Expect(o.Gateway).To(Equal("authorize_net"))
				Expect(o.Id).To(Equal(int64(450789469)))
				Expect(o.LandingSite).To(Equal("http://www.example.com?source=abc"))
				Expect(o.LocationId).To(Equal(int64(0)))
				Expect(o.Name).To(Equal("#1001"))
				Expect(o.Note).To(Equal(""))
				Expect(o.Number).To(Equal(int64(1)))
				Expect(o.ProcessedAt).To(Equal("2008-01-10T11:00:00-05:00"))
				Expect(o.Reference).To(Equal("fhwdgads"))
				Expect(o.ReferringSite).To(Equal("http://www.otherexample.com"))
				Expect(o.SourceIdentifier).To(Equal("fhwdgads"))
				Expect(o.SourceUrl).To(Equal(""))
				Expect(o.SubtotalPrice).To(Equal("398.00"))
				Expect(o.TaxesIncluded).To(Equal(false))
				Expect(o.Test).To(Equal(false))
				Expect(o.Token).To(Equal("b1946ac92492d2347c6235b4d2611184"))
				Expect(o.TotalDiscount).To(Equal("0.00"))
				Expect(o.TotalLineItemsPrice).To(Equal("398.00"))
				Expect(o.TotalPrice).To(Equal("409.94"))
				Expect(o.TotalPriceUsd).To(Equal("409.94"))
				Expect(o.TotalTax).To(Equal("11.94"))
				Expect(o.TotalWeight).To(Equal(int64(0)))
				Expect(o.UpdatedAt).To(Equal("2008-01-10T11:00:00-05:00"))
				Expect(o.UserId).To(Equal(int64(0)))
				Expect(o.BrowserIp).To(Equal(""))
				Expect(o.LandingSiteRef).To(Equal("abc"))
				Expect(o.OrderNumber).To(Equal(int64(1001)))

				Expect(len(o.DiscountCodes)).To(Equal(1))
				d := o.DiscountCodes[0]
				Expect(d).NotTo(BeNil())
				Expect(d.Amount).To(Equal("10.00"))
				Expect(d.Code).To(Equal("TENOFF"))
				Expect(d.Type).To(Equal("percentage"))

				Expect(len(o.NoteAttributes)).To(Equal(2))
				n := o.NoteAttributes[0]
				Expect(n).NotTo(BeNil())
				Expect(n.Name).To(Equal("custom engraving"))
				Expect(n.Value).To(Equal("Happy Birthday"))
				n = o.NoteAttributes[1]
				Expect(n).NotTo(BeNil())
				Expect(n.Name).To(Equal("colour"))
				Expect(n.Value).To(Equal("green"))

				Expect(o.ProcessingMethod).To(Equal("direct"))
				Expect(o.Source).To(Equal("browser"))
				Expect(o.CheckoutId).To(Equal(int64(450789469)))
				Expect(o.SourceName).To(Equal("web"))

				Expect(len(o.TaxLines)).To(Equal(1))
				Expect(o.TaxLines[0]).NotTo(BeNil())
				Expect(o.TaxLines[0].Price).To(Equal("11.94"))
				Expect(o.TaxLines[0].Rate).To(Equal(0.06))
				Expect(o.TaxLines[0].Title).To(Equal("State Tax"))

				Expect(o.Tags).To(Equal(""))

				Expect(len(o.LineItems)).To(Equal(3))
				l := o.LineItems[0]
				Expect(l).NotTo(BeNil())
				Expect(l.FulfillmentService).To(Equal("manual"))
				Expect(l.FulfillmentStatus).To(Equal(""))
				Expect(l.GiftCard).To(Equal(false))
				Expect(l.Grams).To(Equal(float64(200)))
				Expect(l.Id).To(Equal(int64(466157049)))
				Expect(l.Price).To(Equal("199.00"))
				Expect(l.ProductId).To(Equal(int64(632910392)))
				Expect(l.Quantity).To(Equal(int64(1)))
				Expect(l.RequiresShipping).To(Equal(true))
				Expect(l.Sku).To(Equal("IPOD2008GREEN"))
				Expect(l.Taxable).To(Equal(true))
				Expect(l.Title).To(Equal("IPod Nano - 8gb"))
				Expect(l.VariantId).To(Equal(int64(39072856)))
				Expect(l.VariantTitle).To(Equal("green"))
				Expect(l.Vendor).To(Equal(""))
				Expect(l.Name).To(Equal("IPod Nano - 8gb - green"))
				Expect(l.VariantInventoryManagement).To(Equal("shopify"))
				Expect(len(l.Properties)).To(Equal(2))
				Expect(l.Properties[0].Name).To(Equal("Custom Engraving Front"))
				Expect(l.Properties[0].Value).To(Equal("Happy Birthday"))
				Expect(l.Properties[1].Name).To(Equal("Custom Engraving Back"))
				Expect(l.Properties[1].Value).To(Equal("Merry Christmas"))
				Expect(l.ProductExists).To(Equal(true))
				Expect(l.FulfillableQuantity).To(Equal(int64(1)))
				Expect(l.TotalDiscount).To(Equal("0.00"))
				Expect(len(l.TaxLines)).To(Equal(0))

				l = o.LineItems[1]
				Expect(l).NotTo(BeNil())
				Expect(l.FulfillmentService).To(Equal("manual"))
				Expect(l.FulfillmentStatus).To(Equal(""))
				Expect(l.GiftCard).To(Equal(false))
				Expect(l.Grams).To(Equal(float64(200)))
				Expect(l.Id).To(Equal(int64(518995019)))
				Expect(l.Price).To(Equal("199.00"))
				Expect(l.ProductId).To(Equal(int64(632910392)))
				Expect(l.Quantity).To(Equal(int64(1)))
				Expect(l.RequiresShipping).To(Equal(true))
				Expect(l.Sku).To(Equal("IPOD2008RED"))
				Expect(l.Taxable).To(Equal(true))
				Expect(l.Title).To(Equal("IPod Nano - 8gb"))
				Expect(l.VariantId).To(Equal(int64(49148385)))
				Expect(l.VariantTitle).To(Equal("red"))
				Expect(l.Vendor).To(Equal(""))
				Expect(l.Name).To(Equal("IPod Nano - 8gb - red"))
				Expect(l.VariantInventoryManagement).To(Equal("shopify"))
				Expect(len(l.Properties)).To(Equal(0))
				Expect(l.ProductExists).To(Equal(true))
				Expect(l.FulfillableQuantity).To(Equal(int64(1)))
				Expect(l.TotalDiscount).To(Equal("0.00"))
				Expect(len(l.TaxLines)).To(Equal(0))

				l = o.LineItems[2]
				Expect(l).NotTo(BeNil())
				Expect(l.FulfillmentService).To(Equal("manual"))
				Expect(l.FulfillmentStatus).To(Equal(""))
				Expect(l.GiftCard).To(Equal(false))
				Expect(l.Grams).To(Equal(float64(200)))
				Expect(l.Id).To(Equal(int64(703073504)))
				Expect(l.Price).To(Equal("199.00"))
				Expect(l.ProductId).To(Equal(int64(632910392)))
				Expect(l.Quantity).To(Equal(int64(1)))
				Expect(l.RequiresShipping).To(Equal(true))
				Expect(l.Sku).To(Equal("IPOD2008BLACK"))
				Expect(l.Taxable).To(Equal(true))
				Expect(l.Title).To(Equal("IPod Nano - 8gb"))
				Expect(l.VariantId).To(Equal(int64(457924702)))
				Expect(l.VariantTitle).To(Equal("black"))
				Expect(l.Vendor).To(Equal(""))
				Expect(l.Name).To(Equal("IPod Nano - 8gb - black"))
				Expect(l.VariantInventoryManagement).To(Equal("shopify"))
				Expect(len(l.Properties)).To(Equal(0))
				Expect(l.ProductExists).To(Equal(true))
				Expect(l.FulfillableQuantity).To(Equal(int64(1)))
				Expect(l.TotalDiscount).To(Equal("0.00"))
				Expect(len(l.TaxLines)).To(Equal(0))

				Expect(len(o.ShippingLines)).To(Equal(1))
				sl := o.ShippingLines[0]
				Expect(sl).NotTo(BeNil())
				Expect(sl.Code).To(Equal("Free Shipping"))
				Expect(sl.Price).To(Equal("0.00"))
				Expect(sl.Source).To(Equal("shopify"))
				Expect(sl.Title).To(Equal("Free Shipping"))
				Expect(len(sl.TaxLines)).To(Equal(0))

				Expect(o.BillingAddress).NotTo(BeNil())
				Expect(o.BillingAddress.Address1).To(Equal("123 Amoebobacterieae St"))
				Expect(o.BillingAddress.Address2).To(Equal(""))
				Expect(o.BillingAddress.City).To(Equal("Ottawa"))
				Expect(o.BillingAddress.Company).To(Equal(""))
				Expect(o.BillingAddress.Country).To(Equal("Canada"))
				Expect(o.BillingAddress.FirstName).To(Equal("Bob"))
				Expect(o.BillingAddress.LastName).To(Equal("Bobsen"))
				Expect(o.BillingAddress.Latitude).To(Equal(45.41634))
				Expect(o.BillingAddress.Longitude).To(Equal(-75.6868))
				Expect(o.BillingAddress.Phone).To(Equal("(555)555-5555"))
				Expect(o.BillingAddress.Province).To(Equal("Ontario"))
				Expect(o.BillingAddress.Zip).To(Equal("K2P0V6"))
				Expect(o.BillingAddress.Name).To(Equal("Bob Bobsen"))
				Expect(o.BillingAddress.CountryCode).To(Equal("CA"))
				Expect(o.BillingAddress.ProvinceCode).To(Equal("ON"))

				Expect(o.ShippingAddress).NotTo(BeNil())
				Expect(o.ShippingAddress.Address1).To(Equal("123 Amoebobacterieae St"))
				Expect(o.ShippingAddress.Address2).To(Equal(""))
				Expect(o.ShippingAddress.City).To(Equal("Ottawa"))
				Expect(o.ShippingAddress.Company).To(Equal(""))
				Expect(o.ShippingAddress.Country).To(Equal("Canada"))
				Expect(o.ShippingAddress.FirstName).To(Equal("Bob"))
				Expect(o.ShippingAddress.LastName).To(Equal("Bobsen"))
				Expect(o.ShippingAddress.Latitude).To(Equal(45.41634))
				Expect(o.ShippingAddress.Longitude).To(Equal(-75.6868))
				Expect(o.ShippingAddress.Phone).To(Equal("(555)555-5555"))
				Expect(o.ShippingAddress.Province).To(Equal("Ontario"))
				Expect(o.ShippingAddress.Zip).To(Equal("K2P0V6"))
				Expect(o.ShippingAddress.Name).To(Equal("Bob Bobsen"))
				Expect(o.ShippingAddress.CountryCode).To(Equal("CA"))
				Expect(o.ShippingAddress.ProvinceCode).To(Equal("ON"))

				Expect(len(o.Fulfillments)).To(Equal(1))
				f := o.Fulfillments[0]
				Expect(f).NotTo(BeNil())
				Expect(f.CreatedAt).To(Equal("2015-03-28T13:29:45-04:00"))
				Expect(f.Id).To(Equal(int64(255858046)))
				Expect(f.OrderId).To(Equal(int64(450789469)))
				Expect(f.Service).To(Equal("manual"))
				Expect(f.Status).To(Equal("failure"))
				Expect(f.TrackingCompany).To(Equal(""))
				Expect(f.UpdatedAt).To(Equal("2015-03-28T13:29:45-04:00"))
				Expect(f.TrackingNumber).To(Equal("1Z2345"))
				Expect(len(f.TrackingNumbers)).To(Equal(1))
				Expect(f.TrackingUrl).To(Equal("http://wwwapps.ups.com/etracking/tracking.cgi?InquiryNumber1=1Z2345&TypeOfInquiryNumber=T&AcceptUPSLicenseAgreement=yes&submit=Track"))
				Expect(len(f.TrackingUrls)).To(Equal(1))
				Expect(f.Receipt.TestCase).To(Equal(true))
				Expect(f.Receipt.Authorization).To(Equal("123456"))
				Expect(len(f.LineItems)).To(Equal(1))

				Expect(o.ClientDetails).ToNot(BeNil())
				Expect(o.ClientDetails.AcceptLanguage).To(Equal(""))
				Expect(o.ClientDetails.BrowserHeight).To(Equal(int64(0)))
				Expect(o.ClientDetails.BrowserWidth).To(Equal(int64(0)))
				Expect(o.ClientDetails.SessionHash).To(Equal(""))
				Expect(o.ClientDetails.UserAgent).To(Equal(""))
				Expect(o.ClientDetails.BrowserIp).To(Equal("0.0.0.0"))

				Expect(len(o.Refunds)).To(Equal(1))
				r := o.Refunds[0]
				Expect(r).NotTo(BeNil())
				Expect(r.CreatedAt).To(Equal("2015-03-28T13:29:45-04:00"))
				Expect(r.Id).To(Equal(int64(509562969)))
				Expect(r.Note).To(Equal("it broke during shipping"))
				Expect(r.OrderId).To(Equal(int64(450789469)))
				Expect(r.Restock).To(Equal(true))
				Expect(r.UserId).To(Equal(int64(799407056)))
				Expect(len(r.RefundLineItems)).To(Equal(2))
				Expect(len(r.Transactions)).To(Equal(1))

				Expect(o.Customer).NotTo(BeNil())
			})

			g.It("should be able to parse this order", func() {
				mockShopify.SetPayload([]byte(SampleLiveOrdersJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-id", "oauthom"}

				o, err := s.GetOrders(c, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(o).NotTo(BeNil())
				Expect(len(o)).To(Equal(6))
			})

			g.It("should return a specific orders", func() {
				mockShopify.SetPayload([]byte(SampleOrdersSpecificJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-id", "oauthom"}

				o, err := s.GetOrders(c, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(o).NotTo(BeNil())
				Expect(len(o)).To(Equal(1))
			})
		})

		g.Describe("Count", func() {
			g.It("should get an orders count", func() {
				mockShopify.SetPayload([]byte(SampleOrdersCountJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-id", "oauthom"}

				count, err := s.GetOrdersCount(c, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(count).To(Equal(int64(33)))
			})
		})
	})
}

const (
	SampleOrdersCountJson = `{
	"count": 33
}`

	SampleOrdersJson = `{
    "orders": [
        {
            "buyer_accepts_marketing": false,
            "cancel_reason": null,
            "cancelled_at": null,
            "cart_token": "68778783ad298f1c80c3bafcddeea02f",
            "checkout_token": null,
            "closed_at": null,
            "confirmed": false,
            "created_at": "2008-01-10T11:00:00-05:00",
            "currency": "USD",
            "device_id": null,
            "email": "bob.norman@hostmail.com",
            "financial_status": "authorized",
            "fulfillment_status": null,
            "gateway": "authorize_net",
            "id": 450789469,
            "landing_site": "http://www.example.com?source=abc",
            "location_id": null,
            "name": "#1001",
            "note": null,
            "number": 1,
            "processed_at": "2008-01-10T11:00:00-05:00",
            "reference": "fhwdgads",
            "referring_site": "http://www.otherexample.com",
            "source_identifier": "fhwdgads",
            "source_url": null,
            "subtotal_price": "398.00",
            "taxes_included": false,
            "test": false,
            "token": "b1946ac92492d2347c6235b4d2611184",
            "total_discounts": "0.00",
            "total_line_items_price": "398.00",
            "total_price": "409.94",
            "total_price_usd": "409.94",
            "total_tax": "11.94",
            "total_weight": 0,
            "updated_at": "2008-01-10T11:00:00-05:00",
            "user_id": null,
            "browser_ip": null,
            "landing_site_ref": "abc",
            "order_number": 1001,
            "discount_codes": [
                {
                    "code": "TENOFF",
                    "amount": "10.00",
                    "type": "percentage"
                }
            ],
            "note_attributes": [
                {
                    "name": "custom engraving",
                    "value": "Happy Birthday"
                },
                {
                    "name": "colour",
                    "value": "green"
                }
            ],
            "processing_method": "direct",
            "source": "browser",
            "checkout_id": 450789469,
            "source_name": "web",
            "tax_lines": [
                {
                    "price": "11.94",
                    "rate": 0.06,
                    "title": "State Tax"
                }
            ],
            "tags": "",
            "line_items": [
                {
                    "fulfillment_service": "manual",
                    "fulfillment_status": null,
                    "gift_card": false,
                    "grams": 200,
                    "id": 466157049,
                    "price": "199.00",
                    "product_id": 632910392,
                    "quantity": 1,
                    "requires_shipping": true,
                    "sku": "IPOD2008GREEN",
                    "taxable": true,
                    "title": "IPod Nano - 8gb",
                    "variant_id": 39072856,
                    "variant_title": "green",
                    "vendor": null,
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
                    "tax_lines": []
                },
                {
                    "fulfillment_service": "manual",
                    "fulfillment_status": null,
                    "gift_card": false,
                    "grams": 200,
                    "id": 518995019,
                    "price": "199.00",
                    "product_id": 632910392,
                    "quantity": 1,
                    "requires_shipping": true,
                    "sku": "IPOD2008RED",
                    "taxable": true,
                    "title": "IPod Nano - 8gb",
                    "variant_id": 49148385,
                    "variant_title": "red",
                    "vendor": null,
                    "name": "IPod Nano - 8gb - red",
                    "variant_inventory_management": "shopify",
                    "properties": [],
                    "product_exists": true,
                    "fulfillable_quantity": 1,
                    "total_discount": "0.00",
                    "tax_lines": []
                },
                {
                    "fulfillment_service": "manual",
                    "fulfillment_status": null,
                    "gift_card": false,
                    "grams": 200,
                    "id": 703073504,
                    "price": "199.00",
                    "product_id": 632910392,
                    "quantity": 1,
                    "requires_shipping": true,
                    "sku": "IPOD2008BLACK",
                    "taxable": true,
                    "title": "IPod Nano - 8gb",
                    "variant_id": 457924702,
                    "variant_title": "black",
                    "vendor": null,
                    "name": "IPod Nano - 8gb - black",
                    "variant_inventory_management": "shopify",
                    "properties": [],
                    "product_exists": true,
                    "fulfillable_quantity": 1,
                    "total_discount": "0.00",
                    "tax_lines": []
                }
            ],
            "shipping_lines": [
                {
                    "code": "Free Shipping",
                    "price": "0.00",
                    "source": "shopify",
                    "title": "Free Shipping",
                    "tax_lines": []
                }
            ],
            "billing_address": {
                "address1": "123 Amoebobacterieae St",
                "address2": "",
                "city": "Ottawa",
                "company": "",
                "country": "Canada",
                "first_name": "Bob",
                "last_name": "Bobsen",
                "latitude": 45.41634,
                "longitude": -75.6868,
                "phone": "(555)555-5555",
                "province": "Ontario",
                "zip": "K2P0V6",
                "name": "Bob Bobsen",
                "country_code": "CA",
                "province_code": "ON"
            },
            "shipping_address": {
                "address1": "123 Amoebobacterieae St",
                "address2": "",
                "city": "Ottawa",
                "company": "",
                "country": "Canada",
                "first_name": "Bob",
                "last_name": "Bobsen",
                "latitude": 45.41634,
                "longitude": -75.6868,
                "phone": "(555)555-5555",
                "province": "Ontario",
                "zip": "K2P0V6",
                "name": "Bob Bobsen",
                "country_code": "CA",
                "province_code": "ON"
            },
            "fulfillments": [
                {
                    "created_at": "2015-03-28T13:29:45-04:00",
                    "id": 255858046,
                    "order_id": 450789469,
                    "service": "manual",
                    "status": "failure",
                    "tracking_company": null,
                    "updated_at": "2015-03-28T13:29:45-04:00",
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
                            "fulfillment_service": "manual",
                            "fulfillment_status": null,
                            "gift_card": false,
                            "grams": 200,
                            "id": 466157049,
                            "price": "199.00",
                            "product_id": 632910392,
                            "quantity": 1,
                            "requires_shipping": true,
                            "sku": "IPOD2008GREEN",
                            "taxable": true,
                            "title": "IPod Nano - 8gb",
                            "variant_id": 39072856,
                            "variant_title": "green",
                            "vendor": null,
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
                            "tax_lines": []
                        }
                    ]
                }
            ],
            "client_details": {
                "accept_language": null,
                "browser_height": null,
                "browser_ip": "0.0.0.0",
                "browser_width": null,
                "session_hash": null,
                "user_agent": null
            },
            "refunds": [
                {
                    "created_at": "2015-03-28T13:29:45-04:00",
                    "id": 509562969,
                    "note": "it broke during shipping",
                    "order_id": 450789469,
                    "restock": true,
                    "user_id": 799407056,
                    "refund_line_items": [
                        {
                            "id": 104689539,
                            "line_item_id": 703073504,
                            "quantity": 1,
                            "line_item": {
                                "fulfillment_service": "manual",
                                "fulfillment_status": null,
                                "gift_card": false,
                                "grams": 200,
                                "id": 703073504,
                                "price": "199.00",
                                "product_id": 632910392,
                                "quantity": 1,
                                "requires_shipping": true,
                                "sku": "IPOD2008BLACK",
                                "taxable": true,
                                "title": "IPod Nano - 8gb",
                                "variant_id": 457924702,
                                "variant_title": "black",
                                "vendor": null,
                                "name": "IPod Nano - 8gb - black",
                                "variant_inventory_management": "shopify",
                                "properties": [],
                                "product_exists": true,
                                "fulfillable_quantity": 1,
                                "total_discount": "0.00",
                                "tax_lines": []
                            }
                        },
                        {
                            "id": 709875399,
                            "line_item_id": 466157049,
                            "quantity": 1,
                            "line_item": {
                                "fulfillment_service": "manual",
                                "fulfillment_status": null,
                                "gift_card": false,
                                "grams": 200,
                                "id": 466157049,
                                "price": "199.00",
                                "product_id": 632910392,
                                "quantity": 1,
                                "requires_shipping": true,
                                "sku": "IPOD2008GREEN",
                                "taxable": true,
                                "title": "IPod Nano - 8gb",
                                "variant_id": 39072856,
                                "variant_title": "green",
                                "vendor": null,
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
                                "tax_lines": []
                            }
                        }
                    ],
                    "transactions": [
                        {
                            "amount": "209.00",
                            "authorization": "authorization-key",
                            "created_at": "2005-08-05T12:59:12-04:00",
                            "currency": "USD",
                            "gateway": "bogus",
                            "id": 179259969,
                            "kind": "refund",
                            "location_id": null,
                            "message": null,
                            "order_id": 450789469,
                            "parent_id": null,
                            "status": "success",
                            "test": false,
                            "user_id": null,
                            "device_id": null,
                            "receipt": {},
                            "error_code": null,
                            "source_name": "web"
                        }
                    ]
                }
            ],
            "payment_details": {
                "avs_result_code": null,
                "credit_card_bin": null,
                "cvv_result_code": null,
                "credit_card_number": "•••• •••• •••• 4242",
                "credit_card_company": "Visa"
            },
            "customer": {
                "accepts_marketing": false,
                "created_at": "2015-03-28T13:29:45-04:00",
                "email": "bob.norman@hostmail.com",
                "first_name": "Bob",
                "id": 207119551,
                "last_name": "Norman",
                "last_order_id": 450789469,
                "multipass_identifier": null,
                "note": null,
                "orders_count": 1,
                "state": "disabled",
                "total_spent": "41.94",
                "updated_at": "2015-03-28T13:29:45-04:00",
                "verified_email": true,
                "tags": "",
                "last_order_name": "#1001",
                "default_address": {
                    "address1": "Chestnut Street 92",
                    "address2": "",
                    "city": "Louisville",
                    "company": null,
                    "country": "United States",
                    "first_name": null,
                    "id": 207119551,
                    "last_name": null,
                    "phone": "555-625-1199",
                    "province": "Kentucky",
                    "zip": "40202",
                    "name": "",
                    "province_code": "KY",
                    "country_code": "US",
                    "country_name": "United States",
                    "default": true
                }
            }
        }
    ]
}`

	SampleLiveOrdersJson = `{
   "orders":[
      {
         "id":1212953283,
         "email":"dmiles@monsooncommerce.com",
         "closed_at":null,
         "created_at":"2015-09-10T17:41:57-04:00",
         "updated_at":"2015-09-10T17:41:57-04:00",
         "number":6,
         "note":null,
         "token":"cf52eb02bfdc6aba230f230964b3409c",
         "gateway":"usa_epay",
         "test":true,
         "total_price":"31.48",
         "subtotal_price":"20.00",
         "total_weight":454,
         "total_tax":"1.48",
         "taxes_included":false,
         "currency":"USD",
         "financial_status":"authorized",
         "confirmed":true,
         "total_discounts":"0.00",
         "total_line_items_price":"20.00",
         "cart_token":"2d37505e15e4e2d81ae75d0b74480d30",
         "buyer_accepts_marketing":false,
         "name":"#1006",
         "referring_site":"",
         "landing_site":"\/",
         "cancelled_at":null,
         "cancel_reason":null,
         "total_price_usd":"31.48",
         "checkout_token":"527d51b6fa1ff5100efbbb552284c2cf",
         "reference":null,
         "user_id":null,
         "location_id":null,
         "source_identifier":null,
         "source_url":null,
         "processed_at":"2015-09-10T17:41:57-04:00",
         "device_id":null,
         "browser_ip":"98.246.4.93",
         "landing_site_ref":null,
         "order_number":1006,
         "discount_codes":[

         ],
         "note_attributes":[

         ],
         "payment_gateway_names":[
            "usa_epay"
         ],
         "processing_method":"direct",
         "source":"checkout_next",
         "checkout_id":2500066947,
         "source_name":"web",
         "fulfillment_status":null,
         "tax_lines":[
            {
               "title":"CO State Tax",
               "price":"0.58",
               "rate":0.029
            },
            {
               "title":"Larimer County Tax",
               "price":"0.13",
               "rate":0.0065
            },
            {
               "title":"Fort Collins Municipal Tax",
               "price":"0.77",
               "rate":0.0385
            }
         ],
         "tags":"",
         "line_items":[
            {
               "id":2314654403,
               "variant_id":5161004739,
               "title":"Joseph's Amazing Technicolor Dreamcoat",
               "quantity":1,
               "price":"20.00",
               "grams":454,
               "sku":"346723476022",
               "variant_title":"small \/ yellow",
               "vendor":"service tier development store",
               "fulfillment_service":"manual",
               "product_id":1783591875,
               "requires_shipping":true,
               "taxable":true,
               "gift_card":false,
               "name":"Joseph's Amazing Technicolor Dreamcoat - small \/ yellow",
               "variant_inventory_management":"shopify",
               "properties":[

               ],
               "product_exists":true,
               "fulfillable_quantity":1,
               "total_discount":"0.00",
               "fulfillment_status":null,
               "tax_lines":[
                  {
                     "title":"CO State Tax",
                     "price":"0.58",
                     "rate":0.029
                  },
                  {
                     "title":"Larimer County Tax",
                     "price":"0.13",
                     "rate":0.0065
                  },
                  {
                     "title":"Fort Collins Municipal Tax",
                     "price":"0.77",
                     "rate":0.0385
                  }
               ]
            }
         ],
         "shipping_lines":[
            {
               "title":"Standard Shipping",
               "price":"10.00",
               "code":"Standard Shipping",
               "source":"shopify",
               "tax_lines":[

               ]
            }
         ],
         "billing_address":{
            "first_name":"Daniel",
            "address1":"606 Homestead Ct",
            "phone":"",
            "city":"Fort Collins",
            "zip":"80526",
            "province":"Colorado",
            "country":"United States",
            "last_name":"MilesFortCollins",
            "address2":"",
            "company":"",
            "latitude":40.543306,
            "longitude":-105.087134,
            "name":"Daniel MilesFortCollins",
            "country_code":"US",
            "province_code":"CO"
         },
         "shipping_address":{
            "first_name":"Daniel",
            "address1":"606 Homestead Ct",
            "phone":"",
            "city":"Fort Collins",
            "zip":"80526",
            "province":"Colorado",
            "country":"United States",
            "last_name":"MilesFortCollins",
            "address2":"",
            "company":"",
            "latitude":40.543306,
            "longitude":-105.087134,
            "name":"Daniel MilesFortCollins",
            "country_code":"US",
            "province_code":"CO"
         },
         "fulfillments":[

         ],
         "client_details":{
            "browser_ip":"98.246.4.93",
            "accept_language":"en-US,en;q=0.5",
            "user_agent":"Mozilla\/5.0 (Macintosh; Intel Mac OS X 10.10; rv:40.0) Gecko\/20100101 Firefox\/40.0",
            "session_hash":"4f9011df85e3fc17aa84970a491a8e1e0e84caf3f7eb89a8b5cd794005142e1c",
            "browser_width":1440,
            "browser_height":761
         },
         "refunds":[

         ],
         "payment_details":{
            "credit_card_bin":"411111",
            "avs_result_code":"YYY",
            "cvv_result_code":"M",
            "credit_card_number":"•••• •••• •••• 1111",
            "credit_card_company":"Visa"
         },
         "customer":{
            "id":1311218883,
            "email":"dmiles@monsooncommerce.com",
            "accepts_marketing":false,
            "created_at":"2015-09-10T17:39:41-04:00",
            "updated_at":"2015-09-10T17:41:57-04:00",
            "first_name":"Daniel",
            "last_name":"MilesFortCollins",
            "orders_count":0,
            "state":"disabled",
            "total_spent":"0.00",
            "last_order_id":null,
            "note":null,
            "verified_email":true,
            "multipass_identifier":null,
            "tax_exempt":false,
            "tags":"",
            "last_order_name":null,
            "default_address":{
               "id":1423874243,
               "first_name":"Daniel",
               "last_name":"MilesFortCollins",
               "company":"",
               "address1":"606 Homestead Ct",
               "address2":"",
               "city":"Fort Collins",
               "province":"Colorado",
               "country":"United States",
               "zip":"80526",
               "phone":"",
               "name":"Daniel MilesFortCollins",
               "province_code":"CO",
               "country_code":"US",
               "country_name":"United States",
               "default":true
            }
         }
      },
      {
         "id":1212944963,
         "email":"dmiles@monsooncommerce.com",
         "closed_at":null,
         "created_at":"2015-09-10T17:40:20-04:00",
         "updated_at":"2015-09-10T17:40:20-04:00",
         "number":5,
         "note":null,
         "token":"70f011606be84a93656ac2e237b540c1",
         "gateway":"usa_epay",
         "test":true,
         "total_price":"30.58",
         "subtotal_price":"20.00",
         "total_weight":454,
         "total_tax":"0.58",
         "taxes_included":false,
         "currency":"USD",
         "financial_status":"authorized",
         "confirmed":true,
         "total_discounts":"0.00",
         "total_line_items_price":"20.00",
         "cart_token":"cc095172a7d98610fb6d3ba8870d9742",
         "buyer_accepts_marketing":false,
         "name":"#1005",
         "referring_site":"",
         "landing_site":"\/",
         "cancelled_at":null,
         "cancel_reason":null,
         "total_price_usd":"30.58",
         "checkout_token":"0c0780c3cbd30cd659fd32286e7279ff",
         "reference":null,
         "user_id":null,
         "location_id":null,
         "source_identifier":null,
         "source_url":null,
         "processed_at":"2015-09-10T17:40:20-04:00",
         "device_id":null,
         "browser_ip":"98.246.4.93",
         "landing_site_ref":null,
         "order_number":1005,
         "discount_codes":[

         ],
         "note_attributes":[

         ],
         "payment_gateway_names":[
            "usa_epay"
         ],
         "processing_method":"direct",
         "source":"checkout_next",
         "checkout_id":2500038467,
         "source_name":"web",
         "fulfillment_status":null,
         "tax_lines":[
            {
               "title":"CO State Tax",
               "price":"0.58",
               "rate":0.029
            }
         ],
         "tags":"",
         "line_items":[
            {
               "id":2314637379,
               "variant_id":5161005379,
               "title":"Joseph's Amazing Technicolor Dreamcoat",
               "quantity":1,
               "price":"20.00",
               "grams":454,
               "sku":"346723476033",
               "variant_title":"small \/ lilac",
               "vendor":"service tier development store",
               "fulfillment_service":"manual",
               "product_id":1783591875,
               "requires_shipping":true,
               "taxable":true,
               "gift_card":false,
               "name":"Joseph's Amazing Technicolor Dreamcoat - small \/ lilac",
               "variant_inventory_management":"shopify",
               "properties":[

               ],
               "product_exists":true,
               "fulfillable_quantity":1,
               "total_discount":"0.00",
               "fulfillment_status":null,
               "tax_lines":[
                  {
                     "title":"CO State Tax",
                     "price":"0.58",
                     "rate":0.029
                  }
               ]
            }
         ],
         "shipping_lines":[
            {
               "title":"Standard Shipping",
               "price":"10.00",
               "code":"Standard Shipping",
               "source":"shopify",
               "tax_lines":[

               ]
            }
         ],
         "billing_address":{
            "first_name":"Daniel",
            "address1":"12613 Zuni St",
            "phone":"",
            "city":"Broomfield",
            "zip":"80020",
            "province":"Colorado",
            "country":"United States",
            "last_name":"MilesBroomfield",
            "address2":"207",
            "company":"",
            "latitude":39.9263668,
            "longitude":-105.0157143,
            "name":"Daniel MilesBroomfield",
            "country_code":"US",
            "province_code":"CO"
         },
         "shipping_address":{
            "first_name":"Daniel",
            "address1":"12613 Zuni St",
            "phone":"",
            "city":"Broomfield",
            "zip":"80020",
            "province":"Colorado",
            "country":"United States",
            "last_name":"MilesBroomfield",
            "address2":"207",
            "company":"",
            "latitude":39.9263668,
            "longitude":-105.0157143,
            "name":"Daniel MilesBroomfield",
            "country_code":"US",
            "province_code":"CO"
         },
         "fulfillments":[

         ],
         "client_details":{
            "browser_ip":"98.246.4.93",
            "accept_language":"en-US,en;q=0.5",
            "user_agent":"Mozilla\/5.0 (Macintosh; Intel Mac OS X 10.10; rv:40.0) Gecko\/20100101 Firefox\/40.0",
            "session_hash":"4f9011df85e3fc17aa84970a491a8e1e0e84caf3f7eb89a8b5cd794005142e1c",
            "browser_width":1440,
            "browser_height":765
         },
         "refunds":[

         ],
         "payment_details":{
            "credit_card_bin":"411111",
            "avs_result_code":"YYY",
            "cvv_result_code":"M",
            "credit_card_number":"•••• •••• •••• 1111",
            "credit_card_company":"Visa"
         },
         "customer":{
            "id":1311218883,
            "email":"dmiles@monsooncommerce.com",
            "accepts_marketing":false,
            "created_at":"2015-09-10T17:39:41-04:00",
            "updated_at":"2015-09-10T17:41:57-04:00",
            "first_name":"Daniel",
            "last_name":"MilesFortCollins",
            "orders_count":0,
            "state":"disabled",
            "total_spent":"0.00",
            "last_order_id":null,
            "note":null,
            "verified_email":true,
            "multipass_identifier":null,
            "tax_exempt":false,
            "tags":"",
            "last_order_name":null,
            "default_address":{
               "id":1423874243,
               "first_name":"Daniel",
               "last_name":"MilesFortCollins",
               "company":"",
               "address1":"606 Homestead Ct",
               "address2":"",
               "city":"Fort Collins",
               "province":"Colorado",
               "country":"United States",
               "zip":"80526",
               "phone":"",
               "name":"Daniel MilesFortCollins",
               "province_code":"CO",
               "country_code":"US",
               "country_name":"United States",
               "default":true
            }
         }
      },
      {
         "id":1106552451,
         "email":"",
         "closed_at":null,
         "created_at":"2015-08-25T19:44:16-04:00",
         "updated_at":"2015-08-25T19:44:16-04:00",
         "number":4,
         "note":"",
         "token":"f7e8aaf42d0d585a6e0a1ce045480a84",
         "gateway":"usa_epay",
         "test":true,
         "total_price":"25.00",
         "subtotal_price":"25.00",
         "total_weight":3629,
         "total_tax":"0.00",
         "taxes_included":false,
         "currency":"USD",
         "financial_status":"authorized",
         "confirmed":true,
         "total_discounts":"0.00",
         "total_line_items_price":"25.00",
         "cart_token":null,
         "buyer_accepts_marketing":false,
         "name":"#1004",
         "referring_site":null,
         "landing_site":null,
         "cancelled_at":null,
         "cancel_reason":null,
         "total_price_usd":"25.00",
         "checkout_token":"d1da6fcf13ad4d90dc6b28a5091dc054",
         "reference":null,
         "user_id":31175171,
         "location_id":null,
         "source_identifier":null,
         "source_url":null,
         "processed_at":"2015-08-25T19:44:16-04:00",
         "device_id":null,
         "browser_ip":"173.12.185.113",
         "landing_site_ref":null,
         "order_number":1004,
         "discount_codes":[

         ],
         "note_attributes":[

         ],
         "payment_gateway_names":[
            "usa_epay"
         ],
         "processing_method":"direct",
         "source":"shopify_draft_order",
         "checkout_id":2260107395,
         "source_name":"shopify_draft_order",
         "fulfillment_status":null,
         "tax_lines":[

         ],
         "tags":"",
         "line_items":[
            {
               "id":2110785219,
               "variant_id":5161249027,
               "title":"Example T-Shirt",
               "quantity":1,
               "price":"25.00",
               "grams":3629,
               "sku":"346723476084",
               "variant_title":"Lithograph - Height: 9\" x Width: 12\"",
               "vendor":"Acme",
               "fulfillment_service":"manual",
               "product_id":1783688963,
               "requires_shipping":true,
               "taxable":true,
               "gift_card":false,
               "name":"Example T-Shirt - Lithograph - Height: 9\" x Width: 12\"",
               "variant_inventory_management":null,
               "properties":[

               ],
               "product_exists":true,
               "fulfillable_quantity":1,
               "total_discount":"0.00",
               "fulfillment_status":null,
               "tax_lines":[

               ]
            }
         ],
         "shipping_lines":[

         ],
         "billing_address":{
            "first_name":"Service",
            "address1":"520 NW Davis St",
            "phone":"503-555-5555",
            "city":"Portland",
            "zip":"97210",
            "province":"Oregon",
            "country":"United States",
            "last_name":"Tier",
            "address2":"",
            "company":null,
            "latitude":45.5243176,
            "longitude":-122.6758525,
            "name":"Service Tier",
            "country_code":"US",
            "province_code":"OR"
         },
         "fulfillments":[

         ],
         "client_details":{
            "browser_ip":"173.12.185.113",
            "accept_language":"en-US,en;q=0.8",
            "user_agent":"Mozilla\/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit\/537.36 (KHTML, like Gecko) Chrome\/44.0.2403.157 Safari\/537.36",
            "session_hash":"ee89e8731822549e734a6cd41ded988ff2076d2fc8b1ea7137ddcccc770ac29a",
            "browser_width":null,
            "browser_height":null
         },
         "refunds":[

         ],
         "payment_details":{
            "credit_card_bin":"411111",
            "avs_result_code":"YYY",
            "cvv_result_code":"M",
            "credit_card_number":"•••• •••• •••• 1111",
            "credit_card_company":"Visa"
         }
      },
      {
         "id":1106534083,
         "email":"",
         "closed_at":null,
         "created_at":"2015-08-25T19:40:48-04:00",
         "updated_at":"2015-09-03T15:48:49-04:00",
         "number":3,
         "note":"",
         "token":"7f4213d6d195f92c95d7219281bc2f44",
         "gateway":"manual",
         "test":false,
         "total_price":"1.00",
         "subtotal_price":"1.00",
         "total_weight":227,
         "total_tax":"0.00",
         "taxes_included":false,
         "currency":"USD",
         "financial_status":"paid",
         "confirmed":true,
         "total_discounts":"0.00",
         "total_line_items_price":"1.00",
         "cart_token":null,
         "buyer_accepts_marketing":false,
         "name":"#1003",
         "referring_site":null,
         "landing_site":null,
         "cancelled_at":null,
         "cancel_reason":null,
         "total_price_usd":"1.00",
         "checkout_token":null,
         "reference":null,
         "user_id":31175171,
         "location_id":null,
         "source_identifier":null,
         "source_url":null,
         "processed_at":"2015-08-25T19:40:48-04:00",
         "device_id":null,
         "browser_ip":null,
         "landing_site_ref":null,
         "order_number":1003,
         "discount_codes":[

         ],
         "note_attributes":[

         ],
         "payment_gateway_names":[
            "manual"
         ],
         "processing_method":"manual",
         "source":"shopify_draft_order",
         "checkout_id":null,
         "source_name":"shopify_draft_order",
         "fulfillment_status":"fulfilled",
         "tax_lines":[

         ],
         "tags":"",
         "line_items":[
            {
               "id":2110749123,
               "variant_id":6172675075,
               "title":"Ball Point Pen",
               "quantity":1,
               "price":"1.00",
               "grams":227,
               "sku":"toplevelballpointpensku-2",
               "variant_title":"blue",
               "vendor":"service tier development store",
               "fulfillment_service":"manual",
               "product_id":2152810179,
               "requires_shipping":true,
               "taxable":true,
               "gift_card":false,
               "name":"Ball Point Pen - blue",
               "variant_inventory_management":"shopify",
               "properties":[

               ],
               "product_exists":true,
               "fulfillable_quantity":0,
               "total_discount":"0.00",
               "fulfillment_status":"fulfilled",
               "tax_lines":[

               ]
            }
         ],
         "shipping_lines":[

         ],
         "fulfillments":[
            {
               "id":996678147,
               "order_id":1106534083,
               "status":"success",
               "created_at":"2015-09-03T15:48:49-04:00",
               "service":"manual",
               "updated_at":"2015-09-03T15:48:49-04:00",
               "tracking_company":"UPS",
               "tracking_number":null,
               "tracking_numbers":[

               ],
               "tracking_url":null,
               "tracking_urls":[

               ],
               "receipt":{

               },
               "line_items":[
                  {
                     "id":2110749123,
                     "variant_id":6172675075,
                     "title":"Ball Point Pen",
                     "quantity":1,
                     "price":"1.00",
                     "grams":227,
                     "sku":"toplevelballpointpensku-2",
                     "variant_title":"blue",
                     "vendor":"service tier development store",
                     "fulfillment_service":"manual",
                     "product_id":2152810179,
                     "requires_shipping":true,
                     "taxable":true,
                     "gift_card":false,
                     "name":"Ball Point Pen - blue",
                     "variant_inventory_management":"shopify",
                     "properties":[

                     ],
                     "product_exists":true,
                     "fulfillable_quantity":0,
                     "total_discount":"0.00",
                     "fulfillment_status":"fulfilled",
                     "tax_lines":[

                     ]
                  }
               ]
            }
         ],
         "refunds":[

         ]
      },
      {
         "id":1028364611,
         "email":"",
         "closed_at":null,
         "created_at":"2015-08-14T16:58:54-04:00",
         "updated_at":"2015-08-14T16:58:54-04:00",
         "number":2,
         "note":"",
         "token":"3ef489cbd674195fd3039681ff1df356",
         "gateway":"usa_epay",
         "test":true,
         "total_price":"20.00",
         "subtotal_price":"20.00",
         "total_weight":454,
         "total_tax":"0.00",
         "taxes_included":false,
         "currency":"USD",
         "financial_status":"authorized",
         "confirmed":true,
         "total_discounts":"0.00",
         "total_line_items_price":"20.00",
         "cart_token":null,
         "buyer_accepts_marketing":false,
         "name":"#1002",
         "referring_site":null,
         "landing_site":null,
         "cancelled_at":null,
         "cancel_reason":null,
         "total_price_usd":"20.00",
         "checkout_token":"f97647dc0ccbaca5ad8af6baff6e1add",
         "reference":null,
         "user_id":31175171,
         "location_id":null,
         "source_identifier":null,
         "source_url":null,
         "processed_at":"2015-08-14T16:58:54-04:00",
         "device_id":null,
         "browser_ip":"173.12.185.113",
         "landing_site_ref":null,
         "order_number":1002,
         "discount_codes":[

         ],
         "note_attributes":[

         ],
         "payment_gateway_names":[
            "usa_epay"
         ],
         "processing_method":"direct",
         "source":"shopify_draft_order",
         "checkout_id":2088466307,
         "source_name":"shopify_draft_order",
         "fulfillment_status":null,
         "tax_lines":[

         ],
         "tags":"",
         "line_items":[
            {
               "id":1935292035,
               "variant_id":5161009603,
               "title":"Joseph's Amazing Technicolor Dreamcoat",
               "quantity":1,
               "price":"20.00",
               "grams":454,
               "sku":"346723476093",
               "variant_title":"large \/ chocolate",
               "vendor":"service tier development store",
               "fulfillment_service":"manual",
               "product_id":1783591875,
               "requires_shipping":true,
               "taxable":true,
               "gift_card":false,
               "name":"Joseph's Amazing Technicolor Dreamcoat - large \/ chocolate",
               "variant_inventory_management":"shopify",
               "properties":[

               ],
               "product_exists":true,
               "fulfillable_quantity":1,
               "total_discount":"0.00",
               "fulfillment_status":null,
               "tax_lines":[

               ]
            }
         ],
         "shipping_lines":[

         ],
         "fulfillments":[

         ],
         "client_details":{
            "browser_ip":"173.12.185.113",
            "accept_language":"en-US,en;q=0.5",
            "user_agent":"Mozilla\/5.0 (Macintosh; Intel Mac OS X 10.10; rv:39.0) Gecko\/20100101 Firefox\/39.0",
            "session_hash":"c2dad2af6be6667d579507ea4a3476340bdb8e241a95f9ba2a3eb8c75afc9e5f",
            "browser_width":null,
            "browser_height":null
         },
         "refunds":[

         ],
         "payment_details":{
            "credit_card_bin":"411111",
            "avs_result_code":"YYY",
            "cvv_result_code":"N",
            "credit_card_number":"•••• •••• •••• 1111",
            "credit_card_company":"Visa"
         }
      },
      {
         "id":1028133379,
         "email":"",
         "closed_at":null,
         "created_at":"2015-08-14T16:19:42-04:00",
         "updated_at":"2015-08-14T16:19:43-04:00",
         "number":1,
         "note":"",
         "token":"007bb9ca97e01ebb3fcaa18784680643",
         "gateway":"authorize_net",
         "test":true,
         "total_price":"20.00",
         "subtotal_price":"20.00",
         "total_weight":454,
         "total_tax":"0.00",
         "taxes_included":false,
         "currency":"USD",
         "financial_status":"authorized",
         "confirmed":true,
         "total_discounts":"0.00",
         "total_line_items_price":"20.00",
         "cart_token":null,
         "buyer_accepts_marketing":false,
         "name":"#1001",
         "referring_site":null,
         "landing_site":null,
         "cancelled_at":null,
         "cancel_reason":null,
         "total_price_usd":"20.00",
         "checkout_token":"ed86df1d52fb36cc0987158ef9be340b",
         "reference":null,
         "user_id":31175171,
         "location_id":null,
         "source_identifier":null,
         "source_url":null,
         "processed_at":"2015-08-14T16:19:42-04:00",
         "device_id":null,
         "browser_ip":"173.12.185.113",
         "landing_site_ref":null,
         "order_number":1001,
         "discount_codes":[

         ],
         "note_attributes":[

         ],
         "payment_gateway_names":[
            "authorize_net"
         ],
         "processing_method":"direct",
         "source":"shopify_draft_order",
         "checkout_id":2088054915,
         "source_name":"shopify_draft_order",
         "fulfillment_status":null,
         "tax_lines":[

         ],
         "tags":"",
         "line_items":[
            {
               "id":1934849795,
               "variant_id":5161008899,
               "title":"Joseph's Amazing Technicolor Dreamcoat",
               "quantity":1,
               "price":"20.00",
               "grams":454,
               "sku":"346723476084",
               "variant_title":"large \/ black",
               "vendor":"service tier development store",
               "fulfillment_service":"manual",
               "product_id":1783591875,
               "requires_shipping":true,
               "taxable":true,
               "gift_card":false,
               "name":"Joseph's Amazing Technicolor Dreamcoat - large \/ black",
               "variant_inventory_management":"shopify",
               "properties":[

               ],
               "product_exists":true,
               "fulfillable_quantity":1,
               "total_discount":"0.00",
               "fulfillment_status":null,
               "tax_lines":[

               ]
            }
         ],
         "shipping_lines":[

         ],
         "fulfillments":[

         ],
         "client_details":{
            "browser_ip":"173.12.185.113",
            "accept_language":"en-US,en;q=0.5",
            "user_agent":"Mozilla\/5.0 (Macintosh; Intel Mac OS X 10.10; rv:39.0) Gecko\/20100101 Firefox\/39.0",
            "session_hash":"c2dad2af6be6667d579507ea4a3476340bdb8e241a95f9ba2a3eb8c75afc9e5f",
            "browser_width":null,
            "browser_height":null
         },
         "refunds":[

         ],
         "payment_details":{
            "credit_card_bin":"411111",
            "avs_result_code":"Y",
            "cvv_result_code":"P",
            "credit_card_number":"•••• •••• •••• 1111",
            "credit_card_company":"Visa"
         }
      }
   ]
}`

	SampleOrdersSpecificJson = `{
  "orders": [
    {
      "buyer_accepts_marketing": false,
      "cancel_reason": null,
      "cancelled_at": null,
      "cart_token": null,
      "checkout_token": null,
      "closed_at": null,
      "confirmed": true,
      "created_at": "2015-03-28T13:30:39-04:00",
      "currency": "USD",
      "device_id": null,
      "email": "",
      "financial_status": "paid",
      "fulfillment_status": null,
      "gateway": "",
      "id": 1073459968,
      "landing_site": null,
      "location_id": null,
      "name": "#1002",
      "note": null,
      "number": 2,
      "processed_at": "2015-03-28T13:30:39-04:00",
      "reference": null,
      "referring_site": null,
      "source_identifier": null,
      "source_url": null,
      "subtotal_price": "199.00",
      "taxes_included": false,
      "test": false,
      "token": "f0b73ed55096180c2af69ac6fe12f4f8",
      "total_discounts": "0.00",
      "total_line_items_price": "199.00",
      "total_price": "199.00",
      "total_price_usd": "199.00",
      "total_tax": "0.00",
      "total_weight": 0,
      "updated_at": "2015-03-28T13:30:39-04:00",
      "user_id": null,
      "browser_ip": null,
      "landing_site_ref": null,
      "order_number": 1002,
      "discount_codes": [
      ],
      "note_attributes": [
      ],
      "processing_method": "",
      "source": "755357713",
      "checkout_id": null,
      "source_name": "755357713",
      "tax_lines": [
      ],
      "tags": "",
      "line_items": [
        {
          "fulfillment_service": "manual",
          "fulfillment_status": null,
          "gift_card": false,
          "grams": 200,
          "id": 1071823178,
          "price": "199.00",
          "product_id": 921728736,
          "quantity": 1,
          "requires_shipping": true,
          "sku": "IPOD2009BLACK",
          "taxable": true,
          "title": "IPod Touch 8GB",
          "variant_id": 447654529,
          "variant_title": "Black",
          "vendor": null,
          "name": "IPod Touch 8GB - Black",
          "variant_inventory_management": "shopify",
          "properties": [
          ],
          "product_exists": true,
          "fulfillable_quantity": 1,
          "total_discount": "0.00",
          "tax_lines": [
          ]
        }
      ],
      "shipping_lines": [
      ],
      "fulfillments": [
      ],
      "refunds": [
      ]
    }
  ]
}`
)
