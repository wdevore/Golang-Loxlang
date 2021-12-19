package api

type IVisitorStatement interface {
	VisitExpressionStatement(IStatement) (err IRuntimeError)
	VisitPrintStatement(IStatement) (err IRuntimeError)
	VisitVariableStatement(IStatement) (err IRuntimeError)
	VisitBlockStatement(IStatement) (err IRuntimeError)
	VisitIfStatement(IStatement) (err IRuntimeError)
}
