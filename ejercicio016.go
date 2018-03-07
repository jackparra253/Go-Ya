/*Se ingresa por teclado un valor entero, mostrar una leyenda que indique si el número es positivo, negativo o cero*/
package main

import "fmt"

func main(){
	var numero int

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero)

	if numero < 0 {
		fmt.Println("El número es negativo. ")
	} else if numero == 0 {
		fmt.Println("El número es igual a cero. ")
	} else if numero > 0 {
		fmt.Println("El número es positivo. ")
	}
}
