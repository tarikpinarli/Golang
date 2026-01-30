package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var birthYear int
	thisYear := time.Now().Year()
	fmt.Println("What year were you born?")
	
	_, err := fmt.Scan(&birthYear)
	if err != nil {
		fmt.Println("Error: Please enter a valid number.")
		os.Exit(1)
	}
	if birthYear > thisYear || birthYear < 1900 {
		fmt.Println("Error: Please enter a valid birth year.")
		os.Exit(1)
	}
	age := thisYear - birthYear
	fmt.Printf("You are approximately %d years old.\n", age)
}
