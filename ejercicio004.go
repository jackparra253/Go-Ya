//Realizar la carga de dos números enteros por teclado e imprimir su suma y su producto.

package main

import "fmt"

func main(){
	var primerNumero, segundoNumero, suma, producto int	

	fmt.Println("Ingrese el primer número: ")
	fmt.Scan(&primerNumero)

	fmt.Println("Ingrese el segundo número: ")
	fmt.Scan(&segundoNumero)

	suma = primerNumero + segundoNumero
	fmt.Println("El valor de la suma es: " , suma)

	producto = primerNumero * segundoNumero
	fmt.Println("El valor del producto es: ", producto)
}

