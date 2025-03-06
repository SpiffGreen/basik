package main

import (
	"github.com/spiffgreen/basik/initializers"
	"github.com/spiffgreen/basik/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Tenant{}, &models.Task{}, &models.User{})
}
