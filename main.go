package main

import (
	"log"
	"todolist-api/config"
	actData "todolist-api/features/activity/data"
	actHdl "todolist-api/features/activity/handler"
	actSrv "todolist-api/features/activity/services"
	todoData "todolist-api/features/todo/data"
	todoHdl "todolist-api/features/todo/handler"
	todoSrv "todolist-api/features/todo/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	config.Migrate(db)
	aData := actData.New(db)
	aSrv := actSrv.New(aData)
	aHdl := actHdl.New(aSrv)
	tData := todoData.New(db)
	tSrv := todoSrv.New(tData)
	tHdl := todoHdl.New(tSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	// Activity Routes
	e.POST("/activity-groups", aHdl.Create())
	e.PATCH("/activity-groups/:activity_id", aHdl.Update())
	e.DELETE("/activity-groups/:activity_id", aHdl.Delete())
	e.GET("/activity-groups/:activity_id", aHdl.GetOne())
	e.GET("/activity-groups", aHdl.GetAll())
	// Todo Routes
	e.POST("/todo-items", tHdl.Create())
	e.PATCH("/todo-items/:todo_id", tHdl.Update())
	e.DELETE("/todo-items/:todo_id", tHdl.Delete())
	e.GET("/todo-items/:todo_id", tHdl.GetOne())
	e.GET("/todo-items", tHdl.GetAll())

	if err := e.Start(":3030"); err != nil {
		log.Println(err.Error())
	}
}
