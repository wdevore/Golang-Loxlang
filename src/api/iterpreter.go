package api

type IInterpreter interface {
	// Interpret(exprV IExpression) IRuntimeError
	Interpret(statements []IStatement) IRuntimeError
}
