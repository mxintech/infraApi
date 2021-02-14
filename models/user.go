package models

import (
	"encoding/json"
	"net/http"
)

// ValidateUser validates incoming json
// based in https://github.com/mxintech/RestAPI/blob/main/user/user.go#L21
func ValidateUser(r *http.Request) (User, error) {
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

type User struct {
	Curp        string `json:"curp" validate:"required,min=18,max=18"`
	FirstPhone  string `json:"first_phone" validate:"required,min=10,max=10"`
	SecondPhone string `json:"second_phone" validate:"required,min=10,max=10"`
	FirstEmail  string `json:"first_email" validate:"required,email"`
	SecondEmail string `json:"second_email" validate:"required,email"`
	CP          string `json:"cp" validate:"required,min=2,max=10"`
}
