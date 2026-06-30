package model

type OrderItem struct {
	OrderID   uint `gorm:"primaryKey" json:"order_id"`
	ProductID uint `gorm:"primaryKey" json:"product_id"`

	Product Product `gorm:"foreignKey:ProductID;reference:ID" json:"product,omitzero"`

	Quantity int     `gorm:"type:int;not null" json:"quantity"`
	Price    float64 `gorm:"type:numeric(12,2);not null" json:"price"`
}
