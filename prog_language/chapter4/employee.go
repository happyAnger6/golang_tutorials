package main

import (
	"time"
	"fmt"
)

type Employee struct {
	ID int
	Name string
	Address string
	BoB time.Time
	Position string
	Salary int
	ManagerID int
}

func main() {
	e := Employee{ID:1, ManagerID:2}
	f := &e
	fmt.Println(e)
	fmt.Printf("e=%v\n", e)
	fmt.Printf("type:%T %d\n",f, f.ID)
}
