package main

import (
	"github.com/Kimoto-Norihiro/gin-api/infra/db"
	"github.com/Kimoto-Norihiro/gin-api/infra/server"
)

var dsn = "n000r111:password@tcp(127.0.0.1:3306)/message_api?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	err := db.CreateMySqlDB(dsn)
	if err != nil {
		panic(err)
	}

	server := server.NewRouter(8080, "localhost")
	server.Start()
}