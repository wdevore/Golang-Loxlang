package errors

import (
	"fmt"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

type RuntimeError struct {
	token   api.IToken
	message string
}

func NewRuntimeError(token api.IToken, message string) api.IRuntimeError {
	o := new(RuntimeError)
	o.token = token
	o.message = message
	return o
}

func (r *RuntimeError) Token() api.IToken {
	return r.token
}

func (r *RuntimeError) Message() string {
	return r.message
}

func (r RuntimeError) String() string {
	return fmt.Sprintf("[line %d] %s", r.token.Line(), r.message)
}
