/*Se ingresan por teclado tres números, si al menos uno de los valores ingresados es menor a 10, 
imprimir en pantalla la leyenda "Alguno de los números es menor a diez".*/

package main

import "fmt"

func main(){
	var numero1, numero2, numero3 int

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero1)

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero2)
	
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero3)

	if numero1 < 10 || numero2 < 10 || numero3 < 10 {
		fmt.Println("Alguno de los números es menor a diez")
	}

}