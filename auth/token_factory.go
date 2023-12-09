package auth

import (
	"fmt"
	"time"

	"github.com/narendra121/pkghub/utils"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

type CustomTokenValidationFunc func(username string) bool

type TokenFactory interface {
	GenerateSignedToken(expiry int64, signInSalt string, customValidtor CustomTokenValidationFunc) string
	IsTokenValid(signedToken, signInSalt string, customValidtor CustomTokenValidationFunc) bool
	RefreshToken(signedToken string, signInSalt string, tokenExp int64, customValidtor CustomTokenValidationFunc) string
}

func NewTokenFactory(tokenType interface{}) TokenFactory {
	fmt.Printf("%T", tokenType)
	switch tokenType.(type) {
	case *JwtClaims:
		return tokenType.(*JwtClaims)
	default:
		return nil
	}
}

func (j *JwtClaims) GenerateSignedToken(expiry int64, signInSalt string, customValidtor CustomTokenValidationFunc) string {
	if !j.IsCustomValidatorSuccess(customValidtor) {
		return ""
	}

	j.JwtClaim[utils.JWT_CLAIMS_TOKEN_EXP_KEY] = utils.GetExpDuration(expiry)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &j.JwtClaim)

	signedToken, err := token.SignedString([]byte(signInSalt))
	if err != nil {
		log.Errorln("error in generating signed token ", err)
		return ""
	}

	return signedToken
}

func (j *JwtClaims) IsTokenValid(signedToken, signInSalt string, customValidtor CustomTokenValidationFunc) bool {

	token, isValid := j.validateSignInMethod(signedToken, signInSalt)
	if !isValid || !token.Valid {
		return false
	}
	exp := j.JwtClaim[utils.JWT_CLAIMS_TOKEN_EXP_KEY].(int64)

	if !utils.ValidateExpiration(int64(exp)) {
		log.Errorln("token got expired")
		return false
	}
	if !j.IsCustomValidatorSuccess(customValidtor) {
		return false
	}

	return true
}

func (j *JwtClaims) RefreshToken(signedToken string, signInSalt string, tokenExp int64, customValidtor CustomTokenValidationFunc) string {
	if !j.IsTokenValid(signedToken, signInSalt, customValidtor) {
		return ""
	}
	tokenExp = time.Now().Add(time.Minute * time.Duration(tokenExp)).Unix()
	j.JwtClaim[utils.JWT_CLAIMS_TOKEN_EXP_KEY] = tokenExp

	return j.GenerateSignedToken(tokenExp, signInSalt, nil)
}

func (j *JwtClaims) validateSignInMethod(signedToken, signInSalt string) (*jwt.Token, bool) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token signing method is not valid: %v", token.Header["alg"])
		}
		return []byte(signInSalt), nil
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

func (j *JwtClaims) IsCustomValidatorSuccess(validator CustomTokenValidationFunc) bool {
	if validator != nil {
		userName, ok := j.JwtClaim[utils.JWT_CLAIMS_USERNAME_KEY]
		if !ok {
			return false
		}
		if !validator(userName.(string)) {
			return false
		}
	}
	return true

}
