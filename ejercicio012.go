/*Se ingresan tres notas de un alumno, si el promedio es mayor o igual a siete mostrar un mensaje "Promocionado".*/
package main

import "fmt"

func main(){
	var notaUno, notaDos, notaTres int

	fmt.Print("Ingrese la primera nota: ")
	fmt.Scan(&notaUno)

	fmt.Print("Ingrese la segunda nota: ")
	fmt.Scan(&notaDos)

	fmt.Print("Ingrese la tercera nota: ")
	fmt.Scan(&notaTres)

	if notaUno > 0 || notaDos > 0 || notaTres > 0 {
		var promedio float32
		promedio = (float32(notaUno) + float32(notaDos) + float32(notaTres)) / 3
	}

	if promedio >= 7 {
		fmt.Println("Promocionado ")
	}
}