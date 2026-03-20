package main

import (
	"fmt"
	"time" // Added for time tracking
)

func main() {
    fmt.Println("Press ENTER to start the timer...")
    fmt.Scanln() // Pauses and waits for any input

    start := time.Now() // Captures exact moment
    fmt.Println("Timer started at:", start.Format("15:04:05"))

    fmt.Println("Press ENTER to stop the timer...")
    fmt.Scanln()

    elapsed := time.Since(start) // Calculates the difference
    fmt.Printf("Task complete! Duration: %s\n", elapsed)
}
