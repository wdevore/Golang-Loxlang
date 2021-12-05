package main

import (
	"fmt"
	"log"

	"github.com/wdevore/RISCV-Meta-Assembler/src"
	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/ast"
	"github.com/wdevore/RISCV-Meta-Assembler/src/scanner"
	"github.com/wdevore/RISCV-Meta-Assembler/src/scanner/literals"
)

func main() {
	test_expression()
}

func run_assembler() {
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

func test_expression() {
	// (* (- 123) (group 45.67))
	// (* (- 123) (group 45.669998))

	expression := ast.NewBinaryExpression(
		ast.NewUnaryExpression(
			scanner.NewToken(api.MINUS, "-", nil, 1),
			ast.NewLiteralExpression(
				literals.NewIntegerLiteral("123"),
			),
		),
		scanner.NewToken(api.STAR, "*", nil, 1),
		ast.NewGroupingExpression(
			ast.NewLiteralExpression(
				literals.NewNumberLiteral("45.67"),
			),
		),
	)

	astPrinter := ast.NewAstPrinter().(*ast.AstPrinter)
	pretty := astPrinter.Print(expression)
	fmt.Println(pretty)
}
