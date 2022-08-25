package repository

import (
	"encoding/json"
	"net/http"

	a "gitlab.com/NijatShahveridev/go_jwt/auth"
	"gitlab.com/NijatShahveridev/go_jwt/models"
	u "gitlab.com/NijatShahveridev/go_jwt/util"
)

var errorResp u.ErrorResponse

func SignUp(w http.ResponseWriter, r *http.Request) {
	var dbUser models.User

	user := models.UserDecoder(w, r)
	check := models.GetDB().Table("users").Where("email = ?", user.Email).First(&dbUser).Error
	if dbUser.Email != "" {
		msg := errorResp.ErrorMessage(false, "email already in use")
		u.SetHeader(w, http.StatusForbidden)
		u.ErrorHandle(w, msg)
		return
	}
	user.Password, check = a.GenerateHashPassword(user.Password)
	u.Notify(check, "error in password hash")
	models.GetDB().Create(&user)
	u.SetHeader(w, http.StatusCreated)
	encodeErr := json.NewEncoder(w).Encode(user)
	u.Notify(encodeErr, "Error while encoding user data")

}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	authUser := models.AuthDecoder(w, r)
	models.GetDB().Table("users").Where("email = ?", authUser.Email).First(&newUser)
	if newUser.Email == "" {
		u.SetHeader(w, http.StatusBadRequest)
		msg := errorResp.ErrorMessage(false, "Email not found")
		u.ErrorHandle(w, msg)
		return
	}
	check := a.CheckPasswordHash(authUser.Password, newUser.Password)
	if !check {
		u.SetHeader(w, http.StatusForbidden)
		msg := errorResp.ErrorMessage(false, "Password incorrect")
		u.ErrorHandle(w, msg)
		return
	}
	validToken, generateError := a.GenerateJWT(newUser.Email, newUser.Role)
	if generateError != nil {
		u.SetHeader(w, http.StatusNotAcceptable)
		msg := errorResp.ErrorMessage(false, "Failed to generate token")
		u.ErrorHandle(w, msg)
		return
	}
	var token models.Token
	token.Email = newUser.Email
	token.Role = newUser.Role
	token.TokenString = validToken
	u.SetHeader(w, http.StatusOK)
	encodeErr := json.NewEncoder(w).Encode(token)
	if encodeErr != nil {
		msg := errorResp.ErrorMessage(false, "Error thrown while reading body")
		u.SetHeader(w, http.StatusBadRequest)
		u.ErrorHandle(w, msg)
	}

}

func Index(w http.ResponseWriter, r *http.Request) {

	write, err := w.Write([]byte("Default index page"))
	msg := "Error happened with status code" + string(rune(write))
	u.Notify(err, msg)
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		u.SetHeader(w, http.StatusBadRequest)
		msg := errorResp.ErrorMessage(false, "Not Authorized")
		u.ErrorHandle(w, msg)
		return
	}
	write, err := w.Write([]byte("Admin"))
	msg := "Error happened with status code" + string(rune(write))
	u.Notify(err, msg)
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "user" {
		u.SetHeader(w, http.StatusBadRequest)
		msg := errorResp.ErrorMessage(false, "Not Authorized")
		u.ErrorHandle(w, msg)
		return
	}
	write, err := w.Write([]byte("User"))
	msg := "Error happened with status code" + string(rune(write))
	u.Notify(err, msg)
}
