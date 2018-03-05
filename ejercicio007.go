//Realizar un programa que solicite por teclado cuatro valores numéricos enteros e informar su suma y promedio

package main

import "fmt"

func main(){
	var numeroUno, numeroDos, numeroTres, numeroCuatro, suma, promedio int

	fmt.Print("Ingrese un número: ")
    fmt.Scan(&numeroUno)

	fmt.Print("Ingrese un número: ")
    fmt.Scan(&numeroDos)

	fmt.Print("Ingrese un número: ")
    fmt.Scan(&numeroTres)
	
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numeroCuatro)

	suma = numeroUno + numeroDos + numeroTres + numeroCuatro
	promedio = suma / 4

	fmt.Println("El valor de la suma es: ", suma)
    fmt.Println("El valor del promedio es: ", promedio)

}