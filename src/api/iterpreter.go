package api

type IInterpreter interface {
	Interpret(exprV IExpression) IRuntimeError
}
