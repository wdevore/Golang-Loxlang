package api

type IRuntimeError interface {
	Token() IToken
	Message() string
	String() string
}
