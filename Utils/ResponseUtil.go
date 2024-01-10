package Utils

import (
	"github.com/dgrijalva/jwt-go"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Claims struct {
	jwt.StandardClaims
}
