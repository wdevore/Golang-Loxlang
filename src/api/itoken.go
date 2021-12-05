package api

type IToken interface {
	String() string
	Lexeme() string
}
