package config

import (
	"fmt"
	"sportify_services/models"
)

func Migration() {
	DB.AutoMigrate(&models.Player{})
	fmt.Println("Berhasil Migrate")
}