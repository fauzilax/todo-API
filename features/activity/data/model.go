package data

import (
	"time"
	"todolist-api/features/activity"

	"gorm.io/gorm"
)

type Activity struct {
	ActivityID uint `gorm:"primaryKey"`
	Title      string
	Email      string `gorm:"unique"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func DataToCore(data Activity) activity.Core {
	return activity.Core{
		ActivityID: data.ActivityID,
		Title:      data.Title,
		Email:      data.Email,
	}
}
func CoreToData(core activity.Core) Activity {
	return Activity{
		ActivityID: core.ActivityID,
		Title:      core.Title,
		Email:      core.Email,
	}
}
