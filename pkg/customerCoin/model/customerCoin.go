package model

import "time"

type (
	CustomerCoin struct {
		ID         uint64    `json:"id"`
		CustomerID string    `json:"customerID"`
		Amount     int64     `json:"amount"`
		CreatedAt  time.Time `json:"createdAt"`
	}

	CoinAddingReq struct {
		CustomerID string
		Amount     int64 `json:"amount" validate:"required,gt=0"`
	}

	CustomerCoinShowing struct {
		CustomerID string `json:"customerID"`
		Coin       int64  `json:"coin"`
	}
)
