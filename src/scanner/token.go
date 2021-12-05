package scanner

import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

type Token struct {
	ttype   TokenType
	lexeme  string
	literal api.ILiteral
	line    int
}

func NewToken(ttype TokenType, lexeme string, literal api.ILiteral, line int) api.IToken {
	t := new(Token)
	t.ttype = ttype
	t.lexeme = lexeme
	t.literal = literal
	t.line = line
	return t
}

func (t *Token) Lexeme() string {
	return t.lexeme
}

func (t Token) String() string {
	return "Type: '" + t.ttype.String() + "' Lexeme: '" + t.lexeme + "' Literal: '" + t.literal.String() + "'"
}
