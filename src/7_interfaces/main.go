package main

import "fmt"

func main() {
	for _, element := range getList() {
		fmt.Println(element.jump())
	}
}

type jumper interface {
	jump() string
}

type gopher struct {
	name string
	age  int
}

func (g *gopher) jump() string {
	if g.age <= 35 {
		return g.name + " can jump high"
	}
	return g.name + " still can jump"
}

type horse struct {
	name   string
	weight int
}

func (h *horse) jump() string {
	if h.weight <= 1000 {
		return h.name + " can jump too"
	}
	return h.name + " cannot jump"
}

func getList() []jumper {
	gopher1 := &gopher{name: "abc", age: 22}
	gopher2 := &gopher{name: "xyz", age: 32}
	horse1 := &horse{name: "jocky", weight: 320}
	list := []jumper{gopher1, gopher2, horse1}
	return list
}
