//Se debe desarrollar un programa que pida el ingreso del precio de un artículo y la cantidad que lleva el cliente. Mostrar lo que debe abonar el comprador.
//Definir una variable float32 para el precio del artículo.

package main

import "fmt"

func main(){
	var cantidad int
	var precio, abono float32
	
	fmt.Print("Ingrese el precio del artículo: ")
	fmt.Scan(&precio)

	fmt.Print("Ingrese la cantidad de artículos: ")
	fmt.Scan(&cantidad)

	abono = precio* float32(cantidad)

	fmt.Print("El valor que debe abonar el cliente es: " , abono)
}