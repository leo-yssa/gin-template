package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"gin-api/rest/app/constant"
	"gin-api/rest/app/crypto"
	"gin-api/rest/app/domain/dto"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	ID string 
	RoleID int 
	jwt.MapClaims
}

func ValidateJwt() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		claims := &dto.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			session := sessions.Default(c)
			secret := session.Get("jwt.secret")
			key, err := crypto.ExtractPrivateKey(secret)
			if err != nil {
				return "", err
			}
			return key.Public(), nil
		})
		if err != nil || token == nil || !token.Valid {
			dto.PanicException(constant.InvalidState)
		}
		c.Next()
	}
}

func ValidateSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		state := session.Get("state")
		if state == nil || state != c.Request.FormValue("state") {
			dto.PanicException(constant.InvalidState)
		}
		c.Next()
	}
}

func GetState() string {
	b := make([]byte, 32)
    rand.Read(b)
    return base64.StdEncoding.EncodeToString(b)
}