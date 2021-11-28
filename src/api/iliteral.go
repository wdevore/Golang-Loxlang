package api

type ILiteral interface {
	String() string
	Value() interface{}
}

type IStringLiteral interface {
	ILiteral
}

type ICharLiteral interface {
	ILiteral
}

type INumberLiteral interface {
	ILiteral
}

type IHexNumberLiteral interface {
	ILiteral
}

type INilLiteral interface {
	ILiteral
}
