package literals

import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

type CharLiteral struct {
	value string
}

func NewCharLiteral(value string) api.ICharLiteral {
	s := new(CharLiteral)
	s.value = value
	return s
}

func (s CharLiteral) String() string {
	return s.value
}

func (s *CharLiteral) Value() interface{} {
	return s.value
}
