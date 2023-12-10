package jwtpkg

import (
	"github.com/golang-jwt/jwt"
	"github.com/narendra121/pkghub/utils"
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

func (jcb *JwtClaimsBuilder) AddCustomClaim(key, val string) *JwtClaimsBuilder {
	jcb.jwtClaims.JwtClaim[key] = val
	return jcb
}

func (jcb *JwtClaimsBuilder) Build() JwtClaims {
	return jcb.jwtClaims
}
