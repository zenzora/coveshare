package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// login via email - sends login link, with optional "redirect" to encrypted message

type authClaims struct {
	email string
	jwt.StandardClaims
}

// CreateLoginCode creates a JWT with an email and expiration
func CreateLoginCode(email string) (string, error) {
	claims := authClaims{
		email,
		jwt.StandardClaims{ // Not including an expiration for now
			Issuer: "coveshare",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	keyString := viper.GetString("key")
	ss, err := token.SignedString(keyString)
	if err != nil {
		return ss, nil
	}
	return ss, err
}

// Validate login code
