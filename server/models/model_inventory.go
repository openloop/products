package models

type Inventory struct {
	ProductId string `json:"productId,omitempty"`

	VariantId string `json:"variantId,omitempty"`

	Stock int64 `json:"stock,omitempty"`
}
