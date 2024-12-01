package utils

import (
	"fmt"
	"time"
)

const runs = 100

func formatTime(d time.Duration) string {
	ns := d.Nanoseconds()
	switch {
	case ns < 1_000:
		return fmt.Sprintf("%d ns", ns)
	case ns < 1_000_000:
		return fmt.Sprintf("%.2f µs", float64(ns)/1_000)
	case ns < 1_000_000_000:
		return fmt.Sprintf("%.2f ms", float64(ns)/1_000_000)
	default:
		return fmt.Sprintf("%.2f s", float64(ns)/1_000_000_000)
	}
}

func averageTime(fn func() int) time.Duration {

	var totalTime time.Duration
	for i := 0; i < runs; i++ {
		start := time.Now()
		fn()
		elapsed := time.Since(start)
		totalTime += elapsed
	}

	return totalTime / runs
}

func TimeAndFormat(fn func() int, label string) {
	answer := fn()
	time := formatTime(averageTime(fn))

	fmt.Printf("The answer for %s is %d\n", label, answer)
	fmt.Printf("This took an average of %s over %d runs\n", time, runs)
	fmt.Println("----------------------------------------------------")
}

func PrintTitle(title string) {
	fmt.Println()
	fmt.Printf("\033[1m%s\033[0m\n", title)
	fmt.Println()
}
