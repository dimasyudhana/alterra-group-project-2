package helper

import "github.com/golang-jwt/jwt"

func DecodeJWT(token *jwt.Token) string {
	if token.Valid {
		data := token.Claims.(jwt.MapClaims)
		email := data["email"].(string)

		return email
	}
	return ""
}
