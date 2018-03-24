package main

import (
	"bytes"
	"fmt"
)

func int2string(value []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range value {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	s := int2string([]int{1,2,3})
	fmt.Println(len(s))
}
