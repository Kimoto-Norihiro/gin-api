package main

import (
	"github.com/Kimoto-Norihiro/gin-api/server"
	"github.com/Kimoto-Norihiro/gin-api/utils"
	"github.com/Kimoto-Norihiro/gin-api/models"
)

func main() {
	var err error
	err = database.Init(models.User{})
	if err != nil {
		panic(err)
	}
	err = server.Init()
	if err != nil {
		panic(err)
	}
}