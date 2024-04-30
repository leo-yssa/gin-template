package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"gin-api/app/constant"
	"gin-api/app/domain/error"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ValidateJwt(c *gin.Context) {
	// token := c.GetHeader("Authorization")

}

func ValidateSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		state := session.Get("state")
		if state == nil || state != c.Request.FormValue("state") {
			error.PanicException(constant.InvalidState)
		}
		c.Next()
	}
}

func GetState() string {
	b := make([]byte, 32)
    rand.Read(b)
    return base64.StdEncoding.EncodeToString(b)
}