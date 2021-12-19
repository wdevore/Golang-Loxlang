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

func (s *ExpressionStatement) Condition() api.IExpression {
	return nil
}

func (s *ExpressionStatement) ThenBranch() api.IStatement {
	return nil
}

func (s *ExpressionStatement) ElseBranch() api.IStatement {
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

func (s *BlockStatement) Condition() api.IExpression {
	return nil
}

func (s *BlockStatement) ThenBranch() api.IStatement {
	return nil
}

func (s *BlockStatement) ElseBranch() api.IStatement {
	return nil
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

func (s *PrintStatement) Condition() api.IExpression {
	return nil
}

func (s *PrintStatement) ThenBranch() api.IStatement {
	return nil
}

func (s *PrintStatement) ElseBranch() api.IStatement {
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

func (s *VarStatement) Condition() api.IExpression {
	return nil
}

func (s *VarStatement) ThenBranch() api.IStatement {
	return nil
}

func (s *VarStatement) ElseBranch() api.IStatement {
	return nil
}

// ---------------------------------------------------
// "if" statement
// ---------------------------------------------------
type IfStatement struct {
	condition  api.IExpression
	thenBranch api.IStatement
	elseBranch api.IStatement
}

func NewIfStatement(condition api.IExpression, thenBranch, elseBranch api.IStatement) api.IStatement {
	o := new(IfStatement)
	o.condition = condition
	o.thenBranch = thenBranch
	o.elseBranch = elseBranch
	return o
}

func (s *IfStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
	return visitor.VisitIfStatement(s)
}

func (s *IfStatement) Expression() api.IExpression {
	return nil
}

func (s *IfStatement) Initializer() api.IExpression {
	return nil
}

func (s *IfStatement) Name() api.IToken {
	return nil
}

func (s *IfStatement) Statements() []api.IStatement {
	return nil
}

func (s *IfStatement) Condition() api.IExpression {
	return s.condition
}

func (s *IfStatement) ThenBranch() api.IStatement {
	return s.thenBranch
}

func (s *IfStatement) ElseBranch() api.IStatement {
	return s.elseBranch
}

// ---------------------------------------------------
// Logical statement
// ---------------------------------------------------
// type LogicalStatement struct {
// 	left     api.IExpression
// 	operator api.IToken
// 	right    api.IExpression
// }

// func NewLogicalStatement(left api.IExpression, operator api.IToken, right api.IExpression) api.IStatement {
// 	o := new(LogicalStatement)
// 	o.left = left
// 	o.operator = operator
// 	o.right = right
// 	return o
// }

// func (s *LogicalStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
// 	return visitor.VisitIfStatement(s)
// }

// func (s *LogicalStatement) Expression() api.IExpression {
// 	return nil
// }

// func (s *LogicalStatement) Initializer() api.IExpression {
// 	return nil
// }

// func (s *LogicalStatement) Name() api.IToken {
// 	return nil
// }

// func (s *LogicalStatement) Statements() []api.IStatement {
// 	return nil
// }

// func (s *LogicalStatement) Left() api.IExpression {
// 	return s.left
// }

// func (s *LogicalStatement) Operator() api.IToken {
// 	return s.operator
// }

// func (s *LogicalStatement) Right() api.IExpression {
// 	return s.right
// }
