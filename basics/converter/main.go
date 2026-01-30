package main

import "fmt"

func main() {
	celsius := 25.0
	
	Fahrenheit := (celsius * 9.0/5.0) + 32.0
	fmt.Printf("%.2f Celsius is %.2f Fahrenheit.\n", celsius, Fahrenheit)
}