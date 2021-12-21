package interpreter

import (
	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/errors"
)

type Environment struct {
	enclosing api.IEnvironment
	values    map[string]interface{}
}

// for the global scope’s environment
func NewEnvironment() api.IEnvironment {
	o := new(Environment)
	o.values = make(map[string]interface{})
	return o
}

// local scope nested inside the given outer one
func NewEnvironmentEnclosing(enclosing api.IEnvironment) api.IEnvironment {
	o := new(Environment)
	o.values = make(map[string]interface{})
	o.enclosing = enclosing
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

func (e *Environment) Get(name api.IToken) (value interface{}, err api.IRuntimeError) {
	value, ok := e.values[name.Lexeme()]
	if ok {
		return value, nil
	}

	if e.enclosing != nil {
		return e.enclosing.Get(name)
	}

	return nil, errors.NewRuntimeError(name, "Undefined variable '"+name.Lexeme()+"'.")
}

func (e *Environment) Assign(name api.IToken, value interface{}) (err api.IRuntimeError) {
	_, ok := e.values[name.Lexeme()]
	if ok {
		e.values[name.Lexeme()] = value
		return nil
	}

	if e.enclosing != nil {
		e.enclosing.Assign(name, value)
		return nil
	}

	return errors.NewRuntimeError(name, "Undefined variable '"+name.Lexeme()+"'.")
}
