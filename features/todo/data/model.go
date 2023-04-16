package data

import (
	"time"
	"todolist-api/features/todo"

	"gorm.io/gorm"
)

type Todo struct {
	TodoID          uint `gorm:"primaryKey"`
	Title           string
	ActivityGroupID uint
	IsActive        bool
	Priority        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func DataToCore(data Todo) todo.Core {
	return todo.Core{
		TodoID:          data.TodoID,
		Title:           data.Title,
		ActivityGroupID: data.ActivityGroupID,
		IsActive:        data.IsActive,
		Priority:        data.Priority,
	}
}
func CoreToData(core todo.Core) Todo {
	return Todo{
		TodoID:          core.TodoID,
		Title:           core.Title,
		ActivityGroupID: core.ActivityGroupID,
		IsActive:        core.IsActive,
		Priority:        core.Priority,
	}
}
