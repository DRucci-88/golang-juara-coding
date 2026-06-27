package dto

type CheckoutRequest struct {
	UserID uint               `json:"user_id" binding:"required"`
	Items  []OrderItemRequest `json:"items" binding:"required,dive"`
}
