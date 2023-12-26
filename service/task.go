package service

import "blog/serializer"

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做，1是已做
}

func (service *CreateTaskService) Create() serializer.Response {

}
