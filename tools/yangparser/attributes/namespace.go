package attributes

import "../parser"

type Namespace string

func (n *Namespace) Parse(scan *parser.Scanner) (interface{}, error) {
	if scan.Scan() {
		*n = Namespace(scan.Text())
	}
	return n, nil
}
