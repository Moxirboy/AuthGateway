package jwt

import (
	"Auth_service-and-Api_gateway/proto"
	"Auth_service-and-Api_gateway/service/config"
	"Auth_service-and-Api_gateway/service/service/repo"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"

	"time"
)

func GenerateToken(req *proto.UserRequest) (string, error) {
	user, err := repo.GetUser(req.Username, GeneratePasswordHash(req))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(config.Configs().SigningKey))
}
func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return ok, nil
		}
		return []byte(config.Configs().SigningKey), nil
	})
	if err != nil {
		return 0, nil
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		errors.New("token is not type of tokenClaims ")
	}
	return claims.UserId, nil
}

func GeneratePasswordHash(req *proto.UserRequest) string {
	hash := sha1.New()
	hash.Write([]byte(req.Password))
	config.Configs().Salt = "12345"
	return fmt.Sprintf("%x", hash.Sum([]byte(config.Configs().Salt)))
}
