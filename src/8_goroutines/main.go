package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(3)

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
