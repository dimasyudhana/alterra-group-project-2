package helper

import (
	"log"
	"net/http"

	dependecy "github.com/dimasyudhana/alterra-group-project-2/config/dependcy"
	"github.com/dimasyudhana/alterra-group-project-2/entities"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken() {

}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(passhash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passhash), []byte(password))
}

func CreateWebResponse(code int, message string, data any) any {
	return entities.WebResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func GenerateJWT(id int, dp dependecy.Depend) string {
	var informasi = jwt.MapClaims{}
	informasi["id"] = id
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, informasi)
	resultToken, err := rawToken.SignedString([]byte(dp.Config.JwtSecret))
	if err != nil {
		log.Println("generate jwt error ", err.Error())
		return ""
	}
	return resultToken
}

func DeleteCookieCSRF(c echo.Context) {
	deletecookie := &http.Cookie{
		Name:   "_csrf",
		Path:   "/",
		MaxAge: -1,
	}
	c.SetCookie(deletecookie)
}
