package activity

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ActivityID uint   `json:"id"`
	Title      string `json:"title"`
	Email      string `json:"email"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type ActivityHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetOne() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}

type ActivityService interface {
	Create(newData Core) (Core, error)
	Update(id uint, title string) (Core, error)
	Delete(id uint) error
	GetOne(id uint) (Core, error)
	GetAll() ([]Core, error)
}

type ActivityData interface {
	Create(newData Core) (Core, error)
	Update(id uint, title string) (Core, error)
	Delete(id uint) error
	GetOne(id uint) (Core, error)
	GetAll() ([]Core, error)
}
