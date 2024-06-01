package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

// Error errors error type
type Error struct {
	Code        string
	Message     string
	Field       string
	Description string
	Cause       error
	HttpCode    int
}

func New(code, message, description string, httpCode int) Error {
	return Error{
		Code:        code,
		Message:     message,
		Description: description,
		HttpCode:    httpCode,
	}
}

func (e Error) Error() string {
	if e.Cause == nil {
		return fmt.Sprintf("[%s] %s (on %s) : %s",
			e.Code, e.Message, e.Field, e.Description)
	}

	return fmt.Sprintf("[%s] %s (on %s) %s, cause: %s",
		e.Code, e.Message, e.Field, e.Description, e.Cause)
}

func (e Error) SetCauseWrap(err error, message string) Error {
	e.Cause = errors.Wrap(err, message)
	return e
}

func (e Error) SetCause(err error) Error {
	e.Cause = errors.WithStack(err)
	return e
}

func (e Error) SetField(field string) Error {
	e.Field = field
	return e
}

func (e Error) OverrideHttpStatus(status int) Error {
	e.HttpCode = status
	return e
}
