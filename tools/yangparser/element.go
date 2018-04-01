package yangparser

import (
	"./parser"
)

type Element interface {
	Parse(scan *parser.Scanner) (interface{}, error)
}
