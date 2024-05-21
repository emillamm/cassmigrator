package main

import (
	"fmt"
	"regexp"
)

func ParseErrorString(str string) error {
	switch {
	case invalidCredentialsRegex.MatchString(str):
		return &InvalidCredentialsError{str}
	default:
		return &UnknownErrorStringError{fmt.Sprintf("unknown error: %s", str)}
	}
}

type InvalidCredentialsError struct{
	Msg string
}
func (e InvalidCredentialsError) Error() string {
	return e.Msg
}
var invalidCredentialsRegex = regexp.MustCompile(`.*username [a-zA-Z0-9\_\-]+ and/or password are incorrect.*`)

type UnknownErrorStringError struct{
	Msg string
}
func (e UnknownErrorStringError) Error() string {
	return e.Msg
}

