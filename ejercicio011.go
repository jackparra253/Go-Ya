/*Realizar un programa que solicite la carga por teclado de dos números, 
si el primero es mayor al segundo informar su suma y diferencia, en caso contrario informar el producto 
y la división del primero respecto al segundo.*/

package main

import "fmt"

func main(){
	var numeroUno, numeroDos int

	fmt.Print("Ingrese el primer número: ")
	fmt.Scan(&numeroUno)

	fmt.Print("Ingrese el segundo número: ")
	fmt.Scan(&numeroDos)

	if numeroUno > numeroDos {

		var suma, diferencia int

		suma = numeroUno + numeroDos
		diferencia = numeroUno - numeroDos

		fmt.Println("La suma es: " , suma)
		fmt.Println("La diferencia es: " , diferencia)

	} else {
		var producto int
		var division float32

		producto = numeroUno * numeroDos

		fmt.Println("El producto es: " , producto)

		if numeroDos != 0 {
			division = float32(numeroUno) / float32(numeroDos)
			fmt.Println("La división es: ", division)
		}
	}
}