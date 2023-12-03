package auth

import (
	"fmt"
	"pkg-hub/utils"
	"time"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

type TokenFactory interface {
	GenerateSignedToken() string
	IsTokenValid(signedToken string) bool
	RefreshToken(signedToken string) string
}

func NewTokenFactory(tokenType interface{}) TokenFactory {
	switch tokenType.(type) {
	case JwtClaims:
		return tokenType.(*JwtClaims)
	default:
		return nil
	}
}

func (j *JwtClaims) GenerateSignedToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &j.JwtClaim)

	signedToken, err := token.SignedString([]byte(j.JwtClaim[utils.JWT_CLAIMS_SIGNIN_SALT_KEY].(string)))
	if err != nil {
		log.Errorln("error in generating signed token ", err)
		return ""
	}

	return signedToken
}

func (j *JwtClaims) IsTokenValid(signedToken string) bool {
	token, isValid := j.validateSignInMethod(signedToken)
	if !isValid {
		return false
	}
	tokenData := utils.GetTokenData(token)
	if tokenData == nil {
		log.Errorln("error in parsing token data: ")
		return false
	}
	if !utils.ValidateExpiration(int64(tokenData[utils.JWT_CLAIMS_TOKEN_EXP_KEY].(float64))) {
		log.Errorln("token got expired")
		return false
	}
	return true
}

func (j *JwtClaims) RefreshToken(signedToken string) string {
	token, isValid := j.validateSignInMethod(signedToken)
	if !isValid {
		return ""
	}

	tokenData := utils.GetTokenData(token)
	if tokenData == nil {
		log.Errorln("error in parsing token data: ")
		return ""
	}

	if !utils.ValidateExpiration(int64(tokenData[utils.JWT_CLAIMS_TOKEN_EXP_KEY].(float64))) {
		log.Errorln("token got expired")
		return ""
	}

	tokenExp := time.Now().Add(time.Minute * time.Duration(utils.JWT_REFRESH_TOKEN_EXPIRY)).Unix()
	j.JwtClaim[utils.JWT_CLAIMS_TOKEN_EXP_KEY] = tokenExp
	return j.GenerateSignedToken()
}

func (j *JwtClaims) validateSignInMethod(signedToken string) (*jwt.Token, bool) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token signing method is not valid: %v", token.Header["alg"])
		}
		return []byte(j.JwtClaim[utils.JWT_CLAIMS_SIGNIN_SALT_KEY].(string)), nil
	})
	if err != nil {
		log.Errorln("error in validating signed token ", err)
		return nil, false
	}
	if !token.Valid {
		log.Errorln("invalid token ", err)
		return nil, false
	}
	return token, true
}
