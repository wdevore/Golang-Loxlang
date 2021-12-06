package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/ast"
	"github.com/wdevore/RISCV-Meta-Assembler/src/parser"
	"github.com/wdevore/RISCV-Meta-Assembler/src/scanner"
)

type Assembler struct {
	// Configuration and runtime settings
	properties    api.IProperties
	configRelPath string
	errorOccurred bool

	report api.IReporter

	expression api.IExpression
}

// NewAssembler creates a new assembler for compiling assembly code
func NewAssembler() (assembler api.IAssembler, err error) {
	ass := new(Assembler)
	ass.report = NewReport()
	return ass, nil
}

func (a *Assembler) Configure(configRelPath string) error {
	props, err := a.loadProperties(configRelPath)

	if err != nil {
		log.Fatalln(err)
	}

	a.properties = props

	return nil
}

func (a *Assembler) ConfigRelPath() string {
	return a.configRelPath
}

func (a *Assembler) Properties() api.IProperties {
	return a.properties
}

func (a *Assembler) ErrorOccurred() bool {
	return a.errorOccurred
}

func (a *Assembler) SetError(occurred bool) {
	a.errorOccurred = occurred
}

func (a *Assembler) ReportLine(line int, message string) {
	a.report.ReportLine(line, message)
	a.SetError(true)
}

func (a *Assembler) ReportWhere(line int, where, message string) {
	a.report.ReportLine(line, message)
	a.SetError(true)
}

func (a *Assembler) ReportToken(token api.IToken, message string) {
	if token.Type() == api.EOF {
		a.report.ReportWhere(token.Line(), " at end", message)
	} else {
		a.report.ReportWhere(token.Line(), " at '"+token.Lexeme()+"'", message)
	}
}

func (a *Assembler) Run(source string) error {
	scanner := scanner.NewScanner(a)

	tokens, err := scanner.Scan(source)
	if err != nil {
		return fmt.Errorf("unexpected error occurred during scan: %v", err)
	}

	parser := parser.NewParser(a, tokens)

	a.expression, err = parser.Parse()
	if err != nil {
		return fmt.Errorf("unexpected error occurred during parser: %v", err)
	}

	return nil
}

func (a *Assembler) Print() {
	astPrinter := ast.NewAstPrinter().(*ast.AstPrinter)
	pretty := astPrinter.Print(a.expression)
	fmt.Println(pretty)
}

func (a *Assembler) loadProperties(configRelPath string) (properties api.IProperties, err error) {
	dataPath, err := filepath.Abs(configRelPath)
	if err != nil {
		return nil, err
	}

	path := dataPath + "/config.json"

	fmt.Println("Using '" + path + "' file")
	eConfFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer eConfFile.Close()

	bytes, err := ioutil.ReadAll(eConfFile)
	if err != nil {
		return nil, err
	}

	properties = &Properties{}
	err = json.Unmarshal(bytes, properties)

	if err != nil {
		return nil, err
	}

	return properties, nil
}
