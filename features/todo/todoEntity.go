package todo

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	TodoID          uint   `json:"id"`
	Title           string `json:"title"`
	ActivityGroupID uint   `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type TodoHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetOne() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}

type TodoService interface {
	Create(newData Core) (Core, error)
	Update(id uint, updData Core) (Core, error)
	Delete(id uint) error
	GetOne(id uint) (Core, error)
	GetAll(id uint) ([]Core, error)
}

type TodoData interface {
	Create(newData Core) (Core, error)
	Update(id uint, updData Core) (Core, error)
	Delete(id uint) error
	GetOne(id uint) (Core, error)
	GetAll(id uint) ([]Core, error)
}
