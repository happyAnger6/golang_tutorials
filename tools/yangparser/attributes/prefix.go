package attributes

import "../parser"

type Prefix string;

func (p *Prefix) Parse(scan *parser.Scanner) (interface{}, error) {
	if scan.Scan() {
		*p = Prefix(scan.Text())
	}

	return p, nil
}
