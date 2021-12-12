package statements

import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

// ---------------------------------------------------
// Expression statement
// ---------------------------------------------------
type ExpressionStatement struct {
	expression api.IExpression
}

func NewExpressionStatement(expression api.IExpression) api.IStatement {
	o := new(ExpressionStatement)
	o.expression = expression
	return o
}

func (s *ExpressionStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
	return nil
}

func (s *ExpressionStatement) Expression() api.IExpression {
	return s.expression
}

// ---------------------------------------------------
// Print statement
// ---------------------------------------------------
type PrintStatement struct {
	expression api.IExpression
}

func NewPrintStatement(expression api.IExpression) api.IStatement {
	o := new(PrintStatement)
	o.expression = expression
	return o
}

func (s *PrintStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
	return visitor.VisitPrintStatement(s)
}

func (s *PrintStatement) Expression() api.IExpression {
	return s.expression
}

// ---------------------------------------------------
// Block statement
// ---------------------------------------------------
// type BlockStatement struct {
// 	statements []api.IExpression
// }

// func NewBlockStatement(statements []api.IExpression) api.IStatement {
// 	o := new(BlockStatement)
// 	o.statements = statements
// 	return o
// }

// func (s *BlockStatement) Accept(visitor api.IVisitorStatement) (obj interface{}, err api.IRuntimeError) {
// 	return nil, nil
// }

// func (s *BlockStatement) Expression() api.IExpression {
// 	return nil
// }
