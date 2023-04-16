package data

import (
	"errors"
	"log"
	"todolist-api/features/activity"

	"gorm.io/gorm"
)

type activityQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) activity.ActivityData {
	return &activityQuery{
		db: db,
	}
}

// Create implements activity.ActivityData
func (aq *activityQuery) Create(newData activity.Core) (activity.Core, error) {
	cnv := CoreToData(newData)
	err := aq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return activity.Core{}, errors.New("server error")
	}
	res := DataToCore(cnv)
	res.CreatedAt = cnv.CreatedAt.Format("2006-01-02 15:04:05")
	res.UpdatedAt = cnv.UpdatedAt.Format("2006-01-02 15:04:05")
	res.ActivityID = cnv.ActivityID
	return res, nil
}

// Delete implements activity.ActivityData
func (aq *activityQuery) Delete(id uint) error {
	qryDelete := aq.db.Delete(&Activity{}, id)
	affRow := qryDelete.RowsAffected
	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("failed to delete user content, data not found")
	}

	return nil
}

// GetAll implements activity.ActivityData
func (aq *activityQuery) GetAll() ([]activity.Core, error) {
	cnv := []Activity{}
	err := aq.db.Find(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []activity.Core{}, errors.New("server error")
	}
	res := []activity.Core{}
	for i, data := range cnv {

		res = append(res, DataToCore(data))
		res[i].CreatedAt = data.CreatedAt.Format("2006-01-02 15:04:05")
		res[i].UpdatedAt = data.UpdatedAt.Format("2006-01-02 15:04:05")
	}
	return res, nil
}

// GetOne implements activity.ActivityData
func (aq *activityQuery) GetOne(id uint) (activity.Core, error) {
	cnv := Activity{}
	err := aq.db.Where("activity_id = ?", id).First(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return activity.Core{}, errors.New("server error")
	}

	res := DataToCore(cnv)
	res.CreatedAt = cnv.CreatedAt.Format("2006-01-02 15:04:05")
	res.UpdatedAt = cnv.UpdatedAt.Format("2006-01-02 15:04:05")
	return res, nil
}

// Update implements activity.ActivityData
func (aq *activityQuery) Update(id uint, title string) (activity.Core, error) {
	cnv := Activity{}
	cnv.Title = title
	qry := aq.db.Where("activity_id = ?", id).Updates(&cnv)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return activity.Core{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return activity.Core{}, errors.New("activity not found")
	}

	actData := Activity{}
	err = aq.db.Where("activity_id = ?", id).First(&actData).Error
	if err != nil {
		log.Println("query error", err.Error())
		return activity.Core{}, errors.New("server error")
	}
	res := DataToCore(actData)
	res.CreatedAt = actData.CreatedAt.Format("2006-01-02 15:04:05")
	res.UpdatedAt = actData.UpdatedAt.Format("2006-01-02 15:04:05")
	return res, nil
}
