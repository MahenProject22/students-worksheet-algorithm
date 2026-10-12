package main

import "fmt"

func main() {
    massKg := 70.0
    earthGravity := 9.8
    marsFactor := 0.38

    earthWeight := massKg * earthGravity
    marsWeight := earthWeight * marsFactor

    fmt.Printf("Earth weight: %.1f N\n", earthWeight)
    fmt.Printf("Mars weight: %.1f N\n", marsWeight)
}
