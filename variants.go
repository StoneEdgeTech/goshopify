package goshopify

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type VariantsResponse struct {
	Variants []*Variant `json:"variants"`
}

type VariantResponse struct {
	Variant json.RawMessage `json:"variant"`
}

type Variant struct {
	Barcode              string  `json:"barcode"`
	CompareAtPrice       string  `json:"compare_at_price"`
	CreatedAt            string  `json:"created_at"`
	FulfillmentService   string  `json:"fulfillment_service"`
	Grams                float64 `json:"grams"`
	Id                   int64   `json:"id"`
	InventoryManagement  string  `json:"inventory_management"`
	InventoryPolicy      string  `json:"inventory_policy"`
	Option1              string  `json:"option1"`
	Option2              string  `json:"option2"`
	Option3              string  `json:"option3"`
	Position             int     `json:"position"`
	Price                string  `json:"price"`
	ProductId            int64   `json:"product_id"`
	RequiresShipping     bool    `json:"requires_shipping"`
	Sku                  string  `json:"sku"`
	Taxable              bool    `json:"taxable"`
	Title                string  `json:"title"`
	UpdatedAt            string  `json:"updated_at"`
	InventoryQuantity    int64   `json:"inventory_quantity"`
	OldInventoryQuantity int64   `json:"old_inventory_quantity"`
	ImageId              int64   `json:"image_id"`
	Weight               float64 `json:"weight"`
	WeightUnit           string  `json:"weight_unit"`
}

type VariantUpdate struct {
	Id                 int64 `json:"id"`
	QuantityAdjustment int64 `json:"inventory_quantity_adjustment,omitempty"`
	Quantity           int64 `json:"inventory_quantity,omitempty"`
	OriginalQuantity   int64 `json:"old_inventory_quantity,omitempty"`
}

func NewVariantUpdate(id, oldQty, newQty, adjustment int64) *VariantResponse {
	j, err := json.Marshal(&VariantUpdate{
		Id:                 id,
		Quantity:           newQty,
		OriginalQuantity:   oldQty,
		QuantityAdjustment: adjustment,
	})
	if err != nil {
		return nil
	}

	return &VariantResponse{json.RawMessage(j)}
}

func (s *Shopify) GetVariant(variantId string, creds *Credentials, params url.Values) (*Variant, error) {
	uri, err := s.getUri(fmt.Sprintf(VariantEndpoint, variantId), creds, params)
	if err != nil {
		return nil, err
	}

	var variantResponse *VariantResponse
	err = s.DoResponse("GET", uri, creds, nil, &variantResponse)
	if err != nil {
		return nil, fmt.Errorf("Variant: %s", err.Error())
	}

	var variant Variant
	err = json.Unmarshal(variantResponse.Variant, &variant)
	if err != nil {
		return nil, err
	}

	return &variant, nil
}

func (s *Shopify) PutVariant(variantId string, payload interface{}, creds *Credentials) (*Variant, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	uri, err := s.getUri(fmt.Sprintf(VariantEndpoint, variantId), creds, nil)
	if err != nil {
		return nil, err
	}

	var variantResponse *VariantResponse
	err = s.DoResponse("PUT", uri, creds, jsonPayload, &variantResponse)
	if err != nil {
		return nil, fmt.Errorf("Variant: %s", err.Error())
	}

	var variant Variant
	err = json.Unmarshal(variantResponse.Variant, &variant)
	if err != nil {
		return nil, err
	}

	return &variant, nil
}
