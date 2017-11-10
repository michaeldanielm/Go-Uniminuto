
//
//@author Michael Daniel Murillo López
//
// Funcion fibonacci
//Corporación Universitaria Minuto de Dios - UNIMINUTO
//
//
package main

import (
    "fmt"
)

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func main() {
    c := make(chan int, 47)
    go fibonacci(cap(c), c)
    for i := range c {
        fmt.Println(i)
    }
}