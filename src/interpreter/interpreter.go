package interpreter

import (
	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/errors"
)

type Interpreter struct {
	environment api.IEnvironment
}

func NewInterpreter() api.IInterpreter {
	o := new(Interpreter)
	o.environment = NewEnvironment()
	return o
}

// IInterpreter interface method
func (i *Interpreter) Interpret(statements []api.IStatement) api.IRuntimeError {
	for _, statement := range statements {
		err := i.execute(statement)
		if err != nil {
			return err
		}
	}

	return nil
}

// statement analogue to the evaluate() method we have for expressions
func (i *Interpreter) execute(statement api.IStatement) api.IRuntimeError {
	return statement.Accept(i)
}

// expression analogue to the execute() method we have for statements
func (i *Interpreter) evaluate(expr api.IExpression) (obj interface{}, err api.IRuntimeError) {
	return expr.Accept(i)
}

func (i *Interpreter) executeBlock(statements []api.IStatement, parentEnv api.IEnvironment) (err api.IRuntimeError) {
	prevEnv := i.environment

	i.environment = parentEnv

	for _, statement := range statements {
		err = i.execute(statement)
		if err != nil {
			i.environment = prevEnv
			return err
		}
	}

	i.environment = prevEnv

	return nil
}

// ------------------------------------------------------------
// Extractions
// ------------------------------------------------------------
func (i *Interpreter) extractNumber(expr interface{}, token api.IToken) (v float64, err api.IRuntimeError) {
	ev, isNum := expr.(api.INumberLiteral)
	if isNum {
		return float64(ev.NumValue()), nil
	}

	return 0, errors.NewRuntimeError(token, "Operand not suitable.")
}

func (i *Interpreter) extractNumbers(left, right interface{}, token api.IToken) (lv, rv float64, err api.IRuntimeError) {
	lfv, isNumL := left.(api.INumberLiteral)
	rfv, isNumR := right.(api.INumberLiteral)
	if isNumL || isNumR {
		// One maybe an integer literal
		var l, r float64
		if !isNumL {
			v, _ := left.(api.IIntegerLiteral)
			l = float64(v.IntValue())
		} else {
			l = lfv.NumValue()
		}

		if !isNumR {
			v, _ := right.(api.IIntegerLiteral)
			r = float64(v.IntValue())
		} else {
			r = rfv.NumValue()
		}

		return l, r, nil
	}

	return 0, 0, errors.NewRuntimeError(token, "Operands not suitable.")
}

func (i *Interpreter) extractInteger(expr interface{}, token api.IToken) (v int, err api.IRuntimeError) {
	ev, isInt := expr.(api.IIntegerLiteral)
	if isInt {
		return ev.IntValue(), nil
	}

	return 0, errors.NewRuntimeError(token, "Operand not suitable.")
}

func (i *Interpreter) extractIntegers(left, right interface{}, token api.IToken) (lv, rv int, err api.IRuntimeError) {
	l, liok := left.(api.IIntegerLiteral)
	r, riok := right.(api.IIntegerLiteral)
	if liok && riok {
		return l.IntValue(), r.IntValue(), nil
	}

	return 0, 0, errors.NewRuntimeError(token, "Operands not suitable.")
}

func (i *Interpreter) extractStrings(left, right interface{}, token api.IToken) (lv, rv string, err api.IRuntimeError) {
	lsv, isStrL := left.(api.IStringLiteral)
	rsv, isStrR := right.(api.IStringLiteral)
	if (isStrL && !isStrR) || (!isStrL && isStrR) {
		return "", "", errors.NewRuntimeError(token, "Both '+' operands must strings.")
	} else if isStrL && isStrR {
		return lsv.StringValue(), rsv.StringValue(), nil
	}
	return "", "", errors.NewRuntimeError(token, "Operands not suitable.")
}

// "false" and "nil" are falsey and everything else is truthy
func (i *Interpreter) isTruthy(obj interface{}) bool {
	if obj == nil { // This should never happen
		return false
	}

	vb, isBool := obj.(api.IBooleanLiteral)
	if isBool {
		return vb.BoolValue()
	}

	_, isNil := obj.(api.INilLiteral)

	if isNil {
		return false
	} else {
		return true
	}
}

// func (i *Interpreter) isEqual(left, right interface{}) bool {
// 	_, isNilL := left.(api.INilLiteral)
// 	_, isNilR := right.(api.INilLiteral)
// 	if isNilL && isNilR {
// 		return true
// 	}
// 	// if objA == nil && objB == nil {
// 	// 	return true
// 	// }
// 	// if objA == nil {
// 	// 	return false
// 	// }

// 	return false
// }
