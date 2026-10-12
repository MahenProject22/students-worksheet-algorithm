
package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)


    if x*x*x == y {
        fmt.Println(true)
    } else {
        fmt.Println(false)
    }
}
