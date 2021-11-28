package literals

import (
	"fmt"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

type BinaryNumberLiteral struct {
	value string
}

func NewBinaryNumberLiteral(value string) api.INumberLiteral {
	s := new(BinaryNumberLiteral)
	s.value = value
	return s
}

func (n BinaryNumberLiteral) String() string {
	return fmt.Sprintf("0b%s", n.value)
}

func (n *BinaryNumberLiteral) Value() interface{} {
	return n.value
}
