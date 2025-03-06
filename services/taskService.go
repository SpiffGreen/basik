package services

import (
	"database/sql"
	"errors"

	"github.com/spiffgreen/basik/initializers"
	"github.com/spiffgreen/basik/models"
)

func CreateTask(task *models.Task) (*sql.Row, error) {
	if task == nil {
		return nil, errors.New("task cannot be nil")
	}

	result := initializers.DB.Create(task)
	return result.Row(), nil
}

func UpdateTask(tenantID int, taskID string, updatedTask *models.Task) error {
	if updatedTask == nil {
		return errors.New("task cannot be nil")
	}

	existingTask := &models.Task{}
	result := initializers.DB.First(existingTask, "id = ? AND tenant_id = ?", taskID, tenantID)
	if result.Error != nil {
		return result.Error
	}

	result = initializers.DB.Model(existingTask).Updates(updatedTask)
	return result.Error
}

func DeleteTask(tenantID int, taskID string) error {
	result := initializers.DB.Unscoped().Where("id = ? AND tenant_id = ?", taskID, tenantID).Delete(&models.Task{})
	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}
	return result.Error
}

func GetTaskByID(tenantID, taskID string) (*models.Task, error) {
	var task models.Task
	result := initializers.DB.First(&task, "id = ? AND tenant_id = ?", taskID, tenantID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func GetTasks(tenantID int) ([]models.Task, error) {
	var tasks []models.Task
	println("Here")
	result := initializers.DB.Find(&tasks, "tenant_id = ?", tenantID)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}
