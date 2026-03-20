package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("What task are you starting? ")
    taskName, _ := reader.ReadString('\n')

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
