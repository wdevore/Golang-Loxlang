package scanner

import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

type Token struct {
	ttype   TokenType
	lexeme  string
	literal api.ILiteral
	line    int
}

func NewToken(ttype TokenType, lexeme string, literal api.ILiteral, line int) *Token {
	t := new(Token)
	t.ttype = ttype
	t.lexeme = lexeme
	t.literal = literal
	t.line = line
	return t
}

func (t Token) String() string {
	return "Type: '" + t.ttype.String() + "' Lexeme: '" + t.lexeme + "' Literal: '" + t.literal.String() + "'"
}
