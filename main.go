package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spiffgreen/basik/controllers"
	"github.com/spiffgreen/basik/initializers"
	"github.com/spiffgreen/basik/middlewares"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Health check
	r.GET("/health", controllers.CheckHealth)

	// Auth
	r.POST("/auth/register", controllers.AuthRegister)
	r.POST("/auth/login", controllers.AuthLogin)

	// Tenants
	r.POST("/tenants", controllers.CreateTenant)
	r.GET("/tenants", middlewares.AuthMiddleware(), controllers.GetTenants)

	// Task management
	r.GET("/tasks", middlewares.AuthMiddleware(), controllers.GetTasks)
	r.POST("/tasks", middlewares.AuthMiddleware(), controllers.CreateTask)
	r.PUT("/tasks/:id", middlewares.AuthMiddleware(), controllers.UpdateTask)
	r.DELETE("/tasks/:id", middlewares.AuthMiddleware(), controllers.DeleteTask)

	r.Run() // listen and serve on 0.0.0.0:8080
}
