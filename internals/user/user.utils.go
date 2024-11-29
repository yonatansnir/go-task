package user

import (
	"errors"
	"unicode/utf8"
)

func validateUserRequest(user *User) error {
	field := "";
	switch true {
	case user.Email == "":
		field = "email"
	case user.Name == "":
		field = "name";
	case utf8.RuneCountInString(user.Password) < 6:
		return errors.New("password must be greater than 6 chars")
	}

	if (field == "") { 
		return nil 
	}

	return errors.New("Field " + field + " cannot be empty")
}