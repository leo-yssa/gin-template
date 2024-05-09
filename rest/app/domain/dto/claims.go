package dto

import (
	"gin-api/rest/app/crypto"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	ID string `json:"id"`
	RoleID int `json:"role_id"` 
	jwt.MapClaims
}

func NewClaims(id string, roleId int, secret interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := &Claims{
		ID: id,
		RoleID: roleId,
	}
	mapClaims := jwt.MapClaims{}
	mapClaims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	claims.MapClaims = mapClaims
	token.Claims = claims
	key, err := crypto.ExtractPrivateKey(secret)
	if err != nil {
		return "", err
	}
	if signedToken, err := token.SignedString(key); err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}