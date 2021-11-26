package main

import (
	"log"

	"github.com/wdevore/RISCV-Meta-Assembler/src"
)

// Main assembler entry point
func main() {
	assembler, err := src.NewAssembler()

	if err != nil {
		log.Fatalln(err)
	}

	err = assembler.Configure(".")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Assembly done.")
}
