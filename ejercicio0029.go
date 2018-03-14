/*Se ingresan tres notas de un alumno, si el promedio es mayor o igual a siete mostrar un mensaje "Promocionado".
Definir por inferencia el tipo dato para almacenar el promedio*/

package main

import "fmt"

func main(){
	var nota1, nota2, nota3 int

	fmt.Print("Ingrese una nota: ")
	fmt.Scan(&nota1)

	fmt.Print("Ingrese una nota: ")
	fmt.Scan(&nota2)

	fmt.Print("Ingrese una nota: ")
	fmt.Scan(&nota3)

	promedio := (float32(nota1) + float32(nota2) + float32(nota3))/3

	if promedio >= 7{
		fmt.Println("Promocionado")
	}
}