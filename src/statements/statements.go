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
	return visitor.VisitExpressionStatement(s)
}

func (s *ExpressionStatement) Expression() api.IExpression {
	return s.expression
}

func (s *ExpressionStatement) Initializer() api.IExpression {
	return nil
}

func (s *ExpressionStatement) Name() api.IToken {
	return nil
}

func (s *ExpressionStatement) Statements() []api.IStatement {
	return nil
}

// ---------------------------------------------------
// Block statement
// ---------------------------------------------------
type BlockStatement struct {
	statements []api.IStatement
}

func NewBlockStatement(statements []api.IStatement) api.IStatement {
	o := new(BlockStatement)
	o.statements = statements
	return o
}

func (s *BlockStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
	return visitor.VisitBlockStatement(s)
}

func (s *BlockStatement) Expression() api.IExpression {
	return nil
}

func (s *BlockStatement) Initializer() api.IExpression {
	return nil
}

func (s *BlockStatement) Name() api.IToken {
	return nil
}

func (s *BlockStatement) Statements() []api.IStatement {
	return s.statements
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

func (s *PrintStatement) Initializer() api.IExpression {
	return nil
}

func (s *PrintStatement) Name() api.IToken {
	return nil
}

func (s *PrintStatement) Statements() []api.IStatement {
	return nil
}

// ---------------------------------------------------
// var statement
// ---------------------------------------------------
type VarStatement struct {
	name        api.IToken
	initializer api.IExpression
}

func NewVarStatement(name api.IToken, initializer api.IExpression) api.IStatement {
	o := new(VarStatement)
	o.name = name
	o.initializer = initializer
	return o
}

func (s *VarStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
	return visitor.VisitVariableStatement(s)
}

func (s *VarStatement) Expression() api.IExpression {
	return nil
}

func (s *VarStatement) Initializer() api.IExpression {
	return s.initializer
}

func (s *VarStatement) Name() api.IToken {
	return s.name
}

func (s *VarStatement) Statements() []api.IStatement {
	return nil
}
