package literals

import (
	"fmt"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

type BooleanLiteral struct {
	value bool
}

func NewBooleanLiteral(value bool) api.IIntegerLiteral {
	s := new(BooleanLiteral)
	s.value = value
	return s
}

func (n BooleanLiteral) String() string {
	return fmt.Sprintf("%v", n.value)
}

func (n *BooleanLiteral) Value() interface{} {
	return n.value
}
