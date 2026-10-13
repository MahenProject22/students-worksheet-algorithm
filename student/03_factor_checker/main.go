package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	// Validate BEFORE dividing: both numbers must be positive.
	if ____ || ____ {
		fmt.Println("invalid input")
		return
	}
	// x is a factor of y when the remainder of y divided by x is zero.
	fmt.Println(____)
}
