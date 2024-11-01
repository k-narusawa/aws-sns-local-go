package domain

import (
	"errors"
	"fmt"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrTopicNotFound = errors.New("topic not found")
	ErrToDoNotFound  = errors.New("todo not found")
	ErrUnexpected    = errors.New("unexpected error")
)

type ErrorCode int

const (
	ErrUserNotFoundCode ErrorCode = iota
	ErrTopicNotFoundCode
	ErrToDoNotFoundCode
	ErrUnexpectedCode
)

var StatusMap = map[ErrorCode]int{
	ErrUserNotFoundCode:  404,
	ErrToDoNotFoundCode:  404,
	ErrTopicNotFoundCode: 404,
	ErrUnexpectedCode:    500,
}

var ErrorCodeMap = map[ErrorCode]string{
	ErrTopicNotFoundCode: "404-001",
	ErrToDoNotFoundCode:  "404-002",
	ErrUserNotFoundCode:  "404-003",
	ErrUnexpectedCode:    "500-001",
}

var ErrorMassageMap = map[ErrorCode]string{
	ErrUserNotFoundCode:  "user not found.",
	ErrToDoNotFoundCode:  "todo not found.",
	ErrTopicNotFoundCode: "topic not found.",
}

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (e Error) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"code": "%s", "message": "%s"}`, ErrorCodeMap[e.Code], e.Message)), nil
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func ToDomainError(err error) Error {
	if err == ErrTopicNotFound {
		return Error{
			Code:    ErrTopicNotFoundCode,
			Message: ErrorMassageMap[ErrTopicNotFoundCode],
		}
	}
	if err == ErrUserNotFound {
		return Error{
			Code:    ErrUserNotFoundCode,
			Message: ErrorMassageMap[ErrUserNotFoundCode],
		}
	}

	if err == ErrToDoNotFound {
		return Error{
			Code:    ErrToDoNotFoundCode,
			Message: ErrorMassageMap[ErrToDoNotFoundCode],
		}
	}

	return Error{
		Code:    ErrUnexpectedCode,
		Message: err.Error(),
	}
}
