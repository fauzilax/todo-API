package handler

import (
	"net/http"
	"strconv"
	"todolist-api/features/activity"

	"github.com/labstack/echo/v4"
)

type activityController struct {
	srv activity.ActivityService
}

func New(us activity.ActivityService) activity.ActivityHandler {
	return &activityController{
		srv: us,
	}
}

// Create implements activity.ActivityHandler
func (ac *activityController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := CreateRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		res, err := ac.srv.Create(*ReqToCore(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		// log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
			"data":    res,
		})
	}
}

// Delete implements activity.ActivityHandler
func (ac *activityController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		activity_idStr := c.Param("activity_id")
		activity_id, _ := strconv.Atoi(activity_idStr)
		err := ac.srv.Delete(uint(activity_id))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "Not Found",
				"message": "Activity with ID " + activity_idStr + " Not Found",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
		})
	}
}

// GetAll implements activity.ActivityHandler
func (ac *activityController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ac.srv.GetAll()
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

// GetOne implements activity.ActivityHandler
func (ac *activityController) GetOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		activity_idStr := c.Param("activity_id")
		activity_id, _ := strconv.Atoi(activity_idStr)
		res, err := ac.srv.GetOne(uint(activity_id))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "Not Found",
				"message": "Activity with ID " + activity_idStr + " Not Found",
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

// Update implements activity.ActivityHandler
func (ac *activityController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		activity_idStr := c.Param("activity_id")
		activity_id, _ := strconv.Atoi(activity_idStr)
		input := CreateRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		res, err := ac.srv.Update(uint(activity_id), input.Title)
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
