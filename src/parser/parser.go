package parser

import (
	"errors"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/ast"
	"github.com/wdevore/RISCV-Meta-Assembler/src/scanner/literals"
)

type Parser struct {
	assembler api.IAssembler
	tokens    []api.IToken
	current   int
}

func NewParser(assembler api.IAssembler, tokens []api.IToken) *Parser {
	o := new(Parser)
	o.tokens = tokens
	o.assembler = assembler
	return o
}

func (p *Parser) expression() api.IExpression {
	return p.equality()
}

// --------------------------------------------------------
// equality → comparison ( ( "!=" | "==" ) comparison )*
// --------------------------------------------------------
func (p *Parser) equality() api.IExpression {
	expr := p.comparison()

	for p.match(api.BANG_EQUAL, api.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = ast.NewBinaryExpression(expr, operator, right)
	}

	return expr
}

// This checks to see if the current token has any of the given types.
// If so, it consumes the token and returns true.
// Otherwise, it returns false and leavesthe current token alone
func (p *Parser) match(types ...api.TokenType) bool {
	for _, ttype := range types {
		if p.check(ttype) {
			p.advance()
			return true
		}
	}

	return false
}

// returns true if the current token is of the given type
func (p *Parser) check(ttype api.TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type() == ttype
}

// consumes the current token and returns it, similar to
// how our scanner’s corresponding method crawled through characters
func (p *Parser) advance() api.IToken {
	if !p.isAtEnd() {
		p.current++
	}

	return p.previous()
}

// checks if we’ve run out of tokens to parse
func (p *Parser) isAtEnd() bool {
	return p.peek().Type() == api.EOF
}

// returns the current token we have yet to consume
func (p *Parser) peek() api.IToken {
	return p.tokens[p.current]
}

// returns the most recently consumed token
func (p *Parser) previous() api.IToken {
	return p.tokens[p.current-1]
}

// --------------------------------------------------------
// comparison → term ( ( ">" | ">=" | "<" | "<=" ) term )*
// --------------------------------------------------------
func (p *Parser) comparison() api.IExpression {
	expr := p.term()

	for p.match(api.GREATER, api.GREATER_EQUAL, api.LESS, api.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = ast.NewBinaryExpression(expr, operator, right)
	}

	return expr
}

// --------------------------------------------------------
// term → factor ( ( "-" | "+" ) factor )*
// --------------------------------------------------------
func (p *Parser) term() api.IExpression {
	expr := p.factor()

	for p.match(api.MINUS, api.PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = ast.NewBinaryExpression(expr, operator, right)
	}

	return expr
}

// --------------------------------------------------------
// factor → unary ( ( "/" | "*" ) unary )*
// --------------------------------------------------------
func (p *Parser) factor() api.IExpression {
	expr := p.unary()

	for p.match(api.SLASH, api.STAR) {
		operator := p.previous()
		right := p.unary()
		expr = ast.NewBinaryExpression(expr, operator, right)
	}

	return expr
}

// --------------------------------------------------------
// unary → ( "!" | "-" ) unary | primary ;
// --------------------------------------------------------
func (p *Parser) unary() api.IExpression {

	if p.match(api.BANG, api.MINUS) {
		operator := p.previous()
		right := p.unary()
		return ast.NewUnaryExpression(operator, right)
	}

	return p.primary()
}

// --------------------------------------------------------
// primary → NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")"
// --------------------------------------------------------
func (p *Parser) primary() api.IExpression {

	if p.match(api.FALSE) {
		return ast.NewLiteralExpression(literals.NewBooleanLiteral(false))
	}
	if p.match(api.TRUE) {
		return ast.NewLiteralExpression(literals.NewBooleanLiteral(true))
	}
	if p.match(api.NIL) {
		return ast.NewLiteralExpression(literals.NewNilLiteral())
	}

	if p.match(api.NUMBER, api.STRING) {
		// NOTE: may need to copy the literal!!!!
		return ast.NewLiteralExpression(p.previous().Literal())
	}

	if p.match(api.LEFT_PAREN) {
		expr := p.expression()
		p.consume(api.RIGHT_PAREN, "Expect ')' after expression")
		return ast.NewGroupingExpression(expr)
	}

	return nil
}

func (p *Parser) consume(ttype api.TokenType, message string) (token api.IToken, err error) {
	if p.check(ttype) {
		return p.advance(), nil
	}

	token = p.peek()
	return token, p.lerror(token, message)
}

func (p *Parser) lerror(ttype api.IToken, message string) error {
	p.assembler.ReportToken(ttype, message)
	return errors.New(message)
}
