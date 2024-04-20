package util

import (
	"github.com/aadi-1024/ConnectAbility/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJwtToken(uid int, secret []byte, expiry time.Duration) (string, error) {
	claims := models.Claims{}

	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(expiry))
	claims.Id = uid

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return signedToken, err
}
