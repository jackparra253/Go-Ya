//Escribir un programa en el cual se ingresen cuatro números, calcular e informar la suma de los dos primeros y el producto del tercero y el cuarto

package main

import "fmt"

func main(){
	var numeroUno, numeroDos, numeroTres, numeroCuatro, suma, producto int

	fmt.Print("Ingrese el primer número: ")
	fmt.Scan(&numeroUno)

	fmt.Print("Ingrese el segundo número: ")
	fmt.Scan(&numeroDos)

	suma = numeroUno + numeroDos

	fmt.Print("Ingrese el tercer número: ")
	fmt.Scan(&numeroTres)

	fmt.Print("Ingrese el cuarto número: ")
	fmt.Scan(&numeroCuatro)

	producto = numeroTres * numeroCuatro

	fmt.Println("El valor de la suma es: ", suma)

	fmt.Print("El valor del producto es: ", producto)
}