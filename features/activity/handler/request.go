package handler

import "todolist-api/features/activity"

type CreateRequest struct {
	Title string `json:"title" form:"title"`
	Email string `json:"email" form:"email"`
}

func ReqToCore(data interface{}) *activity.Core {
	res := activity.Core{}

	switch data.(type) {
	case CreateRequest:
		cnv := data.(CreateRequest)
		res.Title = cnv.Title
		res.Email = cnv.Email
	default:
		return nil
	}

	return &res
}
