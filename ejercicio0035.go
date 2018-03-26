/*Escribir un programa que solicite ingresar 10 notas de alumnos y nos informe cuántos tienen
 notas mayores o iguales a 7 y cuántos menores.*/

 package main

 import "fmt"

 func main() {
	 var nota, cantidadNotasMayores, cantidadNotasMenores int

	 for i:= 0; i < 10; i++ {
		 fmt.Println("Ingrese la nota del estudiante: ")
		 fmt.Scan(&nota)

		 if nota >= 7 {
			 cantidadNotasMayores = cantidadNotasMayores + 1
		 }

		 if nota >= 0 && nota < 7 {
			 cantidadNotasMenores = cantidadNotasMenores + 1
		 }
	 }

	 fmt.Println("Cantidad de notas mayores o igual a siete: ", cantidadNotasMayores)
	 fmt.Println("Cantidad de notas menores a siete:", cantidadNotasMenores)
 }