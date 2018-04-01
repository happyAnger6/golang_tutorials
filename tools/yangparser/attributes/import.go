package attributes

import (
	"../parser"
)

type Import struct {
	Name string
	Prefix Prefix
}

func (i *Import) Parse(scan *parser.Scanner) (interface{}, error){
	if scan.Scan() {
		i.Name = scan.Text()
	}

	return i, nil
}
