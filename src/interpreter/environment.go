package interpreter

import (
	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/errors"
)

type Environment struct {
	values map[string]interface{}
}

func NewEnvironment() api.IEnvironment {
	o := new(Environment)
	o.values = make(map[string]interface{})
	return o
}

func (e *Environment) Define(name string, obj interface{}) (err api.IRuntimeError) {
	_, ok := e.values[name]
	if !ok {
		e.values[name] = obj
		return nil
	}

	return errors.NewRuntimeError(nil, "Variable '"+name+"' already defined.")
}

func (e *Environment) Get(name api.IToken) (obj interface{}, err api.IRuntimeError) {
	value, ok := e.values[name.Lexeme()]
	if ok {
		return value, nil
	}

	return nil, errors.NewRuntimeError(nil, "Undefined variable '"+name.Lexeme()+"'.")
}

func (e *Environment) Assign(name api.IToken, obj interface{}) (err api.IRuntimeError) {
	_, ok := e.values[name.Lexeme()]
	if ok {
		e.values[name.Lexeme()] = obj
		return nil
	}

	return errors.NewRuntimeError(nil, "Undefined variable '"+name.Lexeme()+"'.")
}
