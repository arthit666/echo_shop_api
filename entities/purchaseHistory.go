package entities

import (
	"time"
)

type PurchaseHistory struct {
	ID                 uint64    `gorm:"primaryKey;autoIncrement;"`
	CustomerID         string    `gorm:"type:varchar(64);not null;"`
	ProductID          uint64    `gorm:"type:bigint;not null;"`
	ProductName        string    `gorm:"type:varchar(64);not null;"`
	ProductDescription string    `gorm:"type:varchar(128);not null;"`
	ProductPrice       uint      `gorm:"not null;"`
	ProductPicture     string    `gorm:"type:varchar(128);not null;"`
	Quantity           uint      `gorm:"not null;"`
	IsBuying           bool      `gorm:"type:boolean;not null;"`
	CreatedAt          time.Time `gorm:"not null;autoCreateTime;"`
}
