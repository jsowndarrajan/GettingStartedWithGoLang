# Getting started with Golang

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
