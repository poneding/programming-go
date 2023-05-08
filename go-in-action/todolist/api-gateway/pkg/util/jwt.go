package util

import (
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/spf13/viper"
)

var jwtSecret = []byte(viper.GetString("server.jwtSecret"))

type Claims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

func GenerateToken(userID uint) (string, error) {
	now := time.Now()
	expire := now.Add(24 * time.Hour)

	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer:    "1004125120-DINGPENG",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
