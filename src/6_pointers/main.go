package main

import "fmt"

type gopher struct {
	name    string
	age     int
	isAdult bool
}

func main() {
	gopher1 := gopher{name: "abc", age: 21}
	fmt.Println(gopher1)

	validateAge(gopher1)
	fmt.Println(gopher1)

	validateAgeByPassingPointers(&gopher1)
	fmt.Println(gopher1)
}

func validateAge(g gopher) {
	if g.age >= 18 {
		g.isAdult = true
	}
}

func validateAgeByPassingPointers(g *gopher) {
	if g.age >= 18 {
		g.isAdult = true
	}
}
