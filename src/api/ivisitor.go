package api

type IVisitor interface {
	VisitBinaryExpression(IExpression) interface{}
	VisitGroupingExpression(IExpression) interface{}
	VisitLiteralExpression(IExpression) interface{}
	VisitUnaryExpression(IExpression) interface{}
}
