package main

import (
	"log"

	"github.com/wdevore/RISCV-Meta-Assembler/src"
)

func main() {
	assembler, err := src.NewAssembler()

	if err != nil {
		log.Fatalln(err)
	}

	err = assembler.Configure(".")
	if err != nil {
		log.Fatalln(err)
	}

	props := assembler.Properties()
	log.Println("Generating output: " + props.BinaryName())

	log.Println("Assembling...")

	for _, source := range props.Files() {
		err = assembler.Run(source)
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("Assembly done.")
}
