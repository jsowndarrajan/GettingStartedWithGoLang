package main

import "fmt"

type gopher struct {
	name string
	age  int
}

func main() {
	gopher1 := gopher{name: "abc", age: 30}
	var gopher2 gopher
	gopher2 = gopher{name: "xyz", age: 95}

	fmt.Println(gopher1.jump())
	fmt.Println(gopher2.jump())
}

func (g gopher) jump() string {
	if g.age < 35 {
		return g.name + " can jump high"
	}
	return g.name + " still can jump"
}
