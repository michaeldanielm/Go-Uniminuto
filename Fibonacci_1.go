//
//@author Michael Daniel Murillo López
//
// Funcion fibonacci
//Corporación Universitaria Minuto de Dios - UNIMINUTO
//
//
package main

import "fmt"

// fib devuelve una función que devuelve
// sucesivos números de Fibonacci.
func fib() func() int {
	a, b := 0, 1
	
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	// Las llamadas a funciones se evalúan de izquierda a derecha.
	fmt.Println(f(), f(), f(), f(), f(),f(), f(), f(), f(), f(),f(), f(), f(), f(), f(),f(), f(), f(), f(), f(),f(), f(), f(), f(), f(),f(), f(), f(), f(), f())
}
