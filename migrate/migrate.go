package main

import (
	"fmt"
	"wj-dashboard/db"
	"wj-dashboard/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.MasterRole{}, &model.Admin{})
}
