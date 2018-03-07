/*Se cargan por teclado tres números distintos. Mostrar por pantalla el mayor de ellos.*/

package main

import "fmt"

func main(){
	var numeroUno, numeroDos, numeroTres int

	fmt.Print("Ingrese el primer número. ")
	fmt.Scan(&numeroUno)

	fmt.Print("Ingrese el segundo número. ")
	fmt.Scan(&numeroDos)

	fmt.Print("Ingrese el tercer número. ")
	fmt.Scan(&numeroTres)

	if numeroUno > numeroDos && numeroUno > numeroTres {
		fmt.Println("El número mayor es: " , numeroUno)
	} else if numeroDos> numeroUno && numeroDos > numeroTres {
		fmt.Println("El número mayor es: ", numeroDos)
	} else if numeroTres > numeroDos && numeroTres > numeroUno {
		fmt.Println("El número mayor es: ", numeroTres)
	} else {
		fmt.Println("Ningún número es mayor a los demás. ")
	}
}