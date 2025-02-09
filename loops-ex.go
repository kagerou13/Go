package main

import (
	"fmt"
)

// ask to order menu
func askOrder() string {
	var input string
	fmt.Print("What would you like to eat: ")
	// Get the input from the user
	fmt.Scan(&input)
	return input
}

// WRITE CONTAINS FUNCTION BELOW, chectk if order is same with menu
func contains(menu []string, order string) bool {
	for _, item := range menu {
		if order == item {
			return true
		}
	}
	return false
}

func main() {
	fastfoodMenu := []string{"Burgers", "Nuggets", "Fries"}
	fmt.Println("The fast food menu has these items:", fastfoodMenu)

	var total int
	// WRITE INDEFINITE LOOP ASKING FOR ORDERS BELOW
	var order string
	for order != "quit" {
		order = askOrder()
		if contains(fastfoodMenu, order) {
			total += 1
		} else {
			fmt.Println("Your order is not in our menu, please order again")
		}
	}

	var amount int
	// WRITE DEFINITE LOOP COUNTING $2 BILLS BELOW
	for i := 1; i <= total; i++ {
		amount += 2
	}

	fmt.Println("The total for the order is", amount)
}
