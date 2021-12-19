package api

type IStatement interface {
	Accept(IVisitorStatement) (err IRuntimeError)

	Expression() IExpression

	// Var statement
	Name() IToken
	Initializer() IExpression

	// Blocks
	Statements() []IStatement

	// "If"
	Condition() IExpression
	ThenBranch() IStatement
	ElseBranch() IStatement
}
