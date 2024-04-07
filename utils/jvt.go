package utils

import (
	"yeni/models"

	"github.com/go-chi/jwtauth/v5"
)

var _tokenAuth *jwtauth.JWTAuth

func InitTokenAuth() {
	_tokenAuth = jwtauth.New("HS256", []byte("asd"), nil)
}

func TokenAuth() *jwtauth.JWTAuth {
	return _tokenAuth
}
func GenerateToken(model models.JwtModel) string {
	mapData, err := StrunctToMap(model)
	if err != nil {
		panic(err)
	}
	_, token, err := _tokenAuth.Encode(mapData)
	if err != nil {
		panic(err)
	}
	return token
}
