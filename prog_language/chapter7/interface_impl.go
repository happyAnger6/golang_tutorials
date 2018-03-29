package main

import "fmt"

/**
接口实现
*/

type IntSet struct {

}

func (i *IntSet) String() string {
	return "IntSet."
}

func main() {
	i := IntSet{}
	i.String()
	//IntSet{}.String()

	//var _ fmt.Stringer = i
	var _ fmt.Stringer = &i
}

