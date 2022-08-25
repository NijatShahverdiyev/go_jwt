package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	u "gitlab.com/NijatShahveridev/go_jwt/util"
	"net/http"
	"strings"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func UserDecoder(w http.ResponseWriter, r *http.Request) *User {
	var errorResp u.ErrorResponse
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		msg := errorResp.ErrorMessage(false, "Error thrown while reading body")
		u.SetHeader(w, http.StatusNotAcceptable)
		u.ErrorHandle(w, msg)
	}
	if !strings.Contains(user.Email, "@") {
		msg := errorResp.ErrorMessage(false, "Email address incorrect")
		u.SetHeader(w, http.StatusBadRequest)
		u.ErrorHandle(w, msg)
	}
	return user
}

func (user *User) VerifyEmail() {

}
