package api

type IExpression interface {
	Accept(IVisitor) interface{}
}
