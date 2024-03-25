package middit

import (
	"github.com/golang-jwt/jwt/v5"
)

const name = "user"

func CreateToken(hmacSampleSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": hmacSampleSecret,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(name))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(name), nil
	})
	if err != nil {
		return "", err
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	return claims["name"], nil

}
