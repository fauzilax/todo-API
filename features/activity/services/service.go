package services

import (
	"errors"
	"log"
	"strings"
	"todolist-api/features/activity"
)

type activityUseCase struct {
	qry activity.ActivityData
}

func New(ud activity.ActivityData) activity.ActivityService {
	return &activityUseCase{
		qry: ud,
	}
}

// Create implements activity.ActivityService
func (auc *activityUseCase) Create(newData activity.Core) (activity.Core, error) {
	res, err := auc.qry.Create(newData)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "email already registered"
		} else {
			msg = "server error"
		}
		return activity.Core{}, errors.New(msg)
	}

	return res, nil
}

// Delete implements activity.ActivityService
func (auc *activityUseCase) Delete(id uint) error {
	err := auc.qry.Delete(uint(id))
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, delete account fail")
	}
	return nil
}

// GetAll implements activity.ActivityService
func (auc *activityUseCase) GetAll() ([]activity.Core, error) {
	res, err := auc.qry.GetAll()
	if err != nil {
		log.Println("server error or id not found", err.Error())
		return []activity.Core{}, errors.New("id not found")
	}

	return res, nil
}

// GetOne implements activity.ActivityService
func (auc *activityUseCase) GetOne(id uint) (activity.Core, error) {
	res, err := auc.qry.GetOne(id)
	if err != nil {
		log.Println("server error or id not found", err.Error())
		return activity.Core{}, errors.New("id not found")
	}

	return res, nil
}

// Update implements activity.ActivityService
func (auc *activityUseCase) Update(id uint, title string) (activity.Core, error) {
	res, err := auc.qry.Update(id, title)
	if err != nil {
		log.Println("server error or id not found", err.Error())
		return activity.Core{}, errors.New("id not found")
	}

	return res, nil
}
