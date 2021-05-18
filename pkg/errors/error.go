package errors

/*
	Base scrumble error object
*/
type errorBase struct {
	errorContext
	errorType
}

/*
	Implement default interface Error() method
*/
func (base errorBase) Error() string {
	if base.errorContext.trace != nil {
		return base.errorContext.trace.Error()
	}

	return base.errorContext.original.Error()
}
