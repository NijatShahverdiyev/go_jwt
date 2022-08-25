package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	u "gitlab.com/NijatShahveridev/go_jwt/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

func GenerateJWT(email, role string) (string, error) {
	var mySignKey = []byte(os.Getenv("secret_key"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySignKey)
	if err != nil {
		log.Fatalf("something went wrong: %s", err.Error())
	}
	return tokenString, nil
}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	var errorResp u.ErrorResponse

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			u.SetHeader(w, http.StatusForbidden)
			msg := errorResp.ErrorMessage(false, "Missing token")
			u.ErrorHandle(w, msg)
			return
		}
		var secretKey = []byte(os.Getenv("secret_key"))
		token, tokenErr := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return secretKey, nil
		})
		if tokenErr != nil {
			u.SetHeader(w, http.StatusForbidden)
			msg := errorResp.ErrorMessage(false, "Token expired")
			u.ErrorHandle(w, msg)
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {
				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return
			} else if claims["role"] == "user" {
				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return
			}
			u.SetHeader(w, http.StatusUnauthorized)
			msg := errorResp.ErrorMessage(false, "Not Authorized")
			u.ErrorHandle(w, msg)
		}
	}
}

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
