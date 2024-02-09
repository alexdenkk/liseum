package token

import (
	"github.com/golang-jwt/jwt/v4"
)

// Claims - JWT token struct
type Claims struct {
	*jwt.StandardClaims

	ID    uint   `json:"id"`
	Login string `json:"login"`
}
