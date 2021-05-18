package errors

import (
	"fmt"
	"testing"
)

func BenchmarkCreateErrorTypeE(b *testing.B) {
	var errorTypeError = fmt.Errorf("error")

	for i := 0; i < b.N; i++ {
		_ = NoType.New(errorTypeError)
	}
}

func BenchmarkCreateErrorTypeA(b *testing.B) {
	var errorTypeError = fmt.Errorf("error")

	for i := 0; i < b.N; i ++ {
		_ = NoType.NewF(errorTypeError, "arg1", "arg2")
	}
}

func BenchmarkCreateErrorTypeS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NoType.NewS("error")
	}
}

func BenchmarkCreateErrorTypeSA(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		_ = NoType.NewSA("error", "arg1", "arg2")
	}
}
