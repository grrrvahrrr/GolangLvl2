package main

import (
	"github.com/pkg/errors"
)

type MyError struct {
	text string
}

func (e *MyError) Error() string {
	return e.text
}

func CreateError(text string) error {
	return &MyError{text: text}
}

func checkError(text string) error {
	err := CreateError(text)
	if err != nil {
		return errors.Wrap(err, "MyError")
	}
	return nil
}
