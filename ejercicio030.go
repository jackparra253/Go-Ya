/*Realizar un programa que solicite la carga por teclado de dos números, 
si el primero es mayor al segundo informar su suma y diferencia, 
en caso contrario informar el producto y la división del primero respecto al segundo. 
Definir variables por inferencia para almacenar la suma, resta, multiplicación y división*/

package main

import "fmt"

func main(){
	var numero1, numero2 int

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero1)

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero2)

	if numero1 > numero2 {
		suma := numero1 + numero2
		resta := numero1 - numero2
		fmt.Println("El valor de la suma es: ", suma)
		fmt.Println("El valor de la resta es: ", resta)
	} else {
		producto:= float32(numero1) * float32(numero2)
		fmt.Println("El valor del producto es: ", producto)
		if numero2 > 0 {
			division := float32(numero1) / float32(numero2)
			fmt.Println("El valor de la división es: ", division)
		}
	}
}