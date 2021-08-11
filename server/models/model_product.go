package models

import (
	"time"
)

type Product struct {
	ID          int64      `json:"id" bson:"id" binding:"required"`
	Title       string     `json:"title" bson:"title" binding:"required"`
	BodyHTML    string     `json:"body_html" bson:"body_html" binding:"required"`
	Vendor      string     `json:"vendor" bson:"vendor" binding:"required"`
	ProductType string     `json:"product_type" bson:"product_type" binding:"required"`
	CreatedAt   *time.Time `json:"created_at" bson:"created_at" binding:"required"`
	Handle      string     `json:"handle" bson:"handle" binding:"required"`
	UpdatedAt   *time.Time `json:"updated_at" bson:"updated_at" binding:"required"`
	Tags        string     `json:"tags" bson:"tags" binding:",omitempty"`
	Variants    []Variant  `json:"variants" bson:"variants" binding:",omitempty"`
}
