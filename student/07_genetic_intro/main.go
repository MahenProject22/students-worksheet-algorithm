package main

import "fmt"

func fitness(candidate string) int {
	count := 0
	for i := 0; i < len(candidate); i++ {
		if candidate[i] == ____ { // count the one-bits (a single character)
			count++
		}
	}
	return count
}

func main() {
	a, b := "1001", "0110"

	// Crossover: the first two bits of a, then the rest of b.
	child := a[:2] + b[____:]

	// Mutation: change the second bit of the child to 1.
	mutated := child[:1] + ____ + child[2:]

	fmt.Println("Parent A:", a)
	fmt.Println("Parent B:", b)
	fmt.Println("Child:", child, "fitness:", fitness(child))
	fmt.Println("Mutated:", mutated, "fitness:", fitness(mutated))
}
