/*Realizar la carga de dos enteros por teclado. Mostrar la suma y diferencia de los mismos.*/

package main

import "fmt"

func main(){
	var valor1, valor2 int
	
	fmt.Print("Ingrese un valor: ")
	fmt.Scan(&valor1)

	fmt.Print("Ingrese un valor: ")
	fmt.Scan(&valor2)

	suma := valor1 + valor2
	fmt.Println("El resultado de la suma es: " , suma)

	diferencia := valor1 - valor2
	fmt.Println("El resultado de la diferencia es: " , diferencia)
}