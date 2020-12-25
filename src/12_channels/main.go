package main

import (
	"fmt"
	"sync"
)

func main() {
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
}

func demo1() {
	fmt.Println("----- demo1 : One sender, One receiver -----")
	wg := sync.WaitGroup{}
	ch := make(chan int)
	defer close(ch)

	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println("received message:", i)
		wg.Done()
	}()

	go func() {
		ch <- 02
		wg.Done()
	}()

	wg.Wait()
}

func demo2() {
	fmt.Println("----- demo2: Multiple sender, multiple receiver -----")
	wg := sync.WaitGroup{}
	ch := make(chan int)
	defer close(ch)

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go func() {
			j := <-ch
			fmt.Println("received message:", j)
			wg.Done()
		}()

		go func(message int) {
			ch <- message
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func demo3() {
	fmt.Println("----- demo3: send and receive multiple messages -----")
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	go func() {
		for i := range ch {
			fmt.Println("received message:", i)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		wg.Done()
		close(ch)
	}()
	wg.Wait()
}

func demo4() {
	fmt.Println("----- demo4: send message on close channel -----")
	wg := sync.WaitGroup{}
	ch := make(chan int, 3)

	wg.Add(2)
	go func() {
		for i := range ch {
			fmt.Println("received message:", i)
		}
		wg.Done()
	}()
	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Error:", err)
			}
		}()
		for i := 0; i < 5; i++ {
			ch <- i
			if i == 3 {
				close(ch)
			}
		}
	}()
	wg.Wait()
}

func demo5() {
	fmt.Println("----- demo5: use select statment to close the channel on different conditions -----")
	wg := sync.WaitGroup{}
	ch := make(chan int, 3)
	doneChannel := make(chan struct{})
	defer close(ch)
	defer close(doneChannel)
	wg.Add(2)
	go func() {
		for {
			select {
			case message := <-ch:
				fmt.Println("received message:", message)
			case <-doneChannel:
				wg.Done()
				break
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
		}
		doneChannel <- struct{}{}
	}()
	wg.Wait()
}
