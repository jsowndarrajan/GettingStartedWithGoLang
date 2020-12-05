package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	hour := time.Now().Hour()
	greeting, err := getGreeting(hour)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}

func getGreeting(hour int) (string, error) {
	var message string

	if hour < 4 {
		err := errors.New("too early for greeting: go and sleep again")
		return message, err
	}

	if hour < 12 {
		message = "Good Morning!"
	} else if hour < 17 {
		message = "Good Afternoon"
	} else {
		message = "Good Evening"
	}
	return message, nil
}
