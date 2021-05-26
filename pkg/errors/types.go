package errors

import "fmt"

type errorType string

/*
	Pre-define errors types
*/
const (
	NoType     = errorType("No Type")
	Undefined  = errorType("Undefined")
	System     = errorType("System")
	Core       = errorType("Core")
	Initialize = errorType("Initialize")
	Response   = errorType("Response")
	Internal   = errorType("Internal")
	NotFound   = errorType("Not Found")
	Request    = errorType("Request")
	Validation = errorType("Validation")
	Middleware = errorType("Middleware")
)

/*
	Create scrumble.errors.errorBase object from income error, based from selected error type
*/
func (e errorType) New(original error) errorBase {
	context := errorContext{
		original: original,
		trace:    original,
	}

	return errorBase{
		errorContext: context,
		errorType:    e,
	}
}

/*
	Create scrumble.errors.errorBase object from income string, based from selected error type
*/
func (e errorType) NewS(original string) errorBase {
	context := errorContext{
		original: fmt.Errorf(original),
		trace:    fmt.Errorf(original),
	}

	return errorBase{
		errorContext: context,
		errorType:    e,
	}
}

/*
	Create scrumble.errors.errorBase object from income err and args, based from selected error type
*/
func (e errorType) NewF(original error, args ...interface{}) errorBase {
	trace := fmt.Errorf("")
	for _, arg := range args {
		trace = fmt.Errorf("%v %w", arg, trace)
	}

	context := errorContext{
		original: original,
		trace:    fmt.Errorf("%v: %w", trace, original),
	}

	return errorBase{
		errorContext: context,
		errorType:    e,
	}
}

/*
	Create scrumble.errors.errorBase object from income string and args, based from selected error type
*/
func (e errorType) NewSF(original string, args ...interface{}) errorBase {
	trace := fmt.Errorf("")
	for _, arg := range args {
		trace = fmt.Errorf("%v %w", arg, trace)
	}

	context := errorContext{
		original: fmt.Errorf(original),
		trace:    fmt.Errorf("%v: %s", trace, original),
	}

	return errorBase{
		errorContext: context,
		errorType:    e,
	}
}
