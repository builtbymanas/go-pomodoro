package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <duration>")
		fmt.Println("Example: go run main.go 25m")
		fmt.Println("Example: go run main.go 1h30m")
		return
	}

	input := os.Args[1]
	duration, err := time.ParseDuration(input)

	if err != nil {
		fmt.Printf("Error: Could not understand time '%s'.\n", input)
		fmt.Println("Please use formats like: 10s, 5m, 1h30m")
		return
	}

	fmt.Printf("🍅 Pomodoro Timer Started for %s...\n", duration)
	startTimer(int(duration.Seconds()))
}

func startTimer(seconds int) {
	for i := seconds; i > 0; i-- {
		h := i / 3600
		m := (i % 3600) / 60
		s := i % 60

		fmt.Printf("\r⏳ Time Remaining: %02d:%02d:%02d   ", h, m, s)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("\n✅ Timer Finished!")

	triggerFunction()
}

func triggerFunction() {
	cmd := exec.Command("notify-send", "Pomodoro Timer", "Session Complete! Take a break.")
	_ = cmd.Run()
}
