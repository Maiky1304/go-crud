package main

import (
	"github.com/Maiky1304/go-crud/initializers"
	"github.com/Maiky1304/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
