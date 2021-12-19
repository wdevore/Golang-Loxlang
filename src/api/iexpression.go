package api

type ExpressionType int64

const (
	// since iota starts with 0, the first value
	// defined here will be the default
	UNDEFINED_EXPR ExpressionType = iota
	BINARY_EXPR
	GROUPING_EXPR
	LITERAL_EXPR
	UNARY_EXPR
	VAR_EXPR
	ASSIGN_EXPR
	LOGIC_EXPR
)

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

	// What type of expression
	Type() ExpressionType
}
