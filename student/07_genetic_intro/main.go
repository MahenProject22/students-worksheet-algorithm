package main

import "fmt"

func fitness(candidate string) int {
    
    count := 0
    for i := 0; i < len(candidate); i++ {
        if candidate[i] == '1' {
            count++
        }
    }
    return count
}

func main() {
    a, b := "1001", "0110"

  
    child := a[:2] + b[2:]

   
    mutated := child[:1] + "1" + child[2:]

    fmt.Println("Parent A:", a)
    fmt.Println("Parent B:", b)
    fmt.Println("Crossover:", child)
    fmt.Println("Mutation:", mutated)
    fmt.Println("Fitness:", fitness(mutated))
}

