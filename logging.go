package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func logMessage(message string) {
	msg := strings.TrimSpace(message)
	fmt.Printf("Message: %s\n", msg)
	writeMessageToFile(readConfig()["logFileLocation"], msg)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeMessageToFile(filename string, message string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()
	currentTime := time.Now().Format("15:04:05")
	if _, err := f.Write([]byte(currentTime + " " + message + "\n")); err != nil {
		check(err)
	}
}
