package main

import (
	"fmt"

	"github.com/aditya/mvcAPI/Models"
	"github.com/aditya/mvcAPI/Routes"
	"github.com/aditya/mvcAPI/config"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRoutes()
	//running
	r.Run()
}
