/*Se ingresa por teclado un número positivo de uno o dos dígitos (1..99) mostrar un mensaje indicando si el número tiene uno o 
dos dígitos.(Tener en cuenta que condición debe cumplirse para tener dos dígitos un número entero)*/

package main

import "fmt"

func main(){
	var numero int

	fmt.Print("Ingrese un número comprendido entre 1 y 99: ")
	fmt.Scan(&numero)

	if numero > 0 && numero < 10 {
		fmt.Println("El número tiene un dígito. ")
	}

	if numero > 9 && numero < 100 {
		fmt.Println("El número tiene dos dígitos. ")
	}

}

