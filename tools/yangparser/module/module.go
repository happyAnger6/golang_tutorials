package module

import (
	"../attributes"
	"../parser"
)

type Module struct {
	Name string
	YangVersion attributes.YangVersion `json:"yang-version"`
	Namespace attributes.Namespace `json:"namespace"`
	Imports []attributes.Import `json:"import"`
}

func (m *Module) Parse(scan *parser.Scanner) (interface{}, error) {
	if scan.Scan() {
		m.Name = scan.Text()
	}

	return m, nil
}