package main

import (
	"fmt"
	"time"
)

func main() {
	simpleUse()
	formatTime()
	parseTime()
	modifyTime()
	calculateExecutionDuration()
}

func simpleUse() {
	now := time.Now()

	fmt.Printf("Timestamp: %v\n", now.Unix())
	fmt.Printf("RFC3339 format: %s\n", now.Format(time.RFC3339))
	fmt.Printf("RFC822 format: %s\n", now.Format(time.RFC822))
	fmt.Printf(
		"Week day: %d, Day: %d, Month: %d, Year: %d\n",
		now.Weekday(),
		now.Day(),
		now.Month(),
		now.Year(),
	)

	time.Sleep(2 * time.Second)
	newNow := time.Now()
	fmt.Printf("Time difference: %d\n", newNow.Sub(now))
}

func formatTime() {
	now := time.Now()
	timeText := now.Format("2006-01-02 15:04:05")
	fmt.Printf("format 1: %s\n", timeText)
}

func parseTime() {
	timeText := "2023-03-15 12:33:45"
	datetime, err := time.Parse("2006-01-02 15:04:05", timeText)
	if err != nil {
		panic(err)
	}

	fmt.Printf("datetime: %s\n", datetime.Format(time.RFC3339))
}

func modifyTime() {
	datetime := time.Now()
	fmt.Printf("now: %s\n", datetime.Format(time.RFC3339))

	datetime = datetime.Add(10 * time.Second)
	fmt.Printf("after add 10 seconds: %s\n", datetime.Format(time.RFC3339))

	datetime = datetime.Add(-20 * time.Second)
	fmt.Printf("after add -20 seconds: %s\n", datetime.Format(time.RFC3339))
}

func calculateExecutionDuration() {
	fmt.Println("execution start..")
	start := time.Now()
	for i := 1; i < 50000000; i++ {
		_ = i*i - 1 + 300*i - 5000/i
	}
	duration := time.Since(start)

	fmt.Printf("execution duration is: %v\n", duration)
}
