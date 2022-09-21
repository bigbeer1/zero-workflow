package jwtx

import (
	"github.com/golang-jwt/jwt"
)

func GetToken(secretKey string, iat int64, seconds int64, uid string, TokenType, nickName, tenantId string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	claims["tokenType"] = TokenType
	claims["nickName"] = nickName
	claims["tenantId"] = tenantId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
