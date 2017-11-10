//
//@author Michael Daniel Murillo López
//
// Funcion Factorial
//Corporación Universitaria Minuto de Dios - UNIMINUTO
//
//
package main

import (
	"fmt"
	"time"
)

const LIM = 40
var facts [LIM]uint64

func main() {
	fmt.Println("------------FACTORIAL----------------")
	start := time.Now()
	for i:=uint64(0); i < LIM; i++ {
		fmt.Printf("Factorial for %d is : %d \n", i, Factorial(uint64(i)))
	}
	end := time.Now()
	fmt.Printf ("Cálculo finalizado en% s\n", end.Sub (start)) // Cálculo finalizado en 2.0002ms

}
func Factorial(n uint64)(result uint64) {
	if (n > 0) {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}