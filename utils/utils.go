package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/Splash07/go_user/config"
	"github.com/Splash07/go_user/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func MD5(text string) string {
	data := []byte(text)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func GenerateToken(userID int, phone, email string) string {
	signKey := []byte(config.Config.JWTSecret.JWTKey)
	claims := model.UserClaims{
		UserID: userID,
		Phone:  phone,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(), //unix timestamp
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(signKey)

	return ss
}
