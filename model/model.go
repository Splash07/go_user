package model

import (
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type UserClaims struct {
	UserID int    `json:"uid"`
	Phone  string `json:"p"`
	Email  string `json:"e"`
	jwt.StandardClaims
}

type Error struct {
	Code int
	Msg  string
}

type UserUpdateReq struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	*User
	Token string `json:"token"`
}
