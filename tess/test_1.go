package main

import (
	"fmt"
	"strings"
	"time"
)

func parseDuration(durationString string) (time.Duration, error) {
	parts := strings.Split(durationString, "d")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid duration format")
	}

	days, err := time.ParseDuration(parts[0] + "h")
	if err != nil {
		return 0, err
	}

	rest, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, err
	}

	return days + rest, nil
}

func main() {
	durationString := "2d3h30m15s" // Example duration string: 2 days, 3 hours, 30 minutes, and 15 seconds

	duration, err := parseDuration(durationString)
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return
	}

	fmt.Println("Duration:", duration)
}
