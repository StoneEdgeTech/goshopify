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

func TestShopify(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Signing", func() {
		g.It("should sign a request with an access token", func() {
			token := "some-access-token"

			req, err := http.NewRequest("GET", "localhost", nil)
			Expect(err).NotTo(HaveOccurred())

			c := &Credentials{"some-id", token}

			c.SignRequest(req)
			h := req.Header.Get(AccessHeader)
			Expect(h).To(ContainSubstring(token))
		})
	})

	g.Describe("Generic Calls", func() {
		var mockShopify *bogus.Bogus

		g.BeforeEach(func() {
			mockShopify = bogus.New()
			mockShopify.Start()
		})

		g.AfterEach(func() {
			mockShopify.Close()
		})

		g.It("should return a valid cart uri", func() {
			cartId := "wakky-wavey-inflatable-waving-arm-guys"
			baseUri := "https://%s.myshopify.com"
			u, _ := url.Parse(fmt.Sprintf(baseUri, cartId))

			s := &Shopify{u.String()}
			c := &Credentials{cartId, "some-token"}

			uri, err := s.getUri(ProductsEndpoint, c, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(uri.Host).To(Equal(fmt.Sprintf("%s.myshopify.com", cartId)))
			Expect(uri.Path).To(Equal(ProductsEndpoint))
		})

		g.It("should return a valid cart uri with params", func() {
			cartId := "wakky-wavey-inflatable-waving-arm-guys"
			k := "collection"
			v := "5555555"
			baseUri := "https://%s.myshopify.com"
			u, _ := url.Parse(fmt.Sprintf(baseUri, cartId))

			s := &Shopify{u.String()}
			c := &Credentials{cartId, "some-token"}

			p := url.Values{}
			p.Add(k, v)

			uri, err := s.getUri(ProductsEndpoint, c, p)
			Expect(err).NotTo(HaveOccurred())
			Expect(uri.Host).To(Equal(fmt.Sprintf("%s.myshopify.com", cartId)))
			Expect(uri.Path).To(Equal(ProductsEndpoint))
			Expect(uri.Query().Get(k)).To(Equal(v))
		})

		g.It("should return a valid cart uri with a custom endpoint", func() {
			cartId := "wakky-wavey-inflatable-waving-arm-guys"
			baseUri := "https://%s.myshopify.com"
			productId := "654321"
			u, _ := url.Parse(fmt.Sprintf(baseUri, cartId))

			s := &Shopify{u.String()}
			c := &Credentials{cartId, "some-token"}

			uri, err := s.getUri(fmt.Sprintf(ProductEndpoint, productId), c, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(uri.Host).To(Equal(fmt.Sprintf("%s.myshopify.com", cartId)))
			Expect(uri.Path).To(Equal(fmt.Sprintf(ProductEndpoint, productId)))
		})

		g.It("should return an error with an invalid cart id", func() {
			mockShopify.SetStatus(http.StatusBadRequest)
			host, port := mockShopify.HostPort()
			u, _ := url.Parse(fmt.Sprintf("http://%s:%s", host, port))

			s := &Shopify{u.String()}
			c := &Credentials{"some-cart-id", "some-token"}

			type TestJson struct {
			}
			var r TestJson
			err := s.DoResponse("GET", u, c, nil, &r)

			Expect(err).To(HaveOccurred())
		})

		g.It("should return an error if it cannot read a sleep time", func() {
			mockShopify.SetPayload([]byte(""))
			mockShopify.SetStatus(429)
			host, port := mockShopify.HostPort()
			u, _ := url.Parse(fmt.Sprintf("http://%s:%s", host, port))

			s := &Shopify{u.String()}
			c := &Credentials{"some-cart-id", "some-token"}

			type TestJson struct {
			}
			var r TestJson
			err := s.DoResponse("GET", u, c, nil, &r)

			Expect(err).To(HaveOccurred())
		})

		g.It("should return an error if bodybytes is 0 length", func() {
			mockShopify.SetPayload([]byte(""))
			mockShopify.SetStatus(http.StatusBadRequest)
			host, port := mockShopify.HostPort()
			u, _ := url.Parse(fmt.Sprintf("http://%s:%s", host, port))

			s := &Shopify{u.String()}
			c := &Credentials{"some-cart-id", "some-token"}

			type TestJson struct {
			}
			var r TestJson
			err := s.DoResponse("GET", u, c, nil, &r)

			Expect(err).To(HaveOccurred())
		})
	})
}
