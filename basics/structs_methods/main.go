package main

import (
	"fmt"
)

type Product struct {
	Name string
	Price float64
	Stock int
}

func (pr Product) Display() {
	fmt.Printf("Item: %s | Cost: $%.2f | Stock: %d\n", pr.Name, pr.Price, pr.Stock)
}

func (pr *Product) Buy(amount int) error {
	if pr.Stock < amount {
		return fmt.Errorf("No enough stock for this purchase.")
	}
	pr.Stock = pr.Stock - amount
	return nil
}

func main() {
	pr := Product{"Glass", 12.99, 5}
	pr.Display()
	err := pr.Buy(6)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	pr.Display()
}
