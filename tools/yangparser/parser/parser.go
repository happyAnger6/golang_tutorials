package parser

import (
	"io"
	"../module"
	"../attributes"
	"../../yangparser"
	"fmt"
)

var KeysMap map[string]yangparser.Element

func init() {
	KeysMap["yang-version"] = &attributes.YangVersion{}
}

func Parse(r io.Reader) *module.Module {
	m := &module.Module{}
	scan := NewScanner(r)
	for scan.Scan() {
		token := scan.Text();
		if e, ok := KeysMap[token]; ok {
			e.Parse(scan)
		} else {
			fmt.Printf("a unkown keyword:%v found\r\n", token)
		}
		/*switch token := scan.Text(); token {
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
		}*/
	}

	return m
}
