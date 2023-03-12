package middlewares

import (
	"html"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/iqbalpradipta/modul-registrasi-peserta/config"
	"github.com/iqbalpradipta/modul-registrasi-peserta/models"
)

var secretKey = config.JwtSecretKey

func GenerateJWT(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	return token.SignedString(secretKey)
}


func ExtractBearerToken(r *http.Request) string {
	headerAuthorization := r.Header.Get("Authorization")
	bearerToken := strings.Split(headerAuthorization, " ")
	return html.EscapeString(bearerToken[1])
}

func JWTExtractToken(r *http.Request) (map[string]interface{}, error) {
	tokenString := ExtractBearerToken(r)
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

