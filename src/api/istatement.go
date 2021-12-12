package api

type IStatement interface {
	Accept(IVisitorStatement) (err IRuntimeError)

	Expression() IExpression
}
