package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

/*
	Default scrumble error wrapper
*/
func Wrap(err error, msg string, args ...interface{}) error {
	wrapped := errors.Wrapf(err, msg, args...)

	if err, ok := err.(errorBase); ok {
		err.errorContext.trace = wrapped
		return err
	}

	context := errorContext{
		original: err,
		trace:    wrapped,
	}

	return errorBase{
		errorContext: context,
		errorType:    NoType,
	}
}

/*
	Get income error type
	if error is scrumble.errors.errorBase or return NoType
*/
func GetType(err error) string {
	if err, ok := err.(errorBase); ok {
		return fmt.Sprintf("%s", err.errorType)
	}

	return fmt.Sprintf("%s", NoType)
}

/*
	Get trace error (as error)
	if error is scrumble.errors.errorBase or return income err
*/
func GetTrace(err error) error {
	if err, ok := err.(errorBase); ok {
		return err.errorContext.trace
	}

	return err
}

/*
	Get trace error (as string)
	if error is scrumble.errors.errorBase or return income err.Error()
*/
func GetTraceToString(err error) string {
	if err == nil {
		return ""
	}

	if err, ok := err.(errorBase); ok {
		return err.errorContext.trace.Error()
	}

	return err.Error()
}

/*
	Get original error (as error)
	if error is scrumble.errors.errorBase or return income err
*/
func GetOriginal(err error) error {
	if err, ok := err.(errorBase); ok {
		return err.errorContext.original
	}

	return err
}

/*
	Get original error (as string)
	if error is scrumble.errors.errorBase or return income err.Error()
*/
func GetOriginalAsString(err error) string {
	if err == nil {
		return ""
	}

	if err, ok := err.(errorBase); ok {
		return err.errorContext.original.Error()
	}

	return err.Error()
}
