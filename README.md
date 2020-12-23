# Getting started with Golang

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
