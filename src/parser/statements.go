package parser

import (
	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/interpreter"
	"github.com/wdevore/RISCV-Meta-Assembler/src/scanner/literals"
	"github.com/wdevore/RISCV-Meta-Assembler/src/statements"
)

// --------------------------------------------------------
// Print statement
// --------------------------------------------------------
func (p *Parser) printStatement() (statement api.IStatement, err error) {
	value, err := p.expression()

	if err != nil {
		return nil, err
	}

	_, err = p.consume(api.SEMICOLON, "Expect ';' after value.")

	if err != nil {
		return nil, err
	}

	return statements.NewPrintStatement(value), nil
}

// --------------------------------------------------------
// Expression statement
// --------------------------------------------------------
func (p *Parser) expressionStatement() (statement api.IStatement, err error) {
	expr, err := p.expression()

	if err != nil {
		return nil, err
	}

	_, err = p.consume(api.SEMICOLON, "Expect ';' after expression.")

	if err != nil {
		return nil, err
	}

	return statements.NewExpressionStatement(expr), nil
}

// --------------------------------------------------------
// "if" statement
// --------------------------------------------------------
func (p *Parser) ifStatement() (statement api.IStatement, err error) {
	_, err = p.consume(api.LEFT_PAREN, "Expect '(' after 'if'.")
	if err != nil {
		return nil, err
	}

	condition, err := p.expression()
	if err != nil {
		return nil, err
	}

	_, err = p.consume(api.RIGHT_PAREN, "Expect ')' after 'if' condition.")
	if err != nil {
		return nil, err
	}

	thenBranch, err := p.statement()
	if err != nil {
		return nil, err
	}

	var elseBranch api.IStatement
	if p.match(api.ELSE) {
		elseBranch, err = p.statement()
		if err != nil {
			return nil, err
		}
	}

	return statements.NewIfStatement(condition, thenBranch, elseBranch), nil
}

// --------------------------------------------------------
// "while" statement
// --------------------------------------------------------
func (p *Parser) whileStatement() (expr api.IStatement, err error) {
	_, err = p.consume(api.LEFT_PAREN, "Expect '(' after 'while'.")
	if err != nil {
		return nil, err
	}

	condition, err := p.expression()
	if err != nil {
		return nil, err
	}

	_, err = p.consume(api.RIGHT_PAREN, "Expect ')' after 'while' condition.")
	if err != nil {
		return nil, err
	}

	body, err := p.statement()
	if err != nil {
		return nil, err
	}

	return statements.NewWhileStatement(condition, body), nil
}

// --------------------------------------------------------
// "for" statement via desugaring
// --------------------------------------------------------
func (p *Parser) forStatement() (expr api.IStatement, err error) {
	_, err = p.consume(api.LEFT_PAREN, "Expect '(' after 'for'.")
	if err != nil {
		return nil, err
	}

	var initializer api.IStatement

	if p.match(api.SEMICOLON) {
		// no initializer
	} else if p.match(api.VAR) {
		initializer, err = p.varDeclaration()
		if err != nil {
			return nil, err
		}
	} else {
		initializer, err = p.expressionStatement()
		if err != nil {
			return nil, err
		}
	}

	var condition api.IExpression

	if !p.check(api.SEMICOLON) {
		condition, err = p.expression()
		if err != nil {
			return nil, err
		}
	}

	_, err = p.consume(api.SEMICOLON, "Expect ';' after loop condition.")
	if err != nil {
		return nil, err
	}

	var increment api.IExpression

	if !p.check(api.RIGHT_PAREN) {
		increment, err = p.expression()
		if err != nil {
			return nil, err
		}
	}

	_, err = p.consume(api.RIGHT_PAREN, "Expect ';' after 'for' clauses.")
	if err != nil {
		return nil, err
	}

	body, err := p.statement()
	if err != nil {
		return nil, err
	}

	// The increment, if there is one, executes after the body in each iteration of the
	// loop
	if increment != nil {
		stmts := []api.IStatement{body, statements.NewExpressionStatement(increment)}
		body = statements.NewBlockStatement(stmts)
	}

	// If the condition is omitted, we jam in true to make an infinite
	// loop
	if condition == nil {
		condition = interpreter.NewLiteralExpression(literals.NewBooleanLiteral(true))
	}

	body = statements.NewWhileStatement(condition, body)

	// Finally, if there is an initializer, it runs once before the entire loop. We do that
	// by, again, replacing the whole statement with a block that runs the initializer
	// and then executes the loop.
	if initializer != nil {
		stmts := []api.IStatement{initializer, body}
		body = statements.NewBlockStatement(stmts)
	}

	return body, nil
}
