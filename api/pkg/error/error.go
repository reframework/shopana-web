package appErrors

import (
	"fmt"
	"runtime"
	"strings"
)

type AppErrorCode string

const VERSION = "1.0"

const (
	INTERNAL_ERROR  AppErrorCode = "internal_error"
	INVALID_REQUEST AppErrorCode = "invalid_request"
	ACCESS_DENIED   AppErrorCode = "access_denied"
	ALREADY_EXISTS  AppErrorCode = "already_exists"
	NOT_FOUND       AppErrorCode = "not_found"
	UNAUTHORIZED    AppErrorCode = "unauthorized"
)

var (
	Internal       = New(INTERNAL_ERROR, "Internal error", VERSION)
	SchemaNotFound = New(INTERNAL_ERROR, "Schema not found", VERSION)
	InvalidRequest = New(INVALID_REQUEST, "Invalid request", VERSION)
	Unauthorized   = New(UNAUTHORIZED, "Unauthorized", VERSION)
	AccessDenied   = New(ACCESS_DENIED, "Access denied", VERSION)
	AlreadyExists  = New(ALREADY_EXISTS, "Entry already exists", VERSION)
	NotFound       = New(NOT_FOUND, "Entry not found", VERSION)
)

type AppError struct {
	Caller       string       `json:"caller"`
	Code         AppErrorCode `json:"code"`
	Descriptions []string     `json:"description"`
	Message      string       `json:"message"`
	Version      string       `json:"version"`
}

func New(code AppErrorCode, message string, version string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Version: version,
	}
}

func (e AppError) Error() string {
	return fmt.Sprintf(e.Message)
}

func (e AppError) New(err error, descriptions ...string) *AppError {
	newError := New(e.Code, e.Message, e.Version)

	if err != nil {
		newError.Descriptions = append(newError.Descriptions, err.Error())
	}

	if len(descriptions) > 0 {
		newError.Descriptions = append(newError.Descriptions, descriptions...)
	}

	_, file, line, _ := runtime.Caller(1)
	newError.Caller = fmt.Sprintf("%s:%d", file, line)

	// TODO: print based on env
	fmt.Printf("\nApp Error: \n%s. \n%s", strings.Join(newError.Descriptions, "\n"), newError.Caller)
	return newError
}
