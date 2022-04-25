package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"social-network/src/config"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString(config.SecretKey)
}

func ValidateToken(r *http.Request) error {
	tokenStr := getToken(r)

	token, error := jwt.Parse(tokenStr, getSecretKey)
	if error != nil {
		return error
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")

	return nil
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	tokenSplit := strings.Split(token, " ")
	if len(tokenSplit) == 2 {
		return tokenSplit[1]
	}
	return ""
}

func getSecretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected subscription method! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
