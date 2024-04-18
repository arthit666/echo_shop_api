package model

type (
	BuyingReq struct {
		CustomerID string
		ProductID  uint64 `json:"itemID" validate:"required,gt=0"`
		Quantity   uint   `json:"quantity" validate:"required,gt=0"`
	}
)
