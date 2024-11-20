package config

import (
	"os"
	"golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

type AppOauthConfig struct {
	GoogleLoginConfig oauth2.Config
}

func NewAppOauthConfig() AppOauthConfigInterface {
	return &AppOauthConfig{}
}

func (appOauth *AppOauthConfig) GoogleConfig() oauth2.Config {
	appOauth.GoogleLoginConfig = oauth2.Config{
		RedirectURL: "http://localhost:3000/auth/google/callback",
		ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return appOauth.GoogleLoginConfig
}