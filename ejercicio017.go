/*Confeccionar un programa que permita cargar un número entero positivo de hasta tres cifras y muestre un mensaje indicando si tiene 1, 2, o 3 cifras. 
Mostrar un mensaje de error si el número de cifras es mayor.*/

package main

import "fmt"


func main(){

	var numero int

	fmt.Print("Ingrese un número entre 1 y 3 cifras entero positivo: ")
	fmt.Scan(&numero)
 

	if numero > 0 && numero < 10 {
		fmt.Println("El número tiene una cifra. ")
	} else if numero > 9 && numero < 100 {
		fmt.Println("El número tiene dos cifra. ")
	} else if numero > 99 && numero < 1000 {
		fmt.Println("El número tiene tres cifra. ")
	} else if numero > 999 {
		fmt.Println("El número de cifras es mayor. ")
	}

}