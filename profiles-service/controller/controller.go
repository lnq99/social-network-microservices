package controller

import (
	"app/service"
)

type Controller struct {
	service *service.Service
}

func NewGinController(service *service.Service) *Controller {
	return &Controller{service}
}
