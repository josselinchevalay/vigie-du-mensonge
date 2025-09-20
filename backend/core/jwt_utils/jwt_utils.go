package jwt_utils

import (
	"fmt"
	"time"
	"vdm/core/locals"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	AuthedUser locals.AuthedUser
	jwt.RegisteredClaims
}

func GenerateJWT(authedUser locals.AuthedUser, secretKey []byte, expiry time.Time) (string, error) {
	claims := Claims{
		AuthedUser: authedUser,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiry),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ParseJWT(tokenString string, secretKey []byte) (locals.AuthedUser, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return locals.AuthedUser{}, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.AuthedUser, nil
	}

	return locals.AuthedUser{}, fmt.Errorf("invalid token")
}
