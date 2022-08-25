package models

import (
	"encoding/json"
	u "gitlab.com/NijatShahveridev/go_jwt/util"
	"net/http"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

func AuthDecoder(w http.ResponseWriter, r *http.Request) *Authentication {
	var errorResp u.ErrorResponse
	token := &Authentication{}
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		msg := errorResp.ErrorMessage(false, "Error thrown while reading body")
		u.SetHeader(w, http.StatusNotAcceptable)
		u.ErrorHandle(w, msg)
	}
	return token
}
