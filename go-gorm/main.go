package main

import (
	"github.com/DSXRIIIII/go-utils/go-gorm/sqlInit"
	_type "github.com/DSXRIIIII/go-utils/go-gorm/type"
	"gorm.io/gorm"
)

var (
	client *gorm.DB
	err    error
)

func main() {
	client, err = sqlInit.MySQLDBInit()
	if err != nil {
		panic(err)
	}
	shop := _type.NewShop()
	shop.Find(client)
}
