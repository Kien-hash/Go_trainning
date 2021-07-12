package datastruct

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID        int    `json:"ID"`
	Username  string `json:"username"`
	Passwords string `json:"passwords"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}