package models

type Image struct {
	Src string `json:"src" bson:"src" binding:",omitempty"`
}
