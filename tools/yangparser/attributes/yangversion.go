package attributes

import "../parser"

type YangVersion struct {
	Version string
}

func (y *YangVersion) Parse(scan *parser.Scanner) *YangVersion {
	if scan.Scan() {
		y.Version = scan.Text()
	}
	return y
}
