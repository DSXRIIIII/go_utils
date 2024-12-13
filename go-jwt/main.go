package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

type MyCustomClaims struct {
	UserID     int
	Username   string
	GrantScope string
	jwt.StandardClaims
}

// 签名密钥
const sign_key = "hello jwt"

// 随机字符串
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(str_len int) string {
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(rand_bytes)
}

func generateTokenUsingHs256() (string, error) {
	claim := MyCustomClaims{
		UserID:     000001,
		Username:   "Tom",
		GrantScope: "read_user_info",
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Auth_Server",
			Subject:   "Tom",
			ExpiresAt: time.Now().Add(time.Hour).Unix(), //过期时间
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(sign_key))
	return token, err
}

func parseTokenHs256(token_string string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(token_string, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(sign_key), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}

func main() {

	token, err := generateTokenUsingHs256()
	if err != nil {
		panic(err)
	}
	fmt.Println("Token = ", token)

	time.Sleep(time.Second * 2)

	my_claim, err := parseTokenHs256(token)
	if err != nil {
		panic(err)
	}
	fmt.Println("my claim = ", my_claim)

}
