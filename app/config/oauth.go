package config

import (
	"fmt"
	"gin-api/app/constant"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewOAuth() *oauth2.Config {
	redirectURL := fmt.Sprintf("%s:%s%s", os.Getenv("APP_DOMAIN"), os.Getenv("APP_PORT"), constant.GOOGLE_REDIRECT_URL)
	return &oauth2.Config{
		ClientID: os.Getenv("OAUTH_GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_GOOGLE_CLIENT_SECRET"),
		RedirectURL: redirectURL,
		Scopes: []string{constant.GOOGLE_SCOPE_EMAIL, constant.GOOGLE_SCOPE_PROFILE},
		Endpoint: google.Endpoint,
	}
}