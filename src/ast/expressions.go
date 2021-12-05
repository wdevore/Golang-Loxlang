package ast

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

func (e *BinaryExpression) Accept(visitor api.IVisitor) interface{} {
	return visitor.VisitBinaryExpression(e)
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

func (e *GroupingExpression) Accept(visitor api.IVisitor) interface{} {
	return visitor.VisitGroupingExpression(e)
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

func (e *LiteralExpression) Accept(visitor api.IVisitor) interface{} {
	return visitor.VisitLiteralExpression(e)
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

func (e *UnaryExpression) Accept(visitor api.IVisitor) interface{} {
	return visitor.VisitUnaryExpression(e)
}

/*
abstract class Expr {
	// defineVisitor
	interface Visitor<R> {
		R visitBinaryExpr(Binary expr);
		R visitGroupingExpr(Grouping expr);
		// ...
	}

	// loop defineType
	static class Binary extends Expr {
		Binary(left, op, right) {
			this.left = left;
			this.op = op;
			this.right = right;
		}
		@override
		<R> R accept(Visitor<R> visitor) {
			return visitor.visitBinaryExpr(this);
		}

		// Field
		final left
		final op
		final right
	}

	static class Grouping extends Expr {
		Grouping(expr) {
			this.expr = expr
		}
		@override
		<R> R accept(Visitor<R> visitor) {
			return visitor.visitGroupingExpr(this);
		}

		// Field
		final expr
	}
	//...

	abstract <R> R accept(Visitor<R> visitor);
}
*/
