package main

import "github.com/HAL-RO-Developer/shutter/model"

func main() {
	db := model.GetDBConn()

	db.DropTableIfExists(&model.Account{})

	db.AutoMigrate(&model.Account{})
}
