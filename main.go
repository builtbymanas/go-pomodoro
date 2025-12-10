package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <work-duration> <break-duration>")
		fmt.Println("Example: go run main.go 25m 5m")
		fmt.Println("Example: go run main.go 1h30m 15m")
		return
	}

	workInput := os.Args[1]
	breakInput := os.Args[2]

	workDur, err1 := time.ParseDuration(workInput)
	breakDur, err2 := time.ParseDuration(breakInput)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Could not understand time formats.")
		fmt.Println("Please use formats like: 10s, 5m, 1h30m")
		return
	}

	fmt.Printf("🍅 Pomodoro Timer Started for %s...\n", workDur)
	startTimer(int(workDur.Seconds()), "Work")

	breakMsg := fmt.Sprintf("Work session ended. Starting break for %s.", breakDur)
	triggerNotification("Pomodoro Timer", breakMsg)

	fmt.Printf("\n☕ Starting Break Session for %s...\n", breakDur)
	startTimer(int(breakDur.Seconds()), "Break")

	triggerNotification("Pomodoro Timer", "Break Over! Ready for the next round?")
	fmt.Println("\n✅ Session Complete!")
}

func startTimer(totalSeconds int, label string) {
	for i := totalSeconds; i > 0; i-- {
		h := i / 3600
		m := (i % 3600) / 60
		s := i % 60

		fmt.Printf("\r⏳ %s Remaining: %02d:%02d:%02d   ", label, h, m, s)
		time.Sleep(1 * time.Second)
	}
}

func triggerNotification(title string, message string) {
	cmdNotify := exec.Command("notify-send", title, message)
	_ = cmdNotify.Run()

	cmdSound := exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga")
	_ = cmdSound.Run()
}
