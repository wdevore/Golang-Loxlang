package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

type Assembler struct {
	// Configuration and runtime settings
	properties *Properties
}

// NewAssembler creates a new assembler for compiling assembly code
func NewAssembler() (assembler api.IAssembler, err error) {
	assembler = new(Assembler)

	return assembler, nil
}

func (a *Assembler) Configure(runPath string) error {
	dataPath, err := filepath.Abs(runPath)
	if err != nil {
		return err
	}

	path := dataPath + "/config.json"

	fmt.Println("Using '" + path + "' file")
	eConfFile, err := os.Open(path)
	if err != nil {
		return err
	}

	defer eConfFile.Close()

	bytes, err := ioutil.ReadAll(eConfFile)
	if err != nil {
		return err
	}

	a.properties = &Properties{}
	err = json.Unmarshal(bytes, a.properties)

	if err != nil {
		return err
	}

	return nil
}
