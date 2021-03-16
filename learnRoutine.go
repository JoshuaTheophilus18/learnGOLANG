package main

import (
	"fmt"
	"time"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func LearnGoRoutine() {
	go print(5, "halo")   //pemanggilan fungsi dengan goroutine
	print(5, "apa kabar") //pemanggilan fungsi biasa

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var messages = make(chan string) //pembuatan channel (komunikasi antar goroutine)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)

	// var input string
	// fmt.Scanln(&input)
	time.Sleep(time.Second)
	fmt.Println("done")
}
