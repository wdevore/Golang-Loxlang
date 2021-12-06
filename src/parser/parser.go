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

func (p *Parser) Parse() (expr api.IExpression, err error) {
	expr, err = p.expression()
	if err != nil {
		p.assembler.SetError(true)
	}
	return expr, err
}

func (p *Parser) expression() (expr api.IExpression, err error) {
	return p.equality()
}

// --------------------------------------------------------
// equality → comparison ( ( "!=" | "==" ) comparison )*
// --------------------------------------------------------
func (p *Parser) equality() (expr api.IExpression, err error) {
	expr, err = p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(api.BANG_EQUAL, api.EQUAL_EQUAL) {
		operator := p.previous()
		right, errc := p.comparison()
		if errc != nil {
			return nil, errc
		}
		expr = ast.NewBinaryExpression(expr, operator, right)
	}

	return expr, nil
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
func (p *Parser) comparison() (expr api.IExpression, err error) {
	expr, err = p.term()
	if err != nil {
		return nil, err
	}

	for p.match(api.GREATER, api.GREATER_EQUAL, api.LESS, api.LESS_EQUAL) {
		operator := p.previous()
		right, errc := p.term()
		if errc != nil {
			return nil, errc
		}
		expr = ast.NewBinaryExpression(expr, operator, right)
	}

	return expr, nil
}

// --------------------------------------------------------
// term → factor ( ( "-" | "+" ) factor )*
// --------------------------------------------------------
func (p *Parser) term() (expr api.IExpression, err error) {
	expr, err = p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(api.MINUS, api.PLUS) {
		operator := p.previous()
		right, errc := p.factor()
		if errc != nil {
			return nil, errc
		}
		expr = ast.NewBinaryExpression(expr, operator, right)
	}

	return expr, nil
}

// --------------------------------------------------------
// factor → unary ( ( "/" | "*" ) unary )*
// --------------------------------------------------------
func (p *Parser) factor() (expr api.IExpression, err error) {
	expr, err = p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(api.SLASH, api.STAR) {
		operator := p.previous()
		right, errc := p.unary()
		if errc != nil {
			return nil, errc
		}
		expr = ast.NewBinaryExpression(expr, operator, right)
	}

	return expr, nil
}

// --------------------------------------------------------
// unary → ( "!" | "-" ) unary | primary ;
// --------------------------------------------------------
func (p *Parser) unary() (expr api.IExpression, err error) {

	if p.match(api.BANG, api.MINUS) {
		operator := p.previous()
		right, errc := p.unary()
		if errc != nil {
			return nil, errc
		}
		return ast.NewUnaryExpression(operator, right), nil
	}

	return p.primary()
}

// --------------------------------------------------------
// primary → NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")"
// --------------------------------------------------------
func (p *Parser) primary() (expr api.IExpression, err error) {

	if p.match(api.FALSE) {
		return ast.NewLiteralExpression(literals.NewBooleanLiteral(false)), nil
	}
	if p.match(api.TRUE) {
		return ast.NewLiteralExpression(literals.NewBooleanLiteral(true)), nil
	}
	if p.match(api.NIL) {
		return ast.NewLiteralExpression(literals.NewNilLiteral()), nil
	}

	if p.match(api.NUMBER, api.STRING) {
		// NOTE: may need to copy the literal!!!!
		return ast.NewLiteralExpression(p.previous().Literal()), nil
	}

	if p.match(api.LEFT_PAREN) {
		expr, errc := p.expression()
		if errc != nil {
			return nil, errc
		}
		_, err = p.consume(api.RIGHT_PAREN, "Expect ')' after expression")
		if err != nil {
			return nil, err
		}
		return ast.NewGroupingExpression(expr), nil
	}

	// If none of the cases in there match,
	// it means we are sitting on a token that can’t start an expression.
	return nil, p.lerror(p.previous(), "Expected expression to begin.")
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

// It discards tokens until it thinks it found a statement boundary.
func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type() == api.RIGHT_BRACE {
			return
		}

		switch p.peek().Type() {
		case api.CONST,
			api.IMPORT,
			api.CODE,
			api.ALIGN_TO,
			api.GLOBAL,
			api.AT,
			api.AS,
			api.USE,
			api.READ_ONLY,
			api.BYTE,
			api.HALF,
			api.WORD,
			api.DATA,
			api.INT,
			api.HI,
			api.LO,
			api.ADD,
			api.SUB,
			api.XOR,
			api.OR,
			api.AND,
			api.SLL,
			api.SRL,
			api.SRA,
			api.SLT,
			api.SLTU,
			api.ADDI,
			api.XORI,
			api.ORI,
			api.ANDI,
			api.SLLI,
			api.SRLI,
			api.SRAI,
			api.SLTI,
			api.SLTIU,
			api.LB,
			api.LH,
			api.LW,
			api.LBU,
			api.LHU,
			api.SB,
			api.SH,
			api.SW,
			api.BEQ,
			api.BNE,
			api.BLT,
			api.BGE,
			api.BLTU,
			api.BGEU,
			api.JAL,
			api.JALR,
			api.LUI,
			api.AUIPC,
			api.ECALL,
			api.EBREAK,
			api.LA,
			api.NOP,
			api.LI,
			api.MV,
			api.NOT,
			api.NEG,
			api.NEGW,
			api.SEXT,
			api.SEQZ,
			api.SNEZ,
			api.SLTZ,
			api.SGTZ,
			api.BEQZ,
			api.BNEZ,
			api.BLEZ,
			api.BGEZ,
			api.BLTZ,
			api.BGTZ,
			api.BGT,
			api.BLE,
			api.BGTU,
			api.BLEU,
			api.J,
			api.RET,
			api.CALL,
			api.TAIL:
			return
		}
	}

	p.advance()
}
