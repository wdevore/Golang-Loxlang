package interpreter

import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

// ---------------------------------------------------
// Binary
// ---------------------------------------------------
type BinaryExpression struct {
	left     api.IExpression
	operator api.IToken
	right    api.IExpression
}

func NewBinaryExpression(left api.IExpression, operator api.IToken, right api.IExpression) api.IExpression {
	e := new(BinaryExpression)
	e.left = left
	e.operator = operator
	e.right = right
	return e
}

func (e *BinaryExpression) Accept(visitor api.IVisitorExpression) (obj interface{}, err api.IRuntimeError) {
	return visitor.VisitBinaryExpression(e)
}

func (e *BinaryExpression) Value() interface{} {
	return nil
}

func (e *BinaryExpression) Left() api.IExpression {
	return e.left
}

func (e *BinaryExpression) Operator() api.IToken {
	return e.operator
}

func (e *BinaryExpression) Right() api.IExpression {
	return e.right
}

func (e *BinaryExpression) Expression() api.IExpression {
	return nil
}

func (e *BinaryExpression) Name() api.IToken {
	return nil
}

// ---------------------------------------------------
// Grouping
// ---------------------------------------------------
type GroupingExpression struct {
	expression api.IExpression
}

func NewGroupingExpression(expression api.IExpression) api.IExpression {
	e := new(GroupingExpression)
	e.expression = expression
	return e
}

func (e *GroupingExpression) Accept(visitor api.IVisitorExpression) (obj interface{}, err api.IRuntimeError) {
	return visitor.VisitGroupingExpression(e)
}

func (e *GroupingExpression) Value() interface{} {
	return nil
}

func (e *GroupingExpression) Left() api.IExpression {
	return nil
}

func (e *GroupingExpression) Operator() api.IToken {
	return nil
}

func (e *GroupingExpression) Right() api.IExpression {
	return nil
}

func (e *GroupingExpression) Expression() api.IExpression {
	return e.expression
}

func (e *GroupingExpression) Name() api.IToken {
	return nil
}

// ---------------------------------------------------
// Literal
// ---------------------------------------------------
type LiteralExpression struct {
	value api.ILiteral
}

func NewLiteralExpression(value api.ILiteral) api.IExpression {
	e := new(LiteralExpression)
	e.value = value
	return e
}

func (e *LiteralExpression) Accept(visitor api.IVisitorExpression) (obj interface{}, err api.IRuntimeError) {
	return visitor.VisitLiteralExpression(e)
}

func (e *LiteralExpression) Value() interface{} {
	return e.value
}

func (e *LiteralExpression) Left() api.IExpression {
	return nil
}

func (e *LiteralExpression) Operator() api.IToken {
	return nil
}

func (e *LiteralExpression) Right() api.IExpression {
	return nil
}

func (e *LiteralExpression) Expression() api.IExpression {
	return nil
}

func (e *LiteralExpression) Name() api.IToken {
	return nil
}

// ---------------------------------------------------
// Unary
// ---------------------------------------------------
type UnaryExpression struct {
	operator api.IToken
	right    api.IExpression
}

func NewUnaryExpression(operator api.IToken, right api.IExpression) api.IExpression {
	e := new(UnaryExpression)
	e.operator = operator
	e.right = right
	return e
}

func (e *UnaryExpression) Accept(visitor api.IVisitorExpression) (obj interface{}, err api.IRuntimeError) {
	return visitor.VisitUnaryExpression(e)
}

func (e *UnaryExpression) Value() interface{} {
	return nil
}

func (e *UnaryExpression) Left() api.IExpression {
	return nil
}

func (e *UnaryExpression) Operator() api.IToken {
	return e.operator
}

func (e *UnaryExpression) Right() api.IExpression {
	return e.right
}

func (e *UnaryExpression) Expression() api.IExpression {
	return nil
}

func (e *UnaryExpression) Name() api.IToken {
	return nil
}

// ---------------------------------------------------
// Variable
// ---------------------------------------------------
type VariableExpression struct {
	name api.IToken
}

func NewVariableExpression(name api.IToken) api.IExpression {
	e := new(VariableExpression)
	e.name = name
	return e
}

func (e *VariableExpression) Accept(visitor api.IVisitorExpression) (obj interface{}, err api.IRuntimeError) {
	return visitor.VisitVariableExpression(e)
}

func (e *VariableExpression) Value() interface{} {
	return nil
}

func (e *VariableExpression) Left() api.IExpression {
	return nil
}

func (e *VariableExpression) Operator() api.IToken {
	return nil
}

func (e *VariableExpression) Right() api.IExpression {
	return nil
}

func (e *VariableExpression) Expression() api.IExpression {
	return nil
}

func (e *VariableExpression) Name() api.IToken {
	return e.name
}
