package error

import (
	"fmt"
	"gin-api/app/constant"
	"gin-api/app/domain/dto"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		key := strArr[0]
		msg := strings.Trim(strArr[1], " ")

		switch key {
		case
			constant.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusBadRequest, dto.BuildResponse_(key, msg, dto.Null()))
			c.Abort()
		case
			constant.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, dto.BuildResponse_(key, msg, dto.Null()))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, dto.BuildResponse_(key, msg, dto.Null()))
			c.Abort()
		}
	}
}
