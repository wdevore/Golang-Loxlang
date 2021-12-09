package interpreter

import (
	"log"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/errors"
	"github.com/wdevore/RISCV-Meta-Assembler/src/scanner/literals"
)

type Interpreter struct {
}

func NewInterpreter() api.IInterpreter {
	o := new(Interpreter)
	return o
}

// IInterpreter interface method
func (i *Interpreter) Interpret(exprV api.IExpression) api.IRuntimeError {
	value, err := i.evaluate(exprV)
	if err != nil {
		return err
	}

	log.Printf("%v", value)

	return nil
}

func (i *Interpreter) VisitLiteralExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	return exprV.Value(), nil
}

func (i *Interpreter) VisitGroupingExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	return i.evaluate(exprV.Expression())
}

func (i *Interpreter) VisitUnaryExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	right, err := i.evaluate(exprV.Right())
	if err != nil {
		return nil, err
	}

	switch exprV.Operator().Type() {
	case api.MINUS:
		return -right.(float64), nil
	case api.BANG:
		return !i.isTruthy(right), nil
	}

	// Unreachable
	return nil, errors.NewRuntimeError(exprV.Operator(), "Unary expression hit unreachable code.")
}

func (i *Interpreter) VisitBinaryExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	left, errl := i.evaluate(exprV.Left())
	if errl != nil {
		return nil, errl
	}
	right, errr := i.evaluate(exprV.Right())
	if errr != nil {
		return nil, errr
	}

	switch exprV.Operator().Type() {
	case api.GREATER:
		return left.(float64) > right.(float64), nil
	case api.GREATER_EQUAL:
		return left.(float64) >= right.(float64), nil
	case api.LESS:
		return left.(float64) < right.(float64), nil
	case api.LESS_EQUAL:
		return left.(float64) <= right.(float64), nil
	case api.BANG_EQUAL:
		return !i.isEqual(left, right), nil
	case api.EQUAL_EQUAL:
		return i.isEqual(left, right), nil
	case api.MINUS:
		err := i.checkNumberOperand(exprV.Operator(), right)
		if err != nil {
			return nil, err
		}
		return left.(float64) - right.(float64), nil
	case api.PLUS:
		// Numbers(floats or ints) or Strings
		// If one is a string then both must be strings
		lsv, isStrL := left.(api.IStringLiteral)
		rsv, isStrR := right.(api.IStringLiteral)
		if (isStrL && !isStrR) || (!isStrL && isStrR) {
			return nil, errors.NewRuntimeError(exprV.Operator(), "Both '+' operands must strings.")
		} else if isStrL && isStrR {
			nl := literals.NewStringLiteral(lsv.StringValue() + rsv.StringValue())
			return nl, nil
		}

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

			nl := literals.NewNumberLiteralVal(l + r)
			return nl, nil
		}

		liv, liok := left.(api.IIntegerLiteral)
		riv, riok := right.(api.IIntegerLiteral)
		if liok && riok {
			nl := literals.NewIntegerLiteralVal(liv.IntValue() + riv.IntValue())
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'+' Operands must be two numbers or two strings.")
	case api.SLASH:
		// Both operands will be converted to floats before division
		lfv, isNumL := left.(api.INumberLiteral)
		rfv, isNumR := right.(api.INumberLiteral)

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

		nl := literals.NewNumberLiteralVal(l / r)
		return nl, nil
	case api.STAR:
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

			nl := literals.NewNumberLiteralVal(l * r)
			return nl, nil
		}

		liv, liok := left.(api.IIntegerLiteral)
		riv, riok := right.(api.IIntegerLiteral)
		if liok && riok {
			nl := literals.NewIntegerLiteralVal(liv.IntValue() * riv.IntValue())
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "At least one operand must be a Number or both Integers.")
	}

	// Unreachable
	return nil, errors.NewRuntimeError(exprV.Operator(), "Binary expression hit unreachable code.")
}

func (i *Interpreter) isEqual(objA, objB interface{}) bool {
	if objA == nil && objB == nil {
		return true
	}
	if objA == nil {
		return false
	}

	return objA == objB
}

func (i *Interpreter) isTruthy(obj interface{}) bool {
	if obj == nil {
		return false
	}

	if value, ok := obj.(bool); ok {
		return value
	}

	return true
}

func (i *Interpreter) evaluate(expr api.IExpression) (obj interface{}, err api.IRuntimeError) {
	return expr.Accept(i)
}

func (i *Interpreter) checkNumberOperand(token api.IToken, obj interface{}) (err api.IRuntimeError) {
	_, okF := obj.(float64)
	_, okI := obj.(int)

	if okF || okI {
		return nil
	}

	return errors.NewRuntimeError(token, "Operand must be a number.")
}

func (i *Interpreter) checkNumberOperands(token api.IToken, left, right interface{}) (err api.IRuntimeError) {
	_, okL := left.(float64)
	_, okR := right.(float64)

	if okL && okR {
		return nil
	}

	_, okL = left.(int)
	_, okR = right.(int)

	if okL && okR {
		return nil
	}

	return errors.NewRuntimeError(token, "Operands must be a numbers.")
}

// nmf := reflect.TypeOf(new(api.INumberLiteral)).Elem()
// hxf := reflect.TypeOf(new(api.IHexNumberLiteral)).Elem()

// t := reflect.TypeOf(left)
// if t.Implements(nmf) {
// 	fmt.Println(t)
// }
// if t.Implements(hxf) {
// 	fmt.Println(t)
// }
