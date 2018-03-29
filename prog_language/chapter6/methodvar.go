package main

import (
	"fmt"
	"time"
)

type Rocket struct {

}

func (r *Rocket) Launch() {
	fmt.Println("Rocket launch.")
}

func main() {
	r := new(Rocket)
	time.AfterFunc(time.Second, r.Launch)
	time.Sleep(time.Second * 2)
}
