package model

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required,min=3"`
	Code string `json:"code" binding:"required"`
}
