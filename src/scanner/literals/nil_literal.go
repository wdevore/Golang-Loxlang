package literals

import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

type NilLiteral struct {
	value string
}

func NewNilLiteral() api.INilLiteral {
	s := new(NilLiteral)
	s.value = "nil"
	return s
}

func (s NilLiteral) String() string {
	return s.value
}

func (s *NilLiteral) Value() interface{} {
	return s.value
}