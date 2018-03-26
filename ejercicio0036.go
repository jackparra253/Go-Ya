/* Se ingresan un conjunto de n alturas de personas por teclado. Mostrar la altura promedio de las personas. */

package main

import "fmt"

func main() {
	var cantidadDeEvaluados int
	var estatura, promedio float32 

	fmt.Println("¿A cuantas personas le va a ingresar la altura por teclado? ")
	fmt.Scan(&cantidadDeEvaluados)

	for i:= 0; i < cantidadDeEvaluados; i++ {
		fmt.Println("Ingrese la estatura de la persona: ")
		fmt.Scan(&estatura)
		
		promedio = (promedio + estatura) / float32(cantidadDeEvaluados)
	}

	fmt.Println("La altura promedio es: ", promedio)
}