package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) RegisterUser(c *gin.Context) {
	_user := ctl.model.NewUser()
	if err := c.ShouldBindJSON(&_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user, err := ctl.model.RegisterUser(_user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})	
	} else {
		c.JSON(http.StatusAccepted, gin.H{"message": user.Id})
	}
	return
}