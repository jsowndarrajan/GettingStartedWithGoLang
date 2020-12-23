package main

import "fmt"

type gopher struct {
	name    string
	age     int
	isAdult bool
}

func main() {
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
}

func demo1() {
	fmt.Println("----demo1----")
	a := 42
	b := a //copy the value to b
	fmt.Println(a, b)
	fmt.Println(&a, &b) //print address
	c := &a
	fmt.Println(c, *c) //dereferencing operator
}

func demo2() {
	fmt.Println("----demo2----")
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	d := &a[2]
	fmt.Println(a, b, c, d)
	*c = 99
	fmt.Println(a, b, c, d)
	c = d
	*c = -1
	fmt.Println(a, b, c, d)
}

func demo3() {
	fmt.Println("----demo3----")
	var nilGopher *gopher
	fmt.Println(nilGopher)

	var defaultGopher *gopher
	defaultGopher = new(gopher)
	fmt.Println(defaultGopher)
}

func demo4() {
	fmt.Println("----demo4----")

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

func demo5() {
	fmt.Println("----demo5----")
	a := [3]int{1, 2, 3} //array
	b := a
	a[1] = 99
	fmt.Println(a, b)

	c := []int{1, 2, 3} //slice
	d := c
	c[1] = 99
	fmt.Println(c, d)

	e := map[string]int{"key1": 1, "key2": 2} //dictionary
	f := e
	e["key1"] = 99
	fmt.Println(e, f)
}
