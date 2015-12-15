package goshopify

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestShopifyVariants(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Variants", func() {
		var mockShopify *bogus.Bogus

		g.BeforeEach(func() {
			mockShopify = bogus.New()
			mockShopify.Start()
		})

		g.AfterEach(func() {
			mockShopify.Close()
		})

		g.Describe("Get One", func() {
			g.It("should get a single variant by id", func() {
				mockShopify.SetPayload([]byte(SampleVariantJson))
				mockShopify.SetStatus(http.StatusOK)
				host, port := mockShopify.HostPort()

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{"some-cart-hash", "oauthom"}

				v, err := s.GetVariant("808950810", c, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(v).NotTo(BeNil())

				Expect(v.Barcode).To(Equal("1234_pink"))
				Expect(v.CompareAtPrice).To(Equal(""))
				Expect(v.CreatedAt).To(Equal("2015-03-28T13:31:19-04:03"))
				Expect(v.FulfillmentService).To(Equal("manual"))
				Expect(v.Grams).To(Equal(float64(200)))
				Expect(v.Id).To(Equal(int64(808950810)))
				Expect(v.InventoryManagement).To(Equal("shopify"))
				Expect(v.InventoryPolicy).To(Equal("continue"))
				Expect(v.Option1).To(Equal("Pink"))
				Expect(v.Option2).To(Equal(""))
				Expect(v.Option3).To(Equal(""))
				Expect(v.Position).To(Equal(1))
				Expect(v.Price).To(Equal("199.00"))
				Expect(v.ProductId).To(Equal(int64(632910392)))
				Expect(v.RequiresShipping).To(BeTrue())
				Expect(v.Sku).To(Equal("IPOD2008PINK"))
				Expect(v.Taxable).To(BeTrue())
				Expect(v.Title).To(Equal("Pink"))
				Expect(v.UpdatedAt).To(Equal("2015-03-28T13:31:19-04:04"))
				Expect(v.InventoryQuantity).To(Equal(int64(10)))
				Expect(v.OldInventoryQuantity).To(Equal(int64(10)))
				Expect(v.ImageId).To(Equal(int64(562641783)))
				Expect(v.Weight).To(Equal(0.2))
				Expect(v.WeightUnit).To(Equal("kg"))
			})
		})

		g.Describe("Updating", func() {
			g.It("should set variant quantity", func() {
				mockShopify.AddPath("/").
					SetPayload([]byte(SampleVariantJsonQuantity(5))).
					SetStatus(http.StatusOK).
					SetMethods("PUT")
				host, port := mockShopify.HostPort()
				id := int64(1)
				oldAmount := int64(1)
				newAmount := int64(5)
				storeIdentifier := "some-cart-hash"
				oauthToken := "oauthom"
				update := NewVariantUpdate(id, oldAmount, newAmount, 0)

				s := &Shopify{fmt.Sprintf("http://%s:%s", host, port)}
				c := &Credentials{storeIdentifier, oauthToken}

				v, err := s.PutVariant(strconv.Itoa(int(id)), update, c)
				Expect(err).NotTo(HaveOccurred())
				Expect(v).NotTo(BeNil())
				Expect(v.Id).To(Equal(int64(1)))
				Expect(v.InventoryQuantity).To(Equal(int64(5)))
			})

			g.It("should decrement variant quantity", func() {
				mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte(SampleVariantJsonQuantity(5)))
				}))
				s := &Shopify{mock.URL}
				c := &Credentials{"some-cart-id", "oauthom"}
				id := int64(1)
				adjustAmount := int64(-5)
				update := NewVariantUpdate(id, 0, 0, adjustAmount)
				v, err := s.PutVariant(strconv.Itoa(int(id)), update, c)
				Expect(err).NotTo(HaveOccurred())
				Expect(v).NotTo(BeNil())
				Expect(v.Id).To(Equal(int64(1)))
				Expect(v.Sku).To(Equal("Sku1"))
				Expect(v.InventoryQuantity).To(Equal(int64(5)))
			})

			g.It("should increment variant quantity", func() {
				mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte(SampleVariantJsonQuantity(15)))
				}))
				s := &Shopify{mock.URL}
				c := &Credentials{"some-cart-id", "oauthom"}
				id := int64(1)
				adjustAmount := int64(5)
				update := NewVariantUpdate(id, 0, 0, adjustAmount)
				v, err := s.PutVariant(strconv.Itoa(int(id)), update, c)
				Expect(err).NotTo(HaveOccurred())
				Expect(v).NotTo(BeNil())
				Expect(v.Id).To(Equal(int64(1)))
				Expect(v.Sku).To(Equal("Sku1"))
				Expect(v.InventoryQuantity).To(Equal(int64(15)))
			})

			g.It("should not change variant quantity", func() {
				mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte(SampleVariantJsonQuantity(10)))
				}))
				s := &Shopify{mock.URL}
				c := &Credentials{"some-cart-id", "oauthom"}
				id := int64(1)
				adjustAmount := int64(0)
				update := NewVariantUpdate(id, 0, 0, adjustAmount)
				v, err := s.PutVariant(strconv.Itoa(int(id)), update, c)
				Expect(err).NotTo(HaveOccurred())
				Expect(v).NotTo(BeNil())
				Expect(v.Id).To(Equal(int64(1)))
				Expect(v.Sku).To(Equal("Sku1"))
				Expect(v.InventoryQuantity).To(Equal(int64(10)))
			})
		})

	})
}

func SampleVariantJsonQuantity(quantity int64) string {
	return fmt.Sprintf(`{
   "variant":{
      "barcode":"1234_pink",
      "compare_at_price":null,
      "created_at":"2015-03-28T13:31:19-04:03",
      "fulfillment_service":"manual",
      "grams":200,
      "id":1,
      "inventory_management":"shopify",
      "inventory_policy":"continue",
      "option1":"Pink",
      "option2":null,
      "option3":null,
      "position":1,
      "price":"199.00",
      "product_id":632910392,
      "requires_shipping":true,
      "sku":"Sku1",
      "taxable":true,
      "title":"Pink",
      "updated_at":"2015-03-28T13:31:19-04:04",
      "inventory_quantity":%d,
      "old_inventory_quantity":10,
      "image_id":562641783,
      "weight":0.2,
      "weight_unit":"kg"
   }}`, quantity)
}

const (
	SampleVariantsSkuIdJson = `{
  "variants": [
    {
      "id": 1,
      "sku": "Sku1"
    },
    {
      "id": 2,
      "sku": "Sku2"
    },
    {
      "id": 3,
      "sku": "Sku3"
    },
    {
      "id": 4,
      "sku": "Sku4"
    }
  ]
}`
	SampleVariantJson = `{
   "variant":{
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
   }
}`
)
