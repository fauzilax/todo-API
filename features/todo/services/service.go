package services

import (
	"errors"
	"log"
	"strings"
	"todolist-api/features/todo"
)

type todoUseCase struct {
	qry todo.TodoData
}

func New(ud todo.TodoData) todo.TodoService {
	return &todoUseCase{
		qry: ud,
	}
}

// Create implements todo.TodoService
func (tuc *todoUseCase) Create(newData todo.Core) (todo.Core, error) {
	res, err := tuc.qry.Create(newData)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "email already registered"
		} else {
			msg = "server error"
		}
		return todo.Core{}, errors.New(msg)
	}

	return res, nil
}

// Delete implements todo.TodoService
func (tuc *todoUseCase) Delete(id uint) error {
	err := tuc.qry.Delete(uint(id))
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, delete account fail")
	}
	return nil
}

// GetAll implements todo.TodoService
func (tuc *todoUseCase) GetAll(id uint) ([]todo.Core, error) {
	res, err := tuc.qry.GetAll(id)
	if err != nil {
		log.Println("server error or id not found", err.Error())
		return []todo.Core{}, errors.New("id not found")
	}

	return res, nil
}

// GetOne implements todo.TodoService
func (tuc *todoUseCase) GetOne(id uint) (todo.Core, error) {
	res, err := tuc.qry.GetOne(id)
	if err != nil {
		log.Println("server error or id not found", err.Error())
		return todo.Core{}, errors.New("id not found")
	}

	return res, nil
}

// Update implements todo.TodoService
func (tuc *todoUseCase) Update(id uint, updData todo.Core) (todo.Core, error) {
	res, err := tuc.qry.Update(id, updData)
	if err != nil {
		log.Println("server error or id not found", err.Error())
		return todo.Core{}, errors.New("id not found")
	}

	return res, nil
}
