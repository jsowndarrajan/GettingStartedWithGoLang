package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//defer function executes before the function ends
func main() {
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
}

func demo1() {
	fmt.Println("----demo 1----")
	fmt.Println("First")
	fmt.Println("Middle")
	fmt.Println("Last")
}

func demo2() {
	fmt.Println("----demo 2----")
	fmt.Println("First")
	defer fmt.Println("Middle")
	fmt.Println("Last")
}

//defer keyword executes last in first out order
func demo3() {
	fmt.Println("----demo 3----")
	defer fmt.Println("First")
	defer fmt.Println("Middle")
	defer fmt.Println("Last")
}

func demo4() {
	fmt.Println("----demo 4----")
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

func demo5() {
	fmt.Println("----demo 5----")
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // open and close the resource next to each other

	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)
}
