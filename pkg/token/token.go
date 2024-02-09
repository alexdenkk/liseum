package token

import (
	"github.com/golang-jwt/jwt/v4"
)

// GenerateJWT - generates the JWT token
func GenerateJWT(claims Claims, key []byte) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := t.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT - decode JWT token or return error
func ParseJWT(token string, key []byte) (*Claims, error) {

	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	}

	return &Claims{}, err
}
