package entities

import (
	"time"

	"github.com/arthit666/shop_api/pkg/customerCoin/model"
)

type (
	CustomerCoin struct {
		ID         uint64    `gorm:"primaryKey;autoIncrement;"`
		CustomerID string    `gorm:"type:varchar(64);not null;"`
		Amount     int64     `gorm:"not null;"`
		CreatedAt  time.Time `gorm:"not null;autoCreateTime;"`
	}
)

func (p *CustomerCoin) ToPlayerCoinModel() *model.CustomerCoin {
	return &model.CustomerCoin{
		ID:         p.ID,
		CustomerID: p.CustomerID,
		Amount:     p.Amount,
		CreatedAt:  p.CreatedAt,
	}
}
