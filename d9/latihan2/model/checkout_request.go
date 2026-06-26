package model

type CheckoutRequest struct {
	CustomerName  string     `json:"customer_name" binding:"required"`
	Items         []CartItem `json:"items" binding:"required,dive"`
	PaymentMethod string     `json:"payment_method" binding:"required,oneof='CASH' 'E-WALLET' 'BANK_TRANSFER'"`
}
