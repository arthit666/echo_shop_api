package main

import (
	"github.com/arthit666/shop_api/config"
	"github.com/arthit666/shop_api/databases"
	"github.com/arthit666/shop_api/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	tx := db.Connect().Begin()

	playerMigration(tx)
	adminMigration(tx)
	itemMigration(tx)
	playerCoinMigration(tx)
	inventoryMigration(tx)
	purchaseHistoryMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func playerMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Customer{})
}

func adminMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Admin{})
}

func itemMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Product{})
}

func playerCoinMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.CustomerCoin{})
}

func inventoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Inventory{})
}

func purchaseHistoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PurchaseHistory{})
}
