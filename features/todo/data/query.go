package data

import (
	"errors"
	"log"
	"todolist-api/features/todo"

	"gorm.io/gorm"
)

type todoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) todo.TodoData {
	return &todoQuery{
		db: db,
	}
}

// Create implements todo.TodoData
func (tq *todoQuery) Create(newData todo.Core) (todo.Core, error) {
	cnv := CoreToData(newData)
	err := tq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return todo.Core{}, errors.New("server error")
	}
	res := DataToCore(cnv)
	res.CreatedAt = cnv.CreatedAt.Format("2006-01-02 15:04:05")
	res.UpdatedAt = cnv.UpdatedAt.Format("2006-01-02 15:04:05")
	return res, nil
}

// Delete implements todo.TodoData
func (tq *todoQuery) Delete(id uint) error {
	qryDelete := tq.db.Delete(&Todo{}, id)
	affRow := qryDelete.RowsAffected
	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("failed to delete user content, data not found")
	}

	return nil
}

// GetAll implements todo.TodoData
func (tq *todoQuery) GetAll(id uint) ([]todo.Core, error) {
	cnv := []Todo{}
	err := tq.db.Where("activity_group_id = ?", id).Find(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []todo.Core{}, errors.New("server error")
	}
	res := []todo.Core{}
	for i, data := range cnv {

		res = append(res, DataToCore(data))
		res[i].CreatedAt = data.CreatedAt.Format("2006-01-02 15:04:05")
		res[i].UpdatedAt = data.UpdatedAt.Format("2006-01-02 15:04:05")
	}
	return res, nil
}

// GetOne implements todo.TodoData
func (tq *todoQuery) GetOne(id uint) (todo.Core, error) {
	cnv := Todo{}
	err := tq.db.Where("todo_id = ?", id).First(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return todo.Core{}, errors.New("server error")
	}

	res := DataToCore(cnv)
	res.CreatedAt = cnv.CreatedAt.Format("2006-01-02 15:04:05")
	res.UpdatedAt = cnv.UpdatedAt.Format("2006-01-02 15:04:05")
	return res, nil
}

// Update implements todo.TodoData
func (tq *todoQuery) Update(id uint, updData todo.Core) (todo.Core, error) {
	cnv := CoreToData(updData)
	log.Println(cnv)
	qry := tq.db.Where("todo_id = ?", id).Updates(&cnv)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return todo.Core{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return todo.Core{}, errors.New("todo not found")
	}

	actData := Todo{}
	err = tq.db.Where("todo_id = ?", id).First(&actData).Error
	if err != nil {
		log.Println("query error", err.Error())
		return todo.Core{}, errors.New("server error")
	}
	res := DataToCore(actData)
	res.CreatedAt = actData.CreatedAt.Format("2006-01-02 15:04:05")
	res.UpdatedAt = actData.UpdatedAt.Format("2006-01-02 15:04:05")
	return res, nil
}
