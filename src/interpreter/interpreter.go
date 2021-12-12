package interpreter

import (
	"fmt"

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

// func (i *Interpreter) Interpret(expression api.IExpression) api.IRuntimeError {
// 	value, err := i.evaluate(expression)
// 	if err != nil {
// 		return err
// 	}

// 	log.Println(value)

// 	return nil
// }

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
		v, err := i.extractNumber(right, exprV.Operator())
		if err == nil {
			nl := literals.NewNumberLiteralVal(-v)
			return nl, nil
		}

		iv, err := i.extractInteger(right, exprV.Operator())
		if err == nil {
			nl := literals.NewIntegerLiteralVal(-iv)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "Minus expression invalid operand.")
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
		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(l > r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(il > ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'>' Unexpected reachable code.")
	case api.GREATER_EQUAL:
		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(l >= r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(il >= ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'>=' Unexpected reachable code.")
	case api.LESS:
		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(l < r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(il < ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'<' Unexpected reachable code.")
	case api.LESS_EQUAL:
		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(l <= r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(il <= ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'<=' Unexpected reachable code.")
	case api.BANG_EQUAL:
		// if !i.isEqual(left, right) {
		// 	return literals.NewBooleanLiteral(true), nil
		// }

		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(l != r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(il != ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'!=' Unexpected reachable code.")
	case api.EQUAL_EQUAL:
		// if i.isEqual(left, right) {
		// 	return literals.NewBooleanLiteral(true), nil
		// }

		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(l == r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(il == ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'==' Unexpected reachable code.")
	case api.MINUS:
		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewNumberLiteralVal(l - r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewIntegerLiteralVal(il - ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'"+exprV.Operator().Lexeme()+"' Operands must be two numbers.")
	case api.PLUS:
		// Numbers(floats or ints) or Strings

		// If one is a string then both must be strings
		sl, sr, err := i.extractStrings(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewStringLiteral(sl + sr)
			return nl, nil
		}

		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewNumberLiteralVal(l + r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewIntegerLiteralVal(il + ir)
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

// -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ --
// IVisitorStatement implementations
// -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ --
func (i *Interpreter) VisitExpressionStatement(statement api.IStatement) (err api.IRuntimeError) {
	_, err = i.evaluate(statement.Expression())
	return err
}

func (i *Interpreter) VisitPrintStatement(statement api.IStatement) (err api.IRuntimeError) {
	value, err := i.evaluate(statement.Expression())

	fmt.Println(value)

	return err
}

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

// false and nil are falsey and everything else is truthy
func (i *Interpreter) isTruthy(obj interface{}) bool {
	if obj == nil {
		return false
	}

	v, isBoo := obj.(api.IBooleanLiteral)
	if isBoo {
		return v.BoolValue()
	}

	return true
}

func (i *Interpreter) evaluate(expr api.IExpression) (obj interface{}, err api.IRuntimeError) {
	return expr.Accept(i)
}

// func (i *Interpreter) checkNumberOperand(token api.IToken, obj interface{}) (err api.IRuntimeError) {
// 	_, okF := obj.(float64)
// 	_, okI := obj.(int)

// 	if okF || okI {
// 		return nil
// 	}

// 	return errors.NewRuntimeError(token, "Operand must be a number.")
// }

// func (i *Interpreter) checkNumberOperands(token api.IToken, left, right interface{}) (err api.IRuntimeError) {
// 	_, okL := left.(float64)
// 	_, okR := right.(float64)

// 	if okL && okR {
// 		return nil
// 	}

// 	_, okL = left.(int)
// 	_, okR = right.(int)

// 	if okL && okR {
// 		return nil
// 	}

// 	return errors.NewRuntimeError(token, "Operands must be a numbers.")
// }

// nmf := reflect.TypeOf(new(api.INumberLiteral)).Elem()
// hxf := reflect.TypeOf(new(api.IHexNumberLiteral)).Elem()

// t := reflect.TypeOf(left)
// if t.Implements(nmf) {
// 	fmt.Println(t)
// }
// if t.Implements(hxf) {
// 	fmt.Println(t)
// }
