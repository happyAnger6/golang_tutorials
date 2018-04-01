package attributes

import "../parser"

type Description struct {
	Description string
}

func (d *Description) Parse(scan parser.Scanner) error {
	if scan.Scan() {
		des := scan.Text()
		d.Description = des
	}
	return nil
}
