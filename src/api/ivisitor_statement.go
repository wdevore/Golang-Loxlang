package api

type IVisitorStatement interface {
	VisitPrintStatement(IStatement) (err IRuntimeError)
	VisitVariableStatement(IStatement) (err IRuntimeError)
}
