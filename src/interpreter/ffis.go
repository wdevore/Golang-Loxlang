package interpreter

import (
	"time"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

// Some are FFIs and some are implementations

// ~~~---~~~---~~~---~~~---~~~---~~~---~~~---~~~---~~~---~~~---
// Clock
// ~~~---~~~---~~~---~~~---~~~---~~~---~~~---~~~---~~~---~~~---
type ClockCallable struct {
}

func NewClockCallable() api.ICallable {
	o := new(ClockCallable)
	return o
}

func (c *ClockCallable) Arity() int {
	return 0
}

func (c *ClockCallable) Call(interpreter api.IInterpreter, arguments []interface{}) (obj interface{}, err api.IRuntimeError) {
	return time.Now().UnixNano() / int64(time.Millisecond), nil
}

func (c ClockCallable) String() string {
	return "<native fn>"
}

// ~~~---~~~---~~~---~~~---~~~---~~~---~~~---~~~---~~~---~~~---
// Function
// ~~~---~~~---~~~---~~~---~~~---~~~---~~~---~~~---~~~---~~~---
type FunctionCallable struct {
	declaration api.IStatement
}

func NewFunctionCallable(declaration api.IStatement) api.ICallable {
	o := new(FunctionCallable)
	o.declaration = declaration
	return o
}

func (c *FunctionCallable) Arity() int {
	return len(c.declaration.Parameters())
}

func (c *FunctionCallable) Call(interpreter api.IInterpreter, arguments []interface{}) (obj interface{}, err api.IRuntimeError) {
	environment := NewEnvironmentEnclosing(interpreter.Globals())

	for i, parm := range c.declaration.Parameters() {
		environment.Define(parm.Lexeme(), arguments[i])
	}

	err = interpreter.ExecuteBlock(c.declaration.Body(), environment)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c FunctionCallable) String() string {
	return "<fn " + c.declaration.Name().Lexeme() + ">"
}
