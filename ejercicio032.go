/*Escribir un programa que solicite la carga de un valor positivo y nos muestre desde 1 hasta el valor ingresado de uno en uno.
Ejemplo: Si ingresamos 30 se debe mostrar en pantalla los números del 1 al 30.*/

package main

import "fmt"

func main(){
	var numero int

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero)

	for i := 1; i<=numero; i++ {
		fmt.Println(i)
	}
}