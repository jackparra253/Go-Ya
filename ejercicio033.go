/*Desarrollar un programa que permita la carga de 10 valores por teclado y nos muestre posteriormente la suma de los valores 
ingresados y su promedio.*/

package main

import "fmt"

func main(){
	var numero int
	sumatoria := 0

	for i:= 1; i <= 10 ;i ++ {
		fmt.Print("Ingrese un número: ")
		fmt.Scan(&numero)
		sumatoria = sumatoria + numero
	}

	fmt.Println("El valor de la sumatoria es: ", sumatoria)
	promedio := sumatoria / 10;
	fmt.Println("El promedio es: ", promedio)
}