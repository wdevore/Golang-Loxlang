package parser

type AstPrinter struct {
}

func NewAstPrinter() *AstPrinter {
	o := new(AstPrinter)
	return o
}

func (a *AstPrinter) Print(expr IExpression) string {
	return expr.Accept(a).(string)
}

func (a *AstPrinter) VisitBinaryExpression(expr *BinaryExpression) interface{} {
	return a.parenthesize(expr.operator.Lexeme(), expr.left, expr.right)
}

func (a *AstPrinter) VisitGroupingExpression(expr *GroupingExpression) interface{} {
	return a.parenthesize("group", expr.expression)
}

func (a *AstPrinter) VisitLiteralExpression(expr *LiteralExpression) interface{} {
	if expr.value == nil {
		return "nil"
	}
	return expr.value.String()
}

func (a *AstPrinter) VisitUnaryExpression(expr *UnaryExpression) interface{} {
	return a.parenthesize(expr.operator.Lexeme(), expr.right)
}

func (a *AstPrinter) parenthesize(name string, expr ...interface{}) string {
	builder := "(" + name

	for _, expres := range expr {
		builder += " "
		ex := expres.(IExpression)
		builder += ex.Accept(a).(string)
	}
	builder += ")"

	return builder
}
