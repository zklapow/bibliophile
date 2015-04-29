//go:generate ffjson book.go

package models

type JsonError struct {
    Id      string
    Message string
}

func NewJsonError(message string) *JsonError {
    return &JsonError{Message: message}
}