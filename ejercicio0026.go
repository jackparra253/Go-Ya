/*Escribir un programa que pida ingresar la coordenada de un punto en el plano, es decir dos valores enteros x e y 
(distintos a cero). Posteriormente imprimir en pantalla en que cuadrante se ubica dicho punto. 
(1º Cuadrante si x > 0 Y y > 0 , 2º Cuadrante: x < 0 Y y > 0, etc.)
*/

package main

import "fmt"

func main(){
	var x, y int

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&x)

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		fmt.Print("El valor se encuentra en el primer cuadrante. ")		
	} 

	if x < 0 && y > 0 {
		fmt.Print("El valor se encuentra en el segundo cuadrante. ")		
	}

	if x < 0 &&  y < 0 {
		fmt.Print("El valor se encuentra en el segundo cuadrante. ")		
	} 

	if x > 0 && y < 0 {
		fmt.Print("El valor se encuentra en el cuarto cuadrante. ")		
	}

	if x == 0 && y == 0 {
		fmt.Print("El valor se encuentra en las coordenadas 0 y 0. ")		
	}
}