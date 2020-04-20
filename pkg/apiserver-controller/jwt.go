package apiserver_controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func JwtDefender(endpoint func(http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte("quarklt-jwt"), nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {

			fmt.Fprintf(w, "Not Authorized")
		}
	}
}
func GenerateJWT(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = username
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	tokenString, err := token.SignedString([]byte("quarklt-jwt"))

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return ""
	}

	return tokenString
}
