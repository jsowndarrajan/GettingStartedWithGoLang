# Getting started with GoLang

## Pointers

* Pointer type use an asterisk (*) as a prefix
  * *int - a pointer to an integer
* Use the address-of (&) operator to get the memory address of a variable
* Deferencing the address type by preceding with an asterisk (*)
* Complex types (e.g., Struct) are automatically dereferenced  
* All assignment operations in Go are copy operations
* Slices and maps contain internal pointers, so copies point to the same underlying data

## Defer

* Used to delay execution of a statement until function ends
* Useful in open and close I/O operations together
  * Be careful with loops
* Run in First-In, Last-Out order
* Arguments evaluated at the time of defer statement is executed, not at the time of function exits

## Panic

* Occurs when a program cannot continue to execute at all
  * Don't use when a file cannot be opened, unless it's a critical file
  * Use for unrecoverable events (e.g. cannot obtain TCP port for a web server)
* Function will stop executing
  * Deferred function will still execute
* If the function not handling panic, then the program will exit

## Recover

* Used to recover from panics
* Only useful in deferred  functions
* Current panic function will not attempt to continue, but higher functions in call stack will

## Goroutines

* It's a lightweight thread and helps to build efficient and highly concurrent applications
* Go runtime has a scheduler and it will map the goroutines into available OS threads
* Use the `go` keyword in front of a function call to create goroutines
* When using anonymous functions, pass data as local variables - avoid a race condition, instead of relying on closures
* Use `sync.WaitGroup` to wait for groups of goroutines to complete
* Use `sync.Mutex` or `sync.RWMutex` to protect data access
* By default, Go will use CPU threads equal to # available cores
* Change the default threads count by using this comment: `runtime.GOMAXPROCS(100)`
* More threads can increase the performance, but too many can slow it down
* Best Practices
  * Don't create goroutines in libraries - let consumer control concurrency
  * When creating goroutine, know how it will end - Avoid subtle memory leaks
  * Check for race conditions at compile time
  > go run -race filePath\fileName.go

## Channels

* Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
* Create channel using `make` statement:
  * unbounded: `make(chan int)`
  * bounded: `make(chan int, 2)`
* Send message into channel: `ch <- message`
* Receive message from channel: `message := <-ch`
* Can have multiple senders and receivers
* By default, channels are bidirectionsl, but we can restrict the data as like below:
  * Send only: `chan <- int`
  * Receive only: `<-chan int`