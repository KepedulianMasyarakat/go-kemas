package main

import (
	"go-kemas/config"
	"go-kemas/models"
	"go-kemas/routes"
)

func main() {
	db := config.SetupDB()
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.Program{})
	db.AutoMigrate(&models.Admin{})
	db.AutoMigrate(&models.User{})
	r := routes.SetupRoutes(db)
  r.Run()
}