package payload

import "github.com/golang-jwt/jwt"

type ClaimsCustom struct {
	UserId   string
	Username string
	jwt.StandardClaims
}
