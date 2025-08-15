package error

import (
	"errors"
	"fmt"
	"testing"
)

type MyError struct {
	Code int
	Msg  string
}

// error 是一个接口，需要实现 Error() 方法
func (m MyError) Error() string {
	return fmt.Sprintf("MyError: code=%d, msg=%s", m.Code, m.Msg)
}

func doSomething() error {
	return MyError{Code: 404, Msg: "Not Found"}
}

func TestWrappedError(t *testing.T) {
	base := doSomething()
	wrapped := fmt.Errorf("wrapped: %w", base)
	fmt.Println("Unwrap:", errors.Unwrap(wrapped))
	fmt.Println("Is base?", errors.Is(wrapped, base))
}
