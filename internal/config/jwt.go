package config

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

func NewJWT(v *viper.Viper) *jwtauth.JWTAuth {
	secret := v.GetString("jwt.secret")

	return jwtauth.New(
		"HS256",
		[]byte(secret),
		nil,
	)
}
