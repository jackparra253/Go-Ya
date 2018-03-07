//Ingresar el sueldo de una persona, si supera los 3000 pesos mostrar un mensaje en pantalla indicando que debe abonar impuestos.

package main

import "fmt"

func main(){
	var sueldo int

	fmt.Print("Ingrese el salario del colaborador: ")
	fmt.Scan(&sueldo)

	if sueldo > 3000
		fmt.Println("Debe abonar impuestos.")
}