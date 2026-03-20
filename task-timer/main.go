package main

import (
    "bufio"
    "fmt"
		"strings"
    "os"
    "time"
)

func main() {
    fmt.Println("1. Start New Task")
    fmt.Println("2. View Summary")
    fmt.Print("Choose an option: ")

    var choice int
    fmt.Scanln(&choice) // This expects a number

    if choice == 1 {
        runTimer() // Tip: You can move your timer code into its own function!
    } else if choice == 2 {
        showSummary()
    } else {
        fmt.Println("Invalid choice.")
    }
}

func showSummary() {
    file, err := os.Open("log.txt")
    if err != nil {
        fmt.Println("No logs found yet!")
        return
    }
    defer file.Close()

    fmt.Println("\n--- YOUR LOGS ---")
    scanner := bufio.NewScanner(file)
    for scanner.Scan() { // This loop runs for every line in the file
        fmt.Println(scanner.Text())
    }
}

func runTimer() {
		reader := bufio.NewReader(os.Stdin)

    fmt.Print("What task are you starting? ")

		// Inside your main function, after reading taskName:
		taskName, _ := reader.ReadString('\n')
		taskName = strings.TrimSpace(taskName) // This removes the hidden "Enter" character

    start := time.Now()
    fmt.Println("Timer started. Press ENTER to stop...")
    fmt.Scanln()

    duration := time.Since(start)

    // --- NEW: FILE SAVING LOGIC ---
    
    // os.OpenFile(name, flags, permissions)
    // O_APPEND: Add to the end of the file
    // O_CREATE: Create the file if it's missing
    // O_WRONLY: Open for writing only
    file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    
    // Error handling: In Go, we check if things went wrong immediately
    if err != nil {
        fmt.Println("Could not open file:", err)
        return
    }
    defer file.Close() // This ensures the file closes when the program ends

    // Format the string: Task Name | Duration | Date
    logEntry := fmt.Sprintf("Task: %s Duration: %v | Date: %s\n", 
        taskName, duration, time.Now().Format("2006-01-02"))

    file.WriteString(logEntry)
    fmt.Println("Task saved to log.txt!")
}
