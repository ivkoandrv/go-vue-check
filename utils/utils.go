package utils

import (
	"fmt"
	"time"
)

func ElapseTimeMemory(fn func()) float64 {
	startTime := time.Now()
	fn()
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	elapsedSeconds := elapsedTime.Seconds()
	fmt.Printf("Program has been running for %.2f seconds\n", elapsedSeconds)

	return elapsedSeconds
}
