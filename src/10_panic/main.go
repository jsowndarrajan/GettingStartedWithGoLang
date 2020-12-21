package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//demo1()
	demo2()
	//demo3()
	//demo4()
}

func demo1() {
	fmt.Println("start")
	panic("something bad happened")
	fmt.Println("end")
}

func demo2() {
	fmt.Println("start")
	defer fmt.Println("this was deffered")
	panic("something bad happened")
	fmt.Println("end")
}

func demo3() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end")
}

func demo4() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
