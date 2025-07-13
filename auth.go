package main

import (
	"time"

	"go-jwt-auth/config"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var appConfig = config.GetApplicationConfig()

// global variable to hold the JWT signing key
var jwtKey = []byte(appConfig.JWT_TOKEN)

var refreshKey = []byte(appConfig.REFERESH_TOKEN)

func GenerateAuthToken(username string) (string, error) {

	// Set the expiration time for the token to 15 minutes from now
	expirationTime := jwt.NewNumericDate(time.Now().Add(15 * time.Minute))

	// Create the claims, which includes the username and expiration time
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime.Time),
		},
	}
	// Create the token using the claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)

}

func GenerateRefreshToken(username string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshKey)
}
