package entities

import "time"

type Inventory struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;"`
	CustomerID string    `gorm:"type:varchar(64);not null;"`
	ProductID  uint64    `gorm:"type:bigint;not null;"`
	IsDeleted  bool      `gorm:"not null;default:false;"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime;"`
}
