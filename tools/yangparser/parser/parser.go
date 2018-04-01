package parser

import (
	"io"
	"../module"
	"../attributes"
)

func Parse(r io.Reader) *module.Module {
	m := &module.Module{}
	scan := NewScanner(r)
	for scan.Scan() {
		switch token := scan.Text(); token {
		case "module":
			m.Parse(scan)
		case "yang-version":
			y := attributes.YangVersion{}
			y.Parse(scan)
			m.YangVersion = y
		case "import":
			i := attributes.Import{}
			i.Parse(scan)
			m.Imports = append(m.Imports, i)
		default:
		}
	}

	return m
}
