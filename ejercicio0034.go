//10 - Estructura de programación repetitiva for (solo con condición)

/*Una planta que fabrica perfiles de hierro posee un lote de n piezas.
Confeccionar un programa que pida ingresar por teclado la cantidad de piezas a procesar y 
luego ingrese la longitud de cada perfil; 
sabiendo que la pieza cuya longitud esté comprendida en el rango de 1.20 y 1.30 son aptas. 
Imprimir por pantalla la cantidad de piezas aptas que hay en el lote.*/
package main

import "fmt"

func main(){
	var cantidad int
	var longitudPerfil float32

	fmt.Print("Ingrese la cantidad de piezas a procesar: ")
	fmt.Scan(&cantidad)

	cantidadPiezasAptas := 0

	for i:= 1; i <= cantidad; i++ {
		fmt.Print("Ingrese la longitud de cada perfil. ")
		fmt.Scan(&longitudPerfil)
		if longitudPerfil >= 1.20 && longitudPerfil <= 1.30 {
			cantidadPiezasAptas = cantidadPiezasAptas + 1
		}
	}

	fmt.Print("La cantidad de piezas aptas que hay en el lote es: " , cantidadPiezasAptas)
}