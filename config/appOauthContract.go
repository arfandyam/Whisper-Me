package config

import "golang.org/x/oauth2"

type AppOauthConfigInterface interface {
	GoogleConfig() oauth2.Config
}