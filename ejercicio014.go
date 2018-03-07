/*Confeccionar un programa que pida por teclado tres notas de un alumno, calcule el promedio e imprima alguno de estos mensajes:
Si el promedio es >=7 mostrar "Promocionado".
Si el promedio es >=4 y <7 mostrar "Regular".
Si el promedio es <4 mostrar "Reprobado".*/

package main

import "fmt"

func main(){
	var notaUno, notaDos, notaTres int

	fmt.Print("Ingrese la primera nota. ")
	fmt.Scan(&notaUno)

	fmt.Print("Ingrese la segunda nota. ")
	fmt.Scan(&notaDos)

	fmt.Print("Ingrese la tercera nota. ")
	fmt.Scan(&notaTres)

	if notaUno >= 0 && notaDos >= 0 && notaTres >= 0 && notaUno <= 10 && notaDos <= 10 && notaTres <= 10{
		var promedio float32

		promedio = (float32(notaUno) + float32(notaDos) + float32(notaTres))/3

		if promedio < 4 {
			fmt.Println("Reprobado. ")
		} else if promedio >= 4 && promedio <= 7 {
			fmt.Println("Regular. ")
		} else if promedio > 7 {
			fmt.Println("Promocionado. ")
		}
	} else {
		fmt.Println("Cada nota ingresada debe estar entre 0 y 10. ")
	}
}
