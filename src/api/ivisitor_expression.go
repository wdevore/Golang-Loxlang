package api

type IVisitorExpression interface {
	VisitBinaryExpression(IExpression) (obj interface{}, err IRuntimeError)
	VisitGroupingExpression(IExpression) (obj interface{}, err IRuntimeError)
	VisitLiteralExpression(IExpression) (obj interface{}, err IRuntimeError)
	VisitUnaryExpression(IExpression) (obj interface{}, err IRuntimeError)
	VisitVariableExpression(IExpression) (obj interface{}, err IRuntimeError)
}
