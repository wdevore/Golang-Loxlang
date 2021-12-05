package parser

import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

type IVisitor interface {
	VisitBinaryExpression(*BinaryExpression) interface{}
	VisitGroupingExpression(*GroupingExpression) interface{}
	VisitLiteralExpression(*LiteralExpression) interface{}
	VisitUnaryExpression(*UnaryExpression) interface{}
}

type IExpression interface {
	Accept(IVisitor) interface{}
}

// ---------------------------------------------------
// Binary
// ---------------------------------------------------
type BinaryExpression struct {
	left     IExpression
	operator api.IToken
	right    IExpression
}

func NewBinaryExpression(left IExpression, operator api.IToken, right IExpression) IExpression {
	e := new(BinaryExpression)
	e.left = left
	e.operator = operator
	e.right = right
	return e
}

func (e *BinaryExpression) Accept(visitor IVisitor) interface{} {
	return visitor.VisitBinaryExpression(e)
}

// ---------------------------------------------------
// Grouping
// ---------------------------------------------------
type GroupingExpression struct {
	expression IExpression
}

func NewGroupingExpression(expression IExpression) IExpression {
	e := new(GroupingExpression)
	e.expression = expression
	return e
}

func (e *GroupingExpression) Accept(visitor IVisitor) interface{} {
	return visitor.VisitGroupingExpression(e)
}

// ---------------------------------------------------
// Literal
// ---------------------------------------------------
type LiteralExpression struct {
	value api.ILiteral
}

func NewLiteralExpression(value api.ILiteral) IExpression {
	e := new(LiteralExpression)
	e.value = value
	return e
}

func (e *LiteralExpression) Accept(visitor IVisitor) interface{} {
	return visitor.VisitLiteralExpression(e)
}

// ---------------------------------------------------
// Unary
// ---------------------------------------------------
type UnaryExpression struct {
	operator api.IToken
	right    IExpression
}

func NewUnaryExpression(operator api.IToken, right IExpression) IExpression {
	e := new(UnaryExpression)
	e.operator = operator
	e.right = right
	return e
}

func (e *UnaryExpression) Accept(visitor IVisitor) interface{} {
	return visitor.VisitUnaryExpression(e)
}
