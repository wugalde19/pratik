
package jwt

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
)

// CustomMapClaims is a reference to jwt.MapClaims type
type CustomMapClaims struct {
	Email string `json:"email"`
	jwt_lib.MapClaims
}
