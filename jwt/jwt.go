package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaim struct {
	jwt.RegisteredClaims
	Username string
	UserId   uint
}

func GenerateToken(username string, userId uint, jwtSecret string) (string, error) {
	customClaim := CustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Bruce",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
			NotBefore: jwt.NewNumericDate(time.Now()), // 立即生效
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
		},
		Username: username,
		UserId:   userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaim)
	return token.SignedString([]byte(jwtSecret))
}
