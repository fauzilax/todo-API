package handler

import "todolist-api/features/todo"

type CreateRequest struct {
	Title           string `json:"title" form:"title"`
	ActivityGroupID uint   `json:"activity_group_id" form:"activity_group_id"`
	IsActive        bool   `json:"is_active" form:"is_active"`
}

type UpdateRequest struct {
	Title           string `json:"title" form:"title"`
	ActivityGroupID uint   `json:"activity_group_id" form:"activity_group_id"`
	IsActive        bool   `json:"is_active" form:"is_active"`
	Priority        string `json:"priority" form:"priority"`
}

func ReqToCore(data interface{}) *todo.Core {
	res := todo.Core{}

	switch data.(type) {
	case CreateRequest:
		cnv := data.(CreateRequest)
		res.Title = cnv.Title
		res.ActivityGroupID = cnv.ActivityGroupID
		res.IsActive = cnv.IsActive
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Title = cnv.Title
		res.ActivityGroupID = cnv.ActivityGroupID
		res.IsActive = cnv.IsActive
		res.Priority = cnv.Priority
	default:
		return nil
	}

	return &res
}
