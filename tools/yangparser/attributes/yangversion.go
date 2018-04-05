package attributes

import (
	"../parser"
	"../../yangparser"
)

type YangVersion struct {
	Version string
}

func (y *YangVersion) Parse(scan *parser.Scanner) (yangparser.Element, error) {
	if scan.Scan() {
		y.Version = scan.Text()
	}
	return y, nil
}
