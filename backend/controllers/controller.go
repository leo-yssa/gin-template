package controllers

import (
	"github.com/leo-yssa/gin-template/backend/models"
)

type Controller struct {
	model *models.Model
}

func NewController() *Controller{
	ctl := &Controller{
		model: models.NewModel(),
	}

	return ctl
}