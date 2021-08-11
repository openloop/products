package models

type Price struct {
	CurrencyCode string `json:"currency_code" bson:"currency_code" binding:"required"`
	Amount       string `json:"amount" bson:"amount" binding:"required"`
}
