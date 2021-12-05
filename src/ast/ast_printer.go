package ast

import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

type AstPrinter struct {
}

func NewAstPrinter() api.IVisitor {
	o := new(AstPrinter)
	return o
}

func (a *AstPrinter) Print(expr api.IExpression) string {
	return expr.Accept(a).(string)
}

func (a *AstPrinter) VisitBinaryExpression(exprV api.IExpression) interface{} {
	expr := exprV.(*BinaryExpression)
	return a.parenthesize(expr.operator.Lexeme(), expr.left, expr.right)
}

func (a *AstPrinter) VisitGroupingExpression(exprV api.IExpression) interface{} {
	expr := exprV.(*GroupingExpression)
	return a.parenthesize("group", expr.expression)
}

func (a *AstPrinter) VisitLiteralExpression(exprV api.IExpression) interface{} {
	expr := exprV.(*LiteralExpression)
	if expr.value == nil {
		return "nil"
	}
	return expr.value.String()
}

func (a *AstPrinter) VisitUnaryExpression(exprV api.IExpression) interface{} {
	expr := exprV.(*UnaryExpression)
	return a.parenthesize(expr.operator.Lexeme(), expr.right)
}

func (a *AstPrinter) parenthesize(name string, expr ...interface{}) string {
	builder := "(" + name

	for _, expres := range expr {
		builder += " "
		ex := expres.(api.IExpression)
		builder += ex.Accept(a).(string)
	}
	builder += ")"

	return builder
}
