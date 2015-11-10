package goshopify

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestShopifyProducts(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Products", func() {
		var mockShopify *bogus.Bogus

		g.BeforeEach(func() {
			mockShopify = bogus.New()
			mockShopify.Start()
		})

		g.AfterEach(func() {
			mockShopify.Close()
		})

		g.Describe("Get All", func() {
			g.It("should get a list of products from shopify", func() {
				mockShopify.SetPayload([]byte(SampleProductsJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-hash", "oauthom"}

				p, err := s.GetProducts(c, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(p).NotTo(BeNil())
				Expect(len(p)).To(Equal(2))

				Expect(p[0].BodyHtml).To(ContainSubstring("It's the small iPod with one very big idea"))
				Expect(p[0].CreatedAt).To(Equal("2015-03-28T13:31:19-04:01"))
				Expect(p[0].Handle).To(Equal("ipod-nano"))
				Expect(p[0].Id).To(Equal(int64(632910392)))
				Expect(p[0].ProductType).To(Equal("Cult Products"))
				Expect(p[0].PublishedAt).To(Equal("2007-12-31T19:00:00-05:00"))
				Expect(p[0].PublishedScope).To(Equal("web"))
				Expect(p[0].TemplateSuffix).To(Equal(""))
				Expect(p[0].Title).To(Equal("IPod Nano - 8GB"))
				Expect(p[0].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:02"))
				Expect(p[0].Vendor).To(Equal("Apple"))
				Expect(p[0].Tags).To(Equal("Emotive, Flash Memory, MP3, Music"))

				Expect(len(p[0].Variants)).To(Equal(4))

				Expect(p[0].Variants[0].Barcode).To(Equal("1234_pink"))
				Expect(p[0].Variants[0].CompareAtPrice).To(Equal(""))
				Expect(p[0].Variants[0].CreatedAt).To(Equal("2015-03-28T13:31:19-04:03"))
				Expect(p[0].Variants[0].FulfillmentService).To(Equal("manual"))
				Expect(p[0].Variants[0].Grams).To(Equal(float64(200)))
				Expect(p[0].Variants[0].Id).To(Equal(int64(808950810)))
				Expect(p[0].Variants[0].InventoryManagement).To(Equal("shopify"))
				Expect(p[0].Variants[0].InventoryPolicy).To(Equal("continue"))
				Expect(p[0].Variants[0].Option1).To(Equal("Pink"))
				Expect(p[0].Variants[0].Option2).To(Equal(""))
				Expect(p[0].Variants[0].Option3).To(Equal(""))
				Expect(p[0].Variants[0].Position).To(Equal(1))
				Expect(p[0].Variants[0].Price).To(Equal("199.00"))
				Expect(p[0].Variants[0].ProductId).To(Equal(int64(632910392)))
				Expect(p[0].Variants[0].RequiresShipping).To(BeTrue())
				Expect(p[0].Variants[0].Sku).To(Equal("IPOD2008PINK"))
				Expect(p[0].Variants[0].Taxable).To(BeTrue())
				Expect(p[0].Variants[0].Title).To(Equal("Pink"))
				Expect(p[0].Variants[0].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:04"))
				Expect(p[0].Variants[0].InventoryQuantity).To(Equal(int64(10)))
				Expect(p[0].Variants[0].OldInventoryQuantity).To(Equal(int64(10)))
				Expect(p[0].Variants[0].ImageId).To(Equal(int64(562641783)))
				Expect(p[0].Variants[0].Weight).To(Equal(0.2))
				Expect(p[0].Variants[0].WeightUnit).To(Equal("kg"))

				Expect(p[0].Variants[1].Barcode).To(Equal("1234_red"))
				Expect(p[0].Variants[1].CompareAtPrice).To(Equal(""))
				Expect(p[0].Variants[1].CreatedAt).To(Equal("2015-03-28T13:31:19-04:05"))
				Expect(p[0].Variants[1].FulfillmentService).To(Equal("manual"))
				Expect(p[0].Variants[1].Grams).To(Equal(float64(200)))
				Expect(p[0].Variants[1].Id).To(Equal(int64(49148385)))
				Expect(p[0].Variants[1].InventoryManagement).To(Equal("shopify"))
				Expect(p[0].Variants[1].InventoryPolicy).To(Equal("continue"))
				Expect(p[0].Variants[1].Option1).To(Equal("Red"))
				Expect(p[0].Variants[1].Option2).To(Equal(""))
				Expect(p[0].Variants[1].Option3).To(Equal(""))
				Expect(p[0].Variants[1].Position).To(Equal(2))
				Expect(p[0].Variants[1].Price).To(Equal("199.00"))
				Expect(p[0].Variants[1].ProductId).To(Equal(int64(632910392)))
				Expect(p[0].Variants[1].RequiresShipping).To(BeTrue())
				Expect(p[0].Variants[1].Sku).To(Equal("IPOD2008RED"))
				Expect(p[0].Variants[1].Taxable).To(BeTrue())
				Expect(p[0].Variants[1].Title).To(Equal("Red"))
				Expect(p[0].Variants[1].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:06"))
				Expect(p[0].Variants[1].InventoryQuantity).To(Equal(int64(20)))
				Expect(p[0].Variants[1].OldInventoryQuantity).To(Equal(int64(20)))
				Expect(p[0].Variants[1].ImageId).To(Equal(int64(0)))
				Expect(p[0].Variants[1].Weight).To(Equal(0.2))
				Expect(p[0].Variants[1].WeightUnit).To(Equal("kg"))

				Expect(len(p[0].Options)).To(Equal(1))

				Expect(p[0].Options[0].Id).To(Equal(int64(594680422)))
				Expect(p[0].Options[0].Name).To(Equal("Title"))
				Expect(p[0].Options[0].Position).To(Equal(1))
				Expect(p[0].Options[0].ProductId).To(Equal(int64(632910392)))

				Expect(len(p[0].Images)).To(Equal(2))

				Expect(p[0].Images[0].CreatedAt).To(Equal("2015-03-28T13:31:19-04:07"))
				Expect(p[0].Images[0].Id).To(Equal(int64(850703190)))
				Expect(p[0].Images[0].Position).To(Equal(1))
				Expect(p[0].Images[0].ProductId).To(Equal(int64(632910392)))
				Expect(p[0].Images[0].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:08"))
				Expect(p[0].Images[0].Source).To(Equal("https://cdn.shopify.com/s/files/1/0006/9093/3842/products/ipod-nano.png?v=1427563879"))
				Expect(len(p[0].Images[0].VariantIds)).To(Equal(0))

				Expect(p[0].Images[1].CreatedAt).To(Equal("2015-03-28T13:31:19-04:09"))
				Expect(p[0].Images[1].Id).To(Equal(int64(562641783)))
				Expect(p[0].Images[1].Position).To(Equal(2))
				Expect(p[0].Images[1].ProductId).To(Equal(int64(632910392)))
				Expect(p[0].Images[1].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:10"))
				Expect(p[0].Images[1].Source).To(Equal("https://cdn.shopify.com/s/files/1/0006/9093/3842/products/ipod-nano-2.png?v=1427563879"))
				Expect(p[0].Images[1].VariantIds[0]).To(Equal(int64(808950810)))

				Expect(p[0].Image.CreatedAt).To(Equal("2015-03-28T13:31:19-04:11"))
				Expect(p[0].Image.Id).To(Equal(int64(850703190)))
				Expect(p[0].Image.Position).To(Equal(1))
				Expect(p[0].Image.ProductId).To(Equal(int64(632910392)))
				Expect(p[0].Image.UpdatedAt).To(Equal("2015-03-28T13:31:19-04:12"))
				Expect(p[0].Image.Source).To(Equal("https://cdn.shopify.com/s/files/1/0006/9093/3842/products/ipod-nano.png?v=1427563879"))
				Expect(len(p[0].Image.VariantIds)).To(Equal(0))

				Expect(len(p[1].Variants)).To(Equal(1))
				Expect(len(p[1].Options)).To(Equal(1))
				Expect(len(p[1].Images)).To(Equal(0))

				Expect(mockShopify.Hits()).To(Equal(1))
			})

			g.It("should return an offset list of products", func() {
				mockShopify.SetPayload([]byte(SampleProductsJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-hash", "oauthom"}

				params := url.Values{}
				params.Add("page", "2")

				p, err := s.GetProducts(c, params)
				Expect(err).NotTo(HaveOccurred())
				Expect(p).NotTo(BeNil())
				Expect(len(p)).To(Equal(2))

				Expect(p[0].BodyHtml).To(ContainSubstring("It's the small iPod with one very big idea"))
				Expect(p[0].CreatedAt).To(Equal("2015-03-28T13:31:19-04:01"))
				Expect(p[0].Handle).To(Equal("ipod-nano"))
				Expect(p[0].Id).To(Equal(int64(632910392)))
				Expect(p[0].ProductType).To(Equal("Cult Products"))
				Expect(p[0].PublishedAt).To(Equal("2007-12-31T19:00:00-05:00"))
				Expect(p[0].PublishedScope).To(Equal("web"))
				Expect(p[0].TemplateSuffix).To(Equal(""))
				Expect(p[0].Title).To(Equal("IPod Nano - 8GB"))
				Expect(p[0].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:02"))
				Expect(p[0].Vendor).To(Equal("Apple"))
				Expect(p[0].Tags).To(Equal("Emotive, Flash Memory, MP3, Music"))

				Expect(len(p[0].Variants)).To(Equal(4))

				Expect(p[0].Variants[0].Barcode).To(Equal("1234_pink"))
				Expect(p[0].Variants[0].CompareAtPrice).To(Equal(""))
				Expect(p[0].Variants[0].CreatedAt).To(Equal("2015-03-28T13:31:19-04:03"))
				Expect(p[0].Variants[0].FulfillmentService).To(Equal("manual"))
				Expect(p[0].Variants[0].Grams).To(Equal(float64(200)))
				Expect(p[0].Variants[0].Id).To(Equal(int64(808950810)))
				Expect(p[0].Variants[0].InventoryManagement).To(Equal("shopify"))
				Expect(p[0].Variants[0].InventoryPolicy).To(Equal("continue"))
				Expect(p[0].Variants[0].Option1).To(Equal("Pink"))
				Expect(p[0].Variants[0].Option2).To(Equal(""))
				Expect(p[0].Variants[0].Option3).To(Equal(""))
				Expect(p[0].Variants[0].Position).To(Equal(1))
				Expect(p[0].Variants[0].Price).To(Equal("199.00"))
				Expect(p[0].Variants[0].ProductId).To(Equal(int64(632910392)))
				Expect(p[0].Variants[0].RequiresShipping).To(BeTrue())
				Expect(p[0].Variants[0].Sku).To(Equal("IPOD2008PINK"))
				Expect(p[0].Variants[0].Taxable).To(BeTrue())
				Expect(p[0].Variants[0].Title).To(Equal("Pink"))
				Expect(p[0].Variants[0].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:04"))
				Expect(p[0].Variants[0].InventoryQuantity).To(Equal(int64(10)))
				Expect(p[0].Variants[0].OldInventoryQuantity).To(Equal(int64(10)))
				Expect(p[0].Variants[0].ImageId).To(Equal(int64(562641783)))
				Expect(p[0].Variants[0].Weight).To(Equal(0.2))
				Expect(p[0].Variants[0].WeightUnit).To(Equal("kg"))

				Expect(p[0].Variants[1].Barcode).To(Equal("1234_red"))
				Expect(p[0].Variants[1].CompareAtPrice).To(Equal(""))
				Expect(p[0].Variants[1].CreatedAt).To(Equal("2015-03-28T13:31:19-04:05"))
				Expect(p[0].Variants[1].FulfillmentService).To(Equal("manual"))
				Expect(p[0].Variants[1].Grams).To(Equal(float64(200)))
				Expect(p[0].Variants[1].Id).To(Equal(int64(49148385)))
				Expect(p[0].Variants[1].InventoryManagement).To(Equal("shopify"))
				Expect(p[0].Variants[1].InventoryPolicy).To(Equal("continue"))
				Expect(p[0].Variants[1].Option1).To(Equal("Red"))
				Expect(p[0].Variants[1].Option2).To(Equal(""))
				Expect(p[0].Variants[1].Option3).To(Equal(""))
				Expect(p[0].Variants[1].Position).To(Equal(2))
				Expect(p[0].Variants[1].Price).To(Equal("199.00"))
				Expect(p[0].Variants[1].ProductId).To(Equal(int64(632910392)))
				Expect(p[0].Variants[1].RequiresShipping).To(BeTrue())
				Expect(p[0].Variants[1].Sku).To(Equal("IPOD2008RED"))
				Expect(p[0].Variants[1].Taxable).To(BeTrue())
				Expect(p[0].Variants[1].Title).To(Equal("Red"))
				Expect(p[0].Variants[1].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:06"))
				Expect(p[0].Variants[1].InventoryQuantity).To(Equal(int64(20)))
				Expect(p[0].Variants[1].OldInventoryQuantity).To(Equal(int64(20)))
				Expect(p[0].Variants[1].ImageId).To(Equal(int64(0)))
				Expect(p[0].Variants[1].Weight).To(Equal(0.2))
				Expect(p[0].Variants[1].WeightUnit).To(Equal("kg"))

				Expect(len(p[0].Options)).To(Equal(1))

				Expect(p[0].Options[0].Id).To(Equal(int64(594680422)))
				Expect(p[0].Options[0].Name).To(Equal("Title"))
				Expect(p[0].Options[0].Position).To(Equal(1))
				Expect(p[0].Options[0].ProductId).To(Equal(int64(632910392)))

				Expect(len(p[0].Images)).To(Equal(2))

				Expect(p[0].Images[0].CreatedAt).To(Equal("2015-03-28T13:31:19-04:07"))
				Expect(p[0].Images[0].Id).To(Equal(int64(850703190)))
				Expect(p[0].Images[0].Position).To(Equal(1))
				Expect(p[0].Images[0].ProductId).To(Equal(int64(632910392)))
				Expect(p[0].Images[0].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:08"))
				Expect(p[0].Images[0].Source).To(Equal("https://cdn.shopify.com/s/files/1/0006/9093/3842/products/ipod-nano.png?v=1427563879"))
				Expect(len(p[0].Images[0].VariantIds)).To(Equal(0))

				Expect(p[0].Images[1].CreatedAt).To(Equal("2015-03-28T13:31:19-04:09"))
				Expect(p[0].Images[1].Id).To(Equal(int64(562641783)))
				Expect(p[0].Images[1].Position).To(Equal(2))
				Expect(p[0].Images[1].ProductId).To(Equal(int64(632910392)))
				Expect(p[0].Images[1].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:10"))
				Expect(p[0].Images[1].Source).To(Equal("https://cdn.shopify.com/s/files/1/0006/9093/3842/products/ipod-nano-2.png?v=1427563879"))
				Expect(p[0].Images[1].VariantIds[0]).To(Equal(int64(808950810)))

				Expect(p[0].Image.CreatedAt).To(Equal("2015-03-28T13:31:19-04:11"))
				Expect(p[0].Image.Id).To(Equal(int64(850703190)))
				Expect(p[0].Image.Position).To(Equal(1))
				Expect(p[0].Image.ProductId).To(Equal(int64(632910392)))
				Expect(p[0].Image.UpdatedAt).To(Equal("2015-03-28T13:31:19-04:12"))
				Expect(p[0].Image.Source).To(Equal("https://cdn.shopify.com/s/files/1/0006/9093/3842/products/ipod-nano.png?v=1427563879"))
				Expect(len(p[0].Image.VariantIds)).To(Equal(0))

				Expect(len(p[1].Variants)).To(Equal(1))
				Expect(len(p[1].Options)).To(Equal(1))
				Expect(len(p[1].Images)).To(Equal(0))

				Expect(mockShopify.Hits()).To(Equal(1))
			})

			g.It("should return a limited list of products", func() {
				mockShopify.SetPayload([]byte(SampleProductsLimitedJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-hash", "oauthom"}
				params := url.Values{}
				params.Add("limit", "1")

				p, err := s.GetProducts(c, params)
				Expect(err).NotTo(HaveOccurred())
				Expect(p).NotTo(BeNil())
				Expect(len(p)).To(Equal(1))

				Expect(p[0].BodyHtml).To(ContainSubstring("The iPod Touch has the iPhone's multi-touch interface"))
				Expect(p[0].CreatedAt).To(Equal("2015-03-28T13:31:19-04:20"))
				Expect(p[0].Handle).To(Equal("ipod-touch"))
				Expect(p[0].Id).To(Equal(int64(921728736)))
				Expect(p[0].ProductType).To(Equal("Cult Products"))
				Expect(p[0].PublishedAt).To(Equal("2008-09-25T20:00:00-04:00"))
				Expect(p[0].PublishedScope).To(Equal("global"))
				Expect(p[0].TemplateSuffix).To(Equal(""))
				Expect(p[0].Title).To(Equal("IPod Touch 8GB"))
				Expect(p[0].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:00"))
				Expect(p[0].Vendor).To(Equal("Apple"))
				Expect(p[0].Tags).To(Equal(""))

				Expect(len(p[0].Variants)).To(Equal(1))

				Expect(p[0].Variants[0].Barcode).To(Equal("1234_black"))
				Expect(p[0].Variants[0].CompareAtPrice).To(Equal(""))
				Expect(p[0].Variants[0].CreatedAt).To(Equal("2015-03-28T13:31:19-04:00"))
				Expect(p[0].Variants[0].FulfillmentService).To(Equal("manual"))
				Expect(p[0].Variants[0].Grams).To(Equal(float64(200)))
				Expect(p[0].Variants[0].Id).To(Equal(int64(447654529)))
				Expect(p[0].Variants[0].InventoryManagement).To(Equal("shopify"))
				Expect(p[0].Variants[0].InventoryPolicy).To(Equal("continue"))
				Expect(p[0].Variants[0].Option1).To(Equal("Black"))
				Expect(p[0].Variants[0].Option2).To(Equal(""))
				Expect(p[0].Variants[0].Option3).To(Equal(""))
				Expect(p[0].Variants[0].Position).To(Equal(1))
				Expect(p[0].Variants[0].Price).To(Equal("199.00"))
				Expect(p[0].Variants[0].ProductId).To(Equal(int64(921728736)))
				Expect(p[0].Variants[0].RequiresShipping).To(BeTrue())
				Expect(p[0].Variants[0].Sku).To(Equal("IPOD2009BLACK"))
				Expect(p[0].Variants[0].Taxable).To(BeTrue())
				Expect(p[0].Variants[0].Title).To(Equal("Black"))
				Expect(p[0].Variants[0].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:00"))
				Expect(p[0].Variants[0].InventoryQuantity).To(Equal(int64(13)))
				Expect(p[0].Variants[0].OldInventoryQuantity).To(Equal(int64(13)))
				Expect(p[0].Variants[0].ImageId).To(Equal(int64(0)))
				Expect(p[0].Variants[0].Weight).To(Equal(0.2))
				Expect(p[0].Variants[0].WeightUnit).To(Equal("kg"))

				Expect(len(p[0].Options)).To(Equal(1))

				Expect(p[0].Options[0].Id).To(Equal(int64(891236591)))
				Expect(p[0].Options[0].Name).To(Equal("Title"))
				Expect(p[0].Options[0].Position).To(Equal(1))
				Expect(p[0].Options[0].ProductId).To(Equal(int64(921728736)))

				Expect(len(p[0].Images)).To(Equal(0))

				Expect(mockShopify.Hits()).To(Equal(1))
			})
		})

		g.Describe("Get One", func() {
			g.It("should get a single product", func() {
				mockShopify.SetPayload([]byte(SampleProductJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-hash", "oauthom"}

				p, err := s.GetProduct("632910392", c, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(p).NotTo(BeNil())

				Expect(p.BodyHtml).To(ContainSubstring("It's the small iPod with one very big idea"))
				Expect(p.CreatedAt).To(Equal("2015-03-28T13:31:19-04:01"))
				Expect(p.Handle).To(Equal("ipod-nano"))
				Expect(p.Id).To(Equal(int64(632910392)))
				Expect(p.ProductType).To(Equal("Cult Products"))
				Expect(p.PublishedAt).To(Equal("2007-12-31T19:00:00-05:00"))
				Expect(p.PublishedScope).To(Equal("web"))
				Expect(p.TemplateSuffix).To(Equal(""))
				Expect(p.Title).To(Equal("IPod Nano - 8GB"))
				Expect(p.UpdatedAt).To(Equal("2015-03-28T13:31:19-04:02"))
				Expect(p.Vendor).To(Equal("Apple"))
				Expect(p.Tags).To(Equal("Emotive, Flash Memory, MP3, Music"))

				Expect(len(p.Variants)).To(Equal(4))

				Expect(p.Variants[0].Barcode).To(Equal("1234_pink"))
				Expect(p.Variants[0].CompareAtPrice).To(Equal(""))
				Expect(p.Variants[0].CreatedAt).To(Equal("2015-03-28T13:31:19-04:03"))
				Expect(p.Variants[0].FulfillmentService).To(Equal("manual"))
				Expect(p.Variants[0].Grams).To(Equal(float64(200)))
				Expect(p.Variants[0].Id).To(Equal(int64(808950810)))
				Expect(p.Variants[0].InventoryManagement).To(Equal("shopify"))
				Expect(p.Variants[0].InventoryPolicy).To(Equal("continue"))
				Expect(p.Variants[0].Option1).To(Equal("Pink"))
				Expect(p.Variants[0].Option2).To(Equal(""))
				Expect(p.Variants[0].Option3).To(Equal(""))
				Expect(p.Variants[0].Position).To(Equal(1))
				Expect(p.Variants[0].Price).To(Equal("199.00"))
				Expect(p.Variants[0].ProductId).To(Equal(int64(632910392)))
				Expect(p.Variants[0].RequiresShipping).To(BeTrue())
				Expect(p.Variants[0].Sku).To(Equal("IPOD2008PINK"))
				Expect(p.Variants[0].Taxable).To(BeTrue())
				Expect(p.Variants[0].Title).To(Equal("Pink"))
				Expect(p.Variants[0].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:04"))
				Expect(p.Variants[0].InventoryQuantity).To(Equal(int64(10)))
				Expect(p.Variants[0].OldInventoryQuantity).To(Equal(int64(10)))
				Expect(p.Variants[0].ImageId).To(Equal(int64(562641783)))
				Expect(p.Variants[0].Weight).To(Equal(0.2))
				Expect(p.Variants[0].WeightUnit).To(Equal("kg"))

				Expect(p.Variants[1].Barcode).To(Equal("1234_red"))
				Expect(p.Variants[1].CompareAtPrice).To(Equal(""))
				Expect(p.Variants[1].CreatedAt).To(Equal("2015-03-28T13:31:19-04:05"))
				Expect(p.Variants[1].FulfillmentService).To(Equal("manual"))
				Expect(p.Variants[1].Grams).To(Equal(float64(200)))
				Expect(p.Variants[1].Id).To(Equal(int64(49148385)))
				Expect(p.Variants[1].InventoryManagement).To(Equal("shopify"))
				Expect(p.Variants[1].InventoryPolicy).To(Equal("continue"))
				Expect(p.Variants[1].Option1).To(Equal("Red"))
				Expect(p.Variants[1].Option2).To(Equal(""))
				Expect(p.Variants[1].Option3).To(Equal(""))
				Expect(p.Variants[1].Position).To(Equal(2))
				Expect(p.Variants[1].Price).To(Equal("199.00"))
				Expect(p.Variants[1].ProductId).To(Equal(int64(632910392)))
				Expect(p.Variants[1].RequiresShipping).To(BeTrue())
				Expect(p.Variants[1].Sku).To(Equal("IPOD2008RED"))
				Expect(p.Variants[1].Taxable).To(BeTrue())
				Expect(p.Variants[1].Title).To(Equal("Red"))
				Expect(p.Variants[1].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:06"))
				Expect(p.Variants[1].InventoryQuantity).To(Equal(int64(20)))
				Expect(p.Variants[1].OldInventoryQuantity).To(Equal(int64(20)))
				Expect(p.Variants[1].ImageId).To(Equal(int64(0)))
				Expect(p.Variants[1].Weight).To(Equal(0.2))
				Expect(p.Variants[1].WeightUnit).To(Equal("kg"))

				Expect(len(p.Options)).To(Equal(1))

				Expect(p.Options[0].Id).To(Equal(int64(594680422)))
				Expect(p.Options[0].Name).To(Equal("Title"))
				Expect(p.Options[0].Position).To(Equal(1))
				Expect(p.Options[0].ProductId).To(Equal(int64(632910392)))

				Expect(len(p.Images)).To(Equal(2))

				Expect(p.Images[0].CreatedAt).To(Equal("2015-03-28T13:31:19-04:07"))
				Expect(p.Images[0].Id).To(Equal(int64(850703190)))
				Expect(p.Images[0].Position).To(Equal(1))
				Expect(p.Images[0].ProductId).To(Equal(int64(632910392)))
				Expect(p.Images[0].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:08"))
				Expect(p.Images[0].Source).To(Equal("https://cdn.shopify.com/s/files/1/0006/9093/3842/products/ipod-nano.png?v=1427563879"))
				Expect(len(p.Images[0].VariantIds)).To(Equal(0))

				Expect(p.Images[1].CreatedAt).To(Equal("2015-03-28T13:31:19-04:09"))
				Expect(p.Images[1].Id).To(Equal(int64(562641783)))
				Expect(p.Images[1].Position).To(Equal(2))
				Expect(p.Images[1].ProductId).To(Equal(int64(632910392)))
				Expect(p.Images[1].UpdatedAt).To(Equal("2015-03-28T13:31:19-04:10"))
				Expect(p.Images[1].Source).To(Equal("https://cdn.shopify.com/s/files/1/0006/9093/3842/products/ipod-nano-2.png?v=1427563879"))
				Expect(p.Images[1].VariantIds[0]).To(Equal(int64(808950810)))

				Expect(p.Image.CreatedAt).To(Equal("2015-03-28T13:31:19-04:11"))
				Expect(p.Image.Id).To(Equal(int64(850703190)))
				Expect(p.Image.Position).To(Equal(1))
				Expect(p.Image.ProductId).To(Equal(int64(632910392)))
				Expect(p.Image.UpdatedAt).To(Equal("2015-03-28T13:31:19-04:12"))
				Expect(p.Image.Source).To(Equal("https://cdn.shopify.com/s/files/1/0006/9093/3842/products/ipod-nano.png?v=1427563879"))
				Expect(len(p.Image.VariantIds)).To(Equal(0))
			})
		})

		g.Describe("Count", func() {
			g.It("should get a products count", func() {
				mockShopify.SetPayload([]byte(SampleProductsCountJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-hash", "oauthom"}

				count, err := s.GetProductsCount(c, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(count).To(Equal(int64(20)))
			})

			g.It("should get a products count that belong to a certain collection", func() {
				mockShopify.SetPayload([]byte(SampleProductsCountCollectionJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-hash", "oauthom"}
				params := url.Values{}
				params.Add("collection_id", "841564295")

				count, err := s.GetProductsCount(c, params)
				Expect(err).NotTo(HaveOccurred())
				Expect(count).To(Equal(int64(1)))
			})
		})
	})
}

const (
	SampleProductsCountJson = `{
	"count": 20
}`

	SampleProductsCountCollectionJson = `{
	"count": 1
}`

	SampleProductsJson = `{
  "products": [
    {
      "body_html": "<p>It's the small iPod with one very big idea: Video. Now the world's most popular music player, available in 4GB and 8GB models, lets you enjoy TV shows, movies, video podcasts, and more. The larger, brighter display means amazing picture quality. In six eye-catching colors, iPod nano is stunning all around. And with models starting at just $149, little speaks volumes.<\/p>",
      "created_at": "2015-03-28T13:31:19-04:01",
      "handle": "ipod-nano",
      "id": 632910392,
      "product_type": "Cult Products",
      "published_at": "2007-12-31T19:00:00-05:00",
      "published_scope": "web",
      "template_suffix": null,
      "title": "IPod Nano - 8GB",
      "updated_at": "2015-03-28T13:31:19-04:02",
      "vendor": "Apple",
      "tags": "Emotive, Flash Memory, MP3, Music",
      "variants": [
        {
          "barcode": "1234_pink",
          "compare_at_price": null,
          "created_at": "2015-03-28T13:31:19-04:03",
          "fulfillment_service": "manual",
          "grams": 200,
          "id": 808950810,
          "inventory_management": "shopify",
          "inventory_policy": "continue",
          "option1": "Pink",
          "option2": null,
          "option3": null,
          "position": 1,
          "price": "199.00",
          "product_id": 632910392,
          "requires_shipping": true,
          "sku": "IPOD2008PINK",
          "taxable": true,
          "title": "Pink",
          "updated_at": "2015-03-28T13:31:19-04:04",
          "inventory_quantity": 10,
          "old_inventory_quantity": 10,
          "image_id": 562641783,
          "weight": 0.2,
          "weight_unit": "kg"
        },
        {
          "barcode": "1234_red",
          "compare_at_price": null,
          "created_at": "2015-03-28T13:31:19-04:05",
          "fulfillment_service": "manual",
          "grams": 200,
          "id": 49148385,
          "inventory_management": "shopify",
          "inventory_policy": "continue",
          "option1": "Red",
          "option2": null,
          "option3": null,
          "position": 2,
          "price": "199.00",
          "product_id": 632910392,
          "requires_shipping": true,
          "sku": "IPOD2008RED",
          "taxable": true,
          "title": "Red",
          "updated_at": "2015-03-28T13:31:19-04:06",
          "inventory_quantity": 20,
          "old_inventory_quantity": 20,
          "image_id": null,
          "weight": 0.2,
          "weight_unit": "kg"
        },
        {
          "barcode": "1234_green",
          "compare_at_price": null,
          "created_at": "2015-03-28T13:31:19-04:00",
          "fulfillment_service": "manual",
          "grams": 200,
          "id": 39072856,
          "inventory_management": "shopify",
          "inventory_policy": "continue",
          "option1": "Green",
          "option2": null,
          "option3": null,
          "position": 3,
          "price": "199.00",
          "product_id": 632910392,
          "requires_shipping": true,
          "sku": "IPOD2008GREEN",
          "taxable": true,
          "title": "Green",
          "updated_at": "2015-03-28T13:31:19-04:00",
          "inventory_quantity": 30,
          "old_inventory_quantity": 30,
          "image_id": null,
          "weight": 0.2,
          "weight_unit": "kg"
        },
        {
          "barcode": "1234_black",
          "compare_at_price": null,
          "created_at": "2015-03-28T13:31:19-04:00",
          "fulfillment_service": "manual",
          "grams": 200,
          "id": 457924702,
          "inventory_management": "shopify",
          "inventory_policy": "continue",
          "option1": "Black",
          "option2": null,
          "option3": null,
          "position": 4,
          "price": "199.00",
          "product_id": 632910392,
          "requires_shipping": true,
          "sku": "IPOD2008BLACK",
          "taxable": true,
          "title": "Black",
          "updated_at": "2015-03-28T13:31:19-04:00",
          "inventory_quantity": 40,
          "old_inventory_quantity": 40,
          "image_id": null,
          "weight": 0.2,
          "weight_unit": "kg"
        }
      ],
      "options": [
        {
          "id": 594680422,
          "name": "Title",
          "position": 1,
          "product_id": 632910392
        }
      ],
      "images": [
        {
          "created_at": "2015-03-28T13:31:19-04:07",
          "id": 850703190,
          "position": 1,
          "product_id": 632910392,
          "updated_at": "2015-03-28T13:31:19-04:08",
          "src": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/products\/ipod-nano.png?v=1427563879",
          "variant_ids": [
          ]
        },
        {
          "created_at": "2015-03-28T13:31:19-04:09",
          "id": 562641783,
          "position": 2,
          "product_id": 632910392,
          "updated_at": "2015-03-28T13:31:19-04:10",
          "src": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/products\/ipod-nano-2.png?v=1427563879",
          "variant_ids": [
            808950810
          ]
        }
      ],
      "image": {
        "created_at": "2015-03-28T13:31:19-04:11",
        "id": 850703190,
        "position": 1,
        "product_id": 632910392,
        "updated_at": "2015-03-28T13:31:19-04:12",
        "src": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/products\/ipod-nano.png?v=1427563879",
        "variant_ids": [
        ]
      }
    },
    {
      "body_html": "<p>The iPod Touch has the iPhone's multi-touch interface, with a physical home button off the touch screen. The home screen has a list of buttons for the available applications.<\/p>",
      "created_at": "2015-03-28T13:31:19-04:00",
      "handle": "ipod-touch",
      "id": 921728736,
      "product_type": "Cult Products",
      "published_at": "2008-09-25T20:00:00-04:00",
      "published_scope": "global",
      "template_suffix": null,
      "title": "IPod Touch 8GB",
      "updated_at": "2015-03-28T13:31:19-04:00",
      "vendor": "Apple",
      "tags": "",
      "variants": [
        {
          "barcode": "1234_black",
          "compare_at_price": null,
          "created_at": "2015-03-28T13:31:19-04:00",
          "fulfillment_service": "manual",
          "grams": 200,
          "id": 447654529,
          "inventory_management": "shopify",
          "inventory_policy": "continue",
          "option1": "Black",
          "option2": null,
          "option3": null,
          "position": 1,
          "price": "199.00",
          "product_id": 921728736,
          "requires_shipping": true,
          "sku": "IPOD2009BLACK",
          "taxable": true,
          "title": "Black",
          "updated_at": "2015-03-28T13:31:19-04:00",
          "inventory_quantity": 13,
          "old_inventory_quantity": 13,
          "image_id": null,
          "weight": 0.2,
          "weight_unit": "kg"
        }
      ],
      "options": [
        {
          "id": 891236591,
          "name": "Title",
          "position": 1,
          "product_id": 921728736
        }
      ],
      "images": [
      ]
    }
  ]
}`

	SampleProductJson = `{
   "product":{
      "body_html":"<p>It's the small iPod with one very big idea: Video. Now the world's most popular music player, available in 4GB and 8GB models, lets you enjoy TV shows, movies, video podcasts, and more. The larger, brighter display means amazing picture quality. In six eye-catching colors, iPod nano is stunning all around. And with models starting at just $149, little speaks volumes.<\/p>",
      "created_at":"2015-03-28T13:31:19-04:01",
      "handle":"ipod-nano",
      "id":632910392,
      "product_type":"Cult Products",
      "published_at":"2007-12-31T19:00:00-05:00",
      "published_scope":"web",
      "template_suffix":null,
      "title":"IPod Nano - 8GB",
      "updated_at":"2015-03-28T13:31:19-04:02",
      "vendor":"Apple",
      "tags":"Emotive, Flash Memory, MP3, Music",
      "variants":[
         {
            "barcode":"1234_pink",
            "compare_at_price":null,
            "created_at":"2015-03-28T13:31:19-04:03",
            "fulfillment_service":"manual",
            "grams":200,
            "id":808950810,
            "inventory_management":"shopify",
            "inventory_policy":"continue",
            "option1":"Pink",
            "option2":null,
            "option3":null,
            "position":1,
            "price":"199.00",
            "product_id":632910392,
            "requires_shipping":true,
            "sku":"IPOD2008PINK",
            "taxable":true,
            "title":"Pink",
            "updated_at":"2015-03-28T13:31:19-04:04",
            "inventory_quantity":10,
            "old_inventory_quantity":10,
            "image_id":562641783,
            "weight":0.2,
            "weight_unit":"kg"
         },
         {
            "barcode":"1234_red",
            "compare_at_price":null,
            "created_at":"2015-03-28T13:31:19-04:05",
            "fulfillment_service":"manual",
            "grams":200,
            "id":49148385,
            "inventory_management":"shopify",
            "inventory_policy":"continue",
            "option1":"Red",
            "option2":null,
            "option3":null,
            "position":2,
            "price":"199.00",
            "product_id":632910392,
            "requires_shipping":true,
            "sku":"IPOD2008RED",
            "taxable":true,
            "title":"Red",
            "updated_at":"2015-03-28T13:31:19-04:06",
            "inventory_quantity":20,
            "old_inventory_quantity":20,
            "image_id":null,
            "weight":0.2,
            "weight_unit":"kg"
         },
         {
            "barcode":"1234_green",
            "compare_at_price":null,
            "created_at":"2015-03-28T13:31:19-04:00",
            "fulfillment_service":"manual",
            "grams":200,
            "id":39072856,
            "inventory_management":"shopify",
            "inventory_policy":"continue",
            "option1":"Green",
            "option2":null,
            "option3":null,
            "position":3,
            "price":"199.00",
            "product_id":632910392,
            "requires_shipping":true,
            "sku":"IPOD2008GREEN",
            "taxable":true,
            "title":"Green",
            "updated_at":"2015-03-28T13:31:19-04:00",
            "inventory_quantity":30,
            "old_inventory_quantity":30,
            "image_id":null,
            "weight":0.2,
            "weight_unit":"kg"
         },
         {
            "barcode":"1234_black",
            "compare_at_price":null,
            "created_at":"2015-03-28T13:31:19-04:00",
            "fulfillment_service":"manual",
            "grams":200,
            "id":457924702,
            "inventory_management":"shopify",
            "inventory_policy":"continue",
            "option1":"Black",
            "option2":null,
            "option3":null,
            "position":4,
            "price":"199.00",
            "product_id":632910392,
            "requires_shipping":true,
            "sku":"IPOD2008BLACK",
            "taxable":true,
            "title":"Black",
            "updated_at":"2015-03-28T13:31:19-04:00",
            "inventory_quantity":40,
            "old_inventory_quantity":40,
            "image_id":null,
            "weight":0.2,
            "weight_unit":"kg"
         }
      ],
      "options":[
         {
            "id":594680422,
            "name":"Title",
            "position":1,
            "product_id":632910392
         }
      ],
      "images":[
         {
            "created_at":"2015-03-28T13:31:19-04:07",
            "id":850703190,
            "position":1,
            "product_id":632910392,
            "updated_at":"2015-03-28T13:31:19-04:08",
            "src":"https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/products\/ipod-nano.png?v=1427563879",
            "variant_ids":[

            ]
         },
         {
            "created_at":"2015-03-28T13:31:19-04:09",
            "id":562641783,
            "position":2,
            "product_id":632910392,
            "updated_at":"2015-03-28T13:31:19-04:10",
            "src":"https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/products\/ipod-nano-2.png?v=1427563879",
            "variant_ids":[
               808950810
            ]
         }
      ],
      "image":{
         "created_at":"2015-03-28T13:31:19-04:11",
         "id":850703190,
         "position":1,
         "product_id":632910392,
         "updated_at":"2015-03-28T13:31:19-04:12",
         "src":"https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/products\/ipod-nano.png?v=1427563879",
         "variant_ids":[

         ]
      }
   }
}`

	SampleProductsLimitedJson = `{
  "products": [
    {
      "body_html": "<p>The iPod Touch has the iPhone's multi-touch interface, with a physical home button off the touch screen. The home screen has a list of buttons for the available applications.<\/p>",
      "created_at": "2015-03-28T13:31:19-04:20",
      "handle": "ipod-touch",
      "id": 921728736,
      "product_type": "Cult Products",
      "published_at": "2008-09-25T20:00:00-04:00",
      "published_scope": "global",
      "template_suffix": null,
      "title": "IPod Touch 8GB",
      "updated_at": "2015-03-28T13:31:19-04:00",
      "vendor": "Apple",
      "tags": "",
      "variants": [
        {
          "barcode": "1234_black",
          "compare_at_price": null,
          "created_at": "2015-03-28T13:31:19-04:00",
          "fulfillment_service": "manual",
          "grams": 200,
          "id": 447654529,
          "inventory_management": "shopify",
          "inventory_policy": "continue",
          "option1": "Black",
          "option2": null,
          "option3": null,
          "position": 1,
          "price": "199.00",
          "product_id": 921728736,
          "requires_shipping": true,
          "sku": "IPOD2009BLACK",
          "taxable": true,
          "title": "Black",
          "updated_at": "2015-03-28T13:31:19-04:00",
          "inventory_quantity": 13,
          "old_inventory_quantity": 13,
          "image_id": null,
          "weight": 0.2,
          "weight_unit": "kg"
        }
      ],
      "options": [
        {
          "id": 891236591,
          "name": "Title",
          "position": 1,
          "product_id": 921728736
        }
      ],
      "images": [
      ]
    }
  ]
}`
)
