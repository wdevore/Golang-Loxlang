package literals

import (
	"fmt"
	"strconv"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

type IntegerLiteral struct {
	value int64
}

func NewIntegerLiteral(value string) api.IIntegerLiteral {
	s := new(IntegerLiteral)
	s.value, _ = strconv.ParseInt(value, 10, 32)
	return s
}

func (n IntegerLiteral) String() string {
	return fmt.Sprintf("%d", n.value)
}

func (n *IntegerLiteral) Value() interface{} {
	return n.value
}
