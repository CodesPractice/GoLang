package helper

import (
	"errors"
	"strings"
)

func NameValidation(name string) error {
	if len(name) < 3 {
		return errors.New("name must be at least 3 letters")
	}
	if name == " " {
		return errors.New("this is required field")
	}
	return nil
}

func EmailValidation(e string) error {
	atSign := strings.Contains(e, "@")
	if !atSign {
		return errors.New("invalied email")
	}
	return nil
}
