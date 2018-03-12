/*Se ingresan por teclado tres números, si todos los valores ingresados son menores a 10, 
imprimir en pantalla la leyenda "Todos los números son menores a diez".*/

package main

import "fmt"

func main(){
	var numero1, numero2, numero3 int
	var diez int = 10

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero1)

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero2)

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero3)

	if numero1 < diez && numero2 < diez && numero3 < diez {
		fmt.Println("Todos los números son menores a diez")
	}
}