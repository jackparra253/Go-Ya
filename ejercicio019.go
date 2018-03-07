
/*Un postulante a un empleo, realiza un test de capacitación, se obtuvo la siguiente información: cantidad total de preguntas que se le realizaron y la cantidad de preguntas que contestó correctamente. 
Se pide confeccionar un programa que ingrese los dos datos por teclado e informe el nivel del mismo según el porcentaje de respuestas correctas que ha obtenido, y sabiendo que:
Nivel máximo:	Porcentaje>=90%.
Nivel medio:	Porcentaje>=75% y <90%.
:	Porcentaje>=50% y <75%.
Fuera de nivel:	Porcentaje<50%.*/

package main

import "fmt"

func main(){
	var cantidadTotalPreguntas, cantidadRespuestasCorrectas int
	var porcentaje float32

	fmt.Print("Ingrese el total de preguntas realizadas. ")
	fmt.Scan(&cantidadTotalPreguntas)

	fmt.Print("Ingrese la cantidad de respuestas correctas. ")
	fmt.Scan(&cantidadRespuestasCorrectas)


	porcentaje = (float32(cantidadRespuestasCorrectas) * float32(100)) / float32(cantidadTotalPreguntas)

	if porcentaje >= 90 {
		fmt.Println("Nivel máximo. ")
	} else if porcentaje >=  75 && porcentaje <= 90 {
		fmt.Println("Nivel medio. ")
	} else if porcentaje >= 50 && porcentaje <75 {
		fmt.Println("Nivel regular. ")
	} else if porcentaje < 50 {
		fmt.Println("Fuera de nivel. ")
	}
}