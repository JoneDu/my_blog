package jwt

import (
	"errors"
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

func ParseToken(tokenString string, jwtSecret string) (*CustomClaim, error) {
	//解析Token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		//验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})
	// 3. 处理解析错误
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("token格式错误")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token已过期")
		} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, errors.New("token尚未生效")
		}
		return nil, errors.New("token验证失败: " + err.Error())
	}

	// 4. 验证claims
	if claims, ok := token.Claims.(*CustomClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的token")
}
