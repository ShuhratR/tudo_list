package services

import (
	"errors"
	"time"

	"github.com/todo_list_shuhrat/internal/models"
)

var tasks = []models.Task{}

func CreateTask(title, description string, dueDate time.Time) models.Task {
	task := models.Task{
		ID:          len(tasks) + 1,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		IsCompleted: false,
	}
	tasks = append(tasks, task)
	return task
}

func GetTasks() []models.Task {
	return tasks
}

func UpdateTask(id int, title, description string, dueDate time.Time, isCompleted bool) (*models.Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = title
			tasks[i].Description = description
			tasks[i].DueDate = dueDate
			tasks[i].IsCompleted = isCompleted
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
