package utils

import (
	"crypto/rand"
	"time"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

func GenerateRandomSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func ValidateExpiration(expiration int64) bool {
	return expiration > time.Now().Unix()
}

func GetTokenData(token *jwt.Token) jwt.MapClaims {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Errorln("error in parsing token ")
		return nil
	}
	return claims
}

func GetExpDuration(exp int64) int64 {
	return time.Now().Add(time.Minute * time.Duration(exp)).Unix()
}
