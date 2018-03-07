//Realizar un programa que solicite al operador ingresar dos números y muestre por pantalla el mayor de ellos.
package main

import "fmt"

func main(){
	var numeroUno, numeroDos int

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numeroUno)

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numeroDos)

	if numeroUno > numeroDos {
		fmt.Println("El número mayor es: ", numeroUno)
	} else {
		fmt.Println("El número mayor es: ", numeroDos)
	}
}