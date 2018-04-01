package main

import "fmt"

type namespace string
func (n *namespace) parse() (interface{}, error) {
	*n = "hello"
	return n, nil
}

func main() {
	var n namespace
	s, _ := n.parse()
	fmt.Println(n)
	fmt.Printf("n:%v n:%[1]T\n", n)
	fmt.Printf("n:%v n:%[1]T\n", s)
	v := s.(*namespace)
	fmt.Printf("n:%v n:%[1]T\n", v)
}
