package errors

/*
	Error context object,
	include:
		- original error
		  (error that initializes the object)
		- trace as full error message
		  (dynamic error value which change by wrapper)
*/
type errorContext struct {
	original error
	trace    error
}
