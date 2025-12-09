package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Pomodoro Timer Started...")
	startTimer(13)
}

func startTimer(seconds int) {
	for i := seconds; i > 0; i-- {
		fmt.Printf("\rTime Remaining: %d seconds... ", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("\nTimer Finished!")
}
