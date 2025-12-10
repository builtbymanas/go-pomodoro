package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

var soundFS embed.FS

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <work-duration> <break-duration>")
		return
	}

	workDur, _ := time.ParseDuration(os.Args[1])
	breakDur, _ := time.ParseDuration(os.Args[2])

	fmt.Printf("🍅 Work for %s...\n", formatDuration(workDur))

	startTimer(int(workDur.Seconds()), "Work")

	triggerNotification("Pomodoro", fmt.Sprintf("Work done! Break for %s.", formatDuration(breakDur)))
	playSound("sounds/session_work.wav")

	fmt.Printf("\n☕ Break for %s...\n", formatDuration(breakDur))

	startTimer(int(breakDur.Seconds()), "Break")

	triggerNotification("Pomodoro", "Break Over! Back to work.")
	playSound("sounds/session_break.wav")
}

func startTimer(totalSeconds int, label string) {
	for i := totalSeconds; i > 0; i-- {
		h := i / 3600
		m := (i % 3600) / 60
		s := i % 60
		fmt.Printf("\r⏳ %s: %02d:%02d:%02d   ", label, h, m, s)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}

func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60

	res := ""

	if h > 0 {
		res += fmt.Sprintf("%dh", h)
	}

	if m > 0 {
		res += fmt.Sprintf("%dm", m)
	}

	if s > 0 {
		res += fmt.Sprintf("%ds", s)
	}

	return res
}

// ---------------------------------------------------------
// CROSS-PLATFORM LOGIC
// ---------------------------------------------------------

func triggerNotification(title, message string) {
	switch runtime.GOOS {
	case "linux":
		exec.Command("notify-send", title, message).Run()
	case "darwin":
		script := fmt.Sprintf("display notification \"%s\" with title \"%s\"", message, title)
		exec.Command("osascript", "-e", script).Run()
	case "windows":
		script := fmt.Sprintf("Add-Type -AssemblyName System.Windows.Forms; [System.Windows.Forms.MessageBox]::Show('%s', '%s')", message, title)
		exec.Command("powershell", "-c", script).Run()
	}
}

func playSound(filename string) {
	data, err := soundFS.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading sound:", err)
		return
	}

	tmpFile := filepath.Join(os.TempDir(), "pomodoro-"+filepath.Base(filename))
	os.WriteFile(tmpFile, data, 0644)
	defer os.Remove(tmpFile)

	switch runtime.GOOS {
	case "linux":
		exec.Command("paplay", tmpFile).Run()
	case "darwin":
		exec.Command("afplay", tmpFile).Run()
	case "windows":
		script := fmt.Sprintf("(New-Object Media.SoundPlayer '%s').PlaySync();", tmpFile)
		exec.Command("powershell", "-c", script).Run()
	}
}
