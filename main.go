package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Pomodoro Timer Started...")
	startTimer(3725)
}

func startTimer(seconds int) {
	for i := seconds; i > 0; i-- {
		h := i / 3600
		m := (i % 3600) / 60
		s := i % 60

		fmt.Printf("\r⏳ Time Remaining: %02d:%02d:%02d   ", h, m, s)

		time.Sleep(1 * time.Second)
	}
	fmt.Println("\nTimer Finished!")
}
