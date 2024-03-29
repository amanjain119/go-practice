package main

import (
	"fmt"
)

type Tester interface {
	Test()
}

type MyFloat float64

func (m MyFloat) Test() {
	fmt.Println(m)
}

func describe(t Tester) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func main() {
	var t Tester
	t = MyFloat(89.7)
	fmt.Println(t)
	describe(t)
	t.Test()
}
