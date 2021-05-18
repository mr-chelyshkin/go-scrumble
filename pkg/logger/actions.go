package logger

import (
	"fmt"
	"github.com/mr-chelyshkin/go-scrumble/pkg/errors"
)

/*
	Write string to log with "Info" level
*/
func (l Logger) Info(msg string) {
	l.zap.Info(msg)
}

/*
	Write string and args to log with "Info" level
*/
func (l Logger) InfoF(msg string, args ...interface{}) {
	l.zap.Sugar().Info(msg, args)
}

// -- >

/*
	Write error to log with "Warning" level
*/
func (l Logger) Warning(err error) {
	if err == nil {
		err = fmt.Errorf("")
	}

	logStr := fmt.Sprintf("%s: %s", errors.GetType(err), errors.GetTraceToString(err))
	l.zap.Warn(logStr)
}

/*
	Write string to log with "Warning" level
*/
func (l Logger) WarningS(msg string) {
	l.zap.Warn(msg)
}

/*
	Write error and args to log with "Warning" level
*/
func (l Logger) WarningF(err error, args ...interface{}) {
	if err == nil {
		err = fmt.Errorf("")
	}

	logStr := fmt.Sprintf("%s: %s", errors.GetType(err), errors.GetTraceToString(err))
	l.zap.Sugar().Warn(logStr, args)
}

/*
	Write string and args to log with "Warning" level
*/
func (l Logger) WarningSF(msg string, args ...interface{}) {
	l.zap.Sugar().Warn(msg, args)
}

// -- >

/*
	Write error to log with "Error" level
*/
func (l Logger) Error(err error) {
	if err == nil {
		err = fmt.Errorf("")
	}

	logStr := fmt.Sprintf("%s: %s", errors.GetType(err), errors.GetTraceToString(err))
	l.zap.Error(logStr)
}

/*
	Write string to log "Error" level
*/
func (l Logger) ErrorS(err string) {
	l.zap.Error(err)
}

/*
	Write error and args to log with "Error" level
*/
func (l Logger) ErrorF(err error, args ...interface{}) {
	if err == nil {
		err = fmt.Errorf("")
	}

	logStr := fmt.Sprintf("%s: %s", errors.GetType(err), errors.GetTraceToString(err))
	l.zap.Sugar().Error(logStr, args)
}

/*
	Write string and args to log with "Error" level
*/
func (l Logger) ErrorSF(msg string, args ...interface{}) {
	l.zap.Sugar().Error(msg, args)
}

// -- >

/*
	Write error to log "Fatal" level
*/
func (l Logger) Fatal(err error) {
	if err == nil {
		err = fmt.Errorf("")
	}

	logStr := fmt.Sprintf("%s: %s", errors.GetType(err), errors.GetTraceToString(err))
	l.zap.Fatal(logStr)
}

/*
	Write string to log "Fatal" level
*/
func (l Logger) FatalS(msg string) {
	logStr := fmt.Sprintf("%s", msg)
	l.zap.Fatal(logStr)
}

/*
	Write error and args to log with "Fatal" level
*/
func (l Logger) FatalF(err error, args ...interface{}) {
	if err == nil {
		err = fmt.Errorf("")
	}

	logStr := fmt.Sprintf("%s: %s", errors.GetType(err), errors.GetTraceToString(err))
	l.zap.Sugar().Fatal(logStr, args)
}

/*
	Write string and args to log with "Fatal" level
*/
func (l Logger) FatalSF(msg string, args ...interface{}) {
	logStr := fmt.Sprintf("%s", msg)
	l.zap.Sugar().Fatal(logStr, args)
}

/*
	Write error with to log "Panic" level
*/
func (l Logger) Panic(err error) {
	if err == nil {
		err = fmt.Errorf("")
	}

	logStr := fmt.Sprintf("%s: %s", errors.GetType(err), errors.GetTraceToString(err))
	l.zap.Panic(logStr)
}
