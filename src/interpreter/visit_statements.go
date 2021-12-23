package interpreter

import (
	"fmt"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

// -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ --
// IVisitorStatement implementations
// -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ --
func (i *Interpreter) VisitExpressionStatement(statement api.IStatement) (err api.IRuntimeError) {
	// Simply decend
	_, err = i.evaluate(statement.Expression())
	return err
}

func (i *Interpreter) VisitPrintStatement(statement api.IStatement) (err api.IRuntimeError) {
	value, err := i.evaluate(statement.Expression())

	if err == nil {
		fmt.Println(value)
	}

	return err
}

func (i *Interpreter) VisitVariableStatement(statement api.IStatement) (err api.IRuntimeError) {
	var value interface{} = nil

	if statement.Initializer() != nil {
		value, err = i.evaluate(statement.Initializer())
		if err != nil {
			return err
		}
	}

	return i.environment.Define(statement.Name().Lexeme(), value)
}

func (i *Interpreter) VisitBlockStatement(statement api.IStatement) (err api.IRuntimeError) {
	childEnv := NewEnvironmentEnclosing(i.environment)
	return i.executeBlock(statement.Statements(), childEnv)
}

func (i *Interpreter) VisitIfStatement(statement api.IStatement) (err api.IRuntimeError) {
	value, err := i.evaluate(statement.Condition())
	if err != nil {
		return err
	}

	if i.isTruthy(value) {
		err = i.execute(statement.ThenBranch())
		if err != nil {
			return err
		}
	} else if statement.ElseBranch() != nil {
		err = i.execute(statement.ElseBranch())
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Interpreter) VisitWhileStatement(statement api.IStatement) (err api.IRuntimeError) {
	value, err := i.evaluate(statement.Condition())
	if err != nil {
		return err
	}

	for i.isTruthy(value) {
		err = i.execute(statement.Body())
		if err != nil {
			return err
		}

		value, err = i.evaluate(statement.Condition())
		if err != nil {
			return err
		}
	}

	return nil
}
