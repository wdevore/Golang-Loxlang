package api

type IExpression interface {
	Accept(IVisitorExpression) (obj interface{}, err IRuntimeError)

	// Literals
	Value() interface{}

	// Unary,Binary
	Left() IExpression
	Operator() IToken
	Right() IExpression

	// Grouping
	Expression() IExpression

	// Var
	Name() IToken
}
