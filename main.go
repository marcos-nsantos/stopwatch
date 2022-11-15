package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var timeInMinutes = flag.Int("minutes", 0, "time in minutes")
var timeInSeconds = flag.Int("seconds", 0, "time in seconds")
var timeInHours = flag.Int("hours", 0, "time in hours")
var name = flag.String("name", "unknown", "activity name")

func main() {
	flag.Parse()

	if !isFlagsValid() || !isTimeValid() {
		fmt.Println("Please enter a valid time")
		return
	}

	definedTime := defineTime()

	fmt.Println("Welcome to the stopwatch app!")
	fmt.Println("\nActivity name:", *name)

	fmt.Println("Starting count time...")
	fmt.Println("\nPress Enter to stop the timer")
	countTime(definedTime)
	fmt.Println("\nTime is over!")
}

func countTime(definedTime int) {
	tick := time.Tick(time.Second * 1)
	abort := make(chan bool)
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- true
	}()

	for value := 0; value < definedTime; value++ {
		select {
		case <-tick:
			showTime(value + 1)
		case <-abort:
			fmt.Println("Abort!")
			return
		}
	}
}

func showTime(value int) {
	var hours int
	var minutes int
	var seconds int

	hours = value / 3600
	minutes = (value % 3600) / 60
	seconds = (value % 3600) % 60

	fmt.Printf("hours: %02d, minutes: %02d, seconds: %02d\r", hours, minutes, seconds)
}

func isFlagsValid() bool {
	if *timeInMinutes == 0 && *timeInSeconds == 0 && *timeInHours == 0 {
		return false
	}
	return true
}

func isTimeValid() bool {
	if *timeInMinutes < 0 || *timeInSeconds < 0 || *timeInHours < 0 {
		return false
	}
	return true
}

func defineTime() int {
	var definedTime int
	if *timeInHours != 0 {
		definedTime = *timeInHours * 3600
	}

	if *timeInMinutes != 0 {
		definedTime = *timeInMinutes * 60
	}

	if *timeInSeconds != 0 {
		definedTime = *timeInSeconds
	}
	return definedTime
}
