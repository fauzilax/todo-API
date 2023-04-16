package handler

import (
	"log"
	"net/http"
	"strconv"
	"todolist-api/features/todo"

	"github.com/labstack/echo/v4"
)

type todoController struct {
	srv todo.TodoService
}

func New(us todo.TodoService) todo.TodoHandler {
	return &todoController{
		srv: us,
	}
}

// Create implements todo.TodoHandler
func (tc *todoController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := CreateRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		res, err := tc.srv.Create(*ReqToCore(input))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})

		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
			"data":    res,
		})

	}
}

// Delete implements todo.TodoHandler
func (tc *todoController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		todo_idStr := c.Param("todo_id")
		todo_id, _ := strconv.Atoi(todo_idStr)
		err := tc.srv.Delete(uint(todo_id))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "Not Found",
				"message": "Todo with ID " + todo_idStr + " Not Found",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
		})
	}
}

// GetAll implements todo.TodoHandler
func (tc *todoController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramId := c.QueryParam("activity_group_id")
		todo_id, _ := strconv.Atoi(paramId)
		res, err := tc.srv.GetAll(uint(todo_id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		// log.Println(res)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
			"data":    res,
		})
	}
}

// GetOne implements todo.TodoHandler
func (tc *todoController) GetOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		todo_idStr := c.Param("todo_id")
		todo_id, _ := strconv.Atoi(todo_idStr)
		res, err := tc.srv.GetOne(uint(todo_id))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "Not Found",
				"message": "Todo with ID " + todo_idStr + " Not Found",
			})
		}
		// log.Println(res)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
			"data":    res,
		})
	}
}

// Update implements todo.TodoHandler
func (tc *todoController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		todo_idStr := c.Param("todo_id")
		todo_id, _ := strconv.Atoi(todo_idStr)
		input := UpdateRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		log.Println(input, todo_id)
		res, err := tc.srv.Update(uint(todo_id), *ReqToCore(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		// log.Println(res)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
			"data":    res,
		})
	}
}
