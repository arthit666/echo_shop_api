package main

import (
	"github.com/arthit666/shop_api/config"
	"github.com/arthit666/shop_api/databases"
	"github.com/arthit666/shop_api/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
