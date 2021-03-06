package goshopify

import (
	"fmt"
	"net/url"
)

type ProductsCount struct {
	Count int64 `json:"count"`
}

type ProductResponse struct {
	Product *Product `json:"product"`
}

type ProductsResponse struct {
	Products []*Product `json:"products"`
}

type Product struct {
	BodyHtml       string     `json:"body_html"`
	CreatedAt      string     `json:"created_at"`
	Handle         string     `json:"handle"`
	Id             int64      `json:"id"`
	ProductType    string     `json:"product_type"`
	PublishedAt    string     `json:"published_at"`
	PublishedScope string     `json:"published_scope"`
	TemplateSuffix string     `json:"template_suffix`
	Title          string     `json:"title"`
	UpdatedAt      string     `json:"updated_at"`
	Vendor         string     `json:"vendor"`
	Tags           string     `json:"tags"`
	Variants       []*Variant `json:"variants"`
	Options        []*Option  `json:"options"`
	Images         []*Image   `json:"images"`
	Image          *Image     `json:"image,omitempty"`
}

type Option struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Position  int    `json:"position"`
	ProductId int64  `json:"product_id"`
}

type Image struct {
	CreatedAt  string  `json:"created_at"`
	Id         int64   `json:"id"`
	Position   int     `json:"position"`
	ProductId  int64   `json:"product_id"`
	UpdatedAt  string  `json:"updated_at"`
	Source     string  `json:"src"`
	VariantIds []int64 `json:"variant_ids"`
}

func (s *Shopify) GetProduct(productId string, creds *Credentials, params url.Values) (*Product, error) {
	uri, err := s.getUri(fmt.Sprintf(ProductEndpoint, productId), creds, params)
	if err != nil {
		return nil, err
	}

	var productResponse *ProductResponse
	err = s.DoRequest("GET", uri, creds, nil, &productResponse)
	if err != nil {
		return nil, fmt.Errorf("Product: %s", err.Error())
	}

	// if no product returned, then just return nil. Do not error out.
	if productResponse == nil {
		return nil, nil
	}

	return productResponse.Product, nil
}

func (s *Shopify) GetProducts(creds *Credentials, params url.Values) ([]*Product, error) {
	uri, err := s.getUri(ProductsEndpoint, creds, params)
	if err != nil {
		return nil, err
	}

	var productsResponse *ProductsResponse
	err = s.DoRequest("GET", uri, creds, nil, &productsResponse)
	if err != nil {
		return nil, fmt.Errorf("Products: %s", err.Error())
	}

	return productsResponse.Products, nil
}

func (s *Shopify) GetProductsCount(creds *Credentials, params url.Values) (*ProductsCount, error) {
	uri, err := s.getUri(ProductsCountEndpoint, creds, params)
	if err != nil {
		return nil, err
	}

	var productCount *ProductsCount
	err = s.DoRequest("GET", uri, creds, nil, &productCount)
	if err != nil {
		return nil, fmt.Errorf("ProductsCount: %s", err.Error())
	}

	return productCount, nil
}
