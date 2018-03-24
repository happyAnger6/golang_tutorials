package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags int
const (
	FlagUp Flags = 1 << iota
	FlagBroadCast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func main() {
	fmt.Println(Sunday, Monday, Tuesday, Saturday)
	fmt.Println(FlagUp, FlagBroadCast, FlagLoopback, FlagPointToPoint, FlagMulticast)
}

