package main

import (
	"fmt"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func LearnGoRoutine() {
	go print(5, "halo")   //pemanggilan fungsi dengan goroutine
	print(5, "apa kabar") //pemanggilan fungsi biasa

	var input string
	fmt.Scanln(&input)
}
