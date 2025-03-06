package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spiffgreen/basik/models"
	"github.com/spiffgreen/basik/services"
)

// CreateTask handles the creation of a new task
func CreateTask(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	role := userRole.(string)

	if (role != "admin") && userRole.(string) != "manager" {
		c.JSON(403, gin.H{"error": "Not allowed to view this resource"})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTask, err := services.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

// UpdateTask handles the updating of an existing task
func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	taskID := c.Param("id")
	tenantID, _ := c.Get("tenantID")
	nTenantID, _ := strconv.Atoi(tenantID.(string))

	err := services.UpdateTask(nTenantID, taskID, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

// DeleteTask handles the deletion of a task
func DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	tenantID, _ := c.Get("tenantID")
	nTenantID, _ := strconv.Atoi(tenantID.(string))

	if err := services.DeleteTask(nTenantID, taskID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Task updated successfully"})
}

// GetTasks handles fetching all tasks
func GetTasks(c *gin.Context) {
	tenantID, _ := c.Get("tenantID")
	mainTenantID, _ := strconv.Atoi(tenantID.(string))

	tasks, err := services.GetTasks(mainTenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}
