package src

import (
	"fmt"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

type Report struct {
}

func NewReport() api.IReporter {
	o := new(Report)
	return o
}

func (r *Report) ReportLine(line int, message string) {
	fmt.Printf("[line %d] Error: %s", line, message)
}

func (r *Report) ReportWhere(line int, where, message string) {
	fmt.Printf("[line %d] Error %s : %s", line, where, message)
}
