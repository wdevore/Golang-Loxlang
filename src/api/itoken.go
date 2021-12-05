package api

type IToken interface {
	String() string
	Lexeme() string
	Type() TokenType
	Literal() ILiteral
	Line() int
}
