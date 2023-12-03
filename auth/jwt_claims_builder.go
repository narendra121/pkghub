package auth

import (
	"pkg-hub/utils"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	JwtClaim jwt.MapClaims
}

type JwtClaimsBuilder struct {
	jwtClaims JwtClaims
}

func NewJwtBuilder() *JwtClaimsBuilder {
	return &JwtClaimsBuilder{jwtClaims: JwtClaims{JwtClaim: jwt.MapClaims{}}}
}

func (jcb *JwtClaimsBuilder) AddUserName(userName string) *JwtClaimsBuilder {
	jcb.jwtClaims.JwtClaim[utils.JWT_CLAIMS_USERNAME_KEY] = userName
	return jcb
}

func (jcb *JwtClaimsBuilder) AddEmail(email string) *JwtClaimsBuilder {
	jcb.jwtClaims.JwtClaim[utils.JWT_CLAIMS_EMAIL_KEY] = email
	return jcb
}

func (jcb *JwtClaimsBuilder) AddSignInSalt(signInSalt string) *JwtClaimsBuilder {
	jcb.jwtClaims.JwtClaim[utils.JWT_CLAIMS_SIGNIN_SALT_KEY] = signInSalt
	return jcb
}

func (jcb *JwtClaimsBuilder) AddTokenExpiry(tokenExp int64) *JwtClaimsBuilder {
	exp := time.Now().Add(time.Minute * time.Duration(tokenExp)).Unix()
	jcb.jwtClaims.JwtClaim[utils.JWT_CLAIMS_TOKEN_EXP_KEY] = exp
	return jcb
}

func (jcb *JwtClaimsBuilder) Build() JwtClaims {
	return jcb.jwtClaims
}

func (jcb *JwtClaimsBuilder) validate(val interface{}) bool {

	switch val.(type) {
	case string:
		return strings.EqualFold(val.(string), "")
	case int:
		return val.(int) == 0
	default:
		return false
	}

}
