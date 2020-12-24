package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("# threads available:", runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(3)
	fmt.Println("# threads available:", runtime.GOMAXPROCS(-1))
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
	demo6()
}

func demo1() {
	fmt.Println("-----demo1-----")
	msg := "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "goodbye"
	time.Sleep(time.Second)
}

func demo2() {
	fmt.Println("-----demo2 : Pass variable to goroutine to avoid race condition-----")
	msg := "Hello"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "goodbye"
	time.Sleep(time.Second)
}

func demo3() {
	fmt.Println("-----demo3: Wait Group-----")
	wg := sync.WaitGroup{}
	msg := "Hello"
	wg.Add(1)
	go func(msg string, w *sync.WaitGroup) {
		fmt.Println(msg)
		w.Done()
	}(msg, &wg)
	msg = "goodbye"
	wg.Wait()
}

func demo4() {
	fmt.Println("-----demo4-----")
	start := time.Now()
	names := []string{"Phil", "Noodles", "Barbaro"}

	var wg sync.WaitGroup
	wg.Add(len(names))

	for _, element := range names {
		go printName(element, &wg)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println("Total time taken", elapsed)
}

func printName(element string, wg *sync.WaitGroup) {
	result := 0.0
	for i := 0; i < 300000000; i++ {
		result += math.Pi * math.Sin(float64(len(element)))
	}
	fmt.Println(element)
	wg.Done()
}

func demo5() {
	fmt.Println("-----demo5 : sync issue-----")
	wg := sync.WaitGroup{}
	count := 0
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func(w *sync.WaitGroup) {
			fmt.Println(i)
			w.Done()
		}(&wg)

		go func(w *sync.WaitGroup) {
			count++
			w.Done()
		}(&wg)
	}

	wg.Wait()
}

func demo6() {
	fmt.Println("-----demo6 : Fix sync issue by using mutex-----")
	wg := sync.WaitGroup{}
	m := sync.RWMutex{}
	count := 0
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go func(w *sync.WaitGroup, m *sync.RWMutex) {
			//m.RLock()
			fmt.Println(i)
			m.RUnlock()
			w.Done()
		}(&wg, &m)

		m.Lock()
		go func(w *sync.WaitGroup, m *sync.RWMutex) {
			//m.Lock()
			count++
			m.Unlock()
			w.Done()
		}(&wg, &m)
	}

	wg.Wait()
}
