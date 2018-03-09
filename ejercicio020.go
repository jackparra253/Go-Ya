//Confeccionar un programa que lea por teclado tres números distintos y nos muestre el mayor.

package main

import "fmt"

func main(){

	var numeroUno, numeroDos, numeroTres int

	fmt.Print("Ingrese un número:")
	fmt.Scan(&numeroUno)

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numeroDos)

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numeroTres)

	if numeroUno > numeroDos && numeroUno > numeroTres {
		fmt.Println("El número mayor es: " , numeroUno)
	} else if numeroDos > numeroTres && numeroDos > numeroUno {
		fmt.Println("El número mayor es: ", numeroDos)
	} else if numeroTres > numeroUno && numeroTres > numeroDos {
		fmt.Println("El número mayor es: ", numeroTres)
	} else {
		fmt.Println("Ningún número es mayor que el otro. ")
	}
}