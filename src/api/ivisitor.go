package api

type IVisitor interface {
	VisitBinaryExpression(IExpression) (obj interface{}, err IRuntimeError)
	VisitGroupingExpression(IExpression) (obj interface{}, err IRuntimeError)
	VisitLiteralExpression(IExpression) (obj interface{}, err IRuntimeError)
	VisitUnaryExpression(IExpression) (obj interface{}, err IRuntimeError)
}

type IVisitorStatement interface {
	VisitPrintStatement(IStatement) (err IRuntimeError)
}
