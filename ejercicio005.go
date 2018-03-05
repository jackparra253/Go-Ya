//Realizar la carga del lado de un cuadrado, mostrar por pantalla el perímetro del mismo
package main

import "fmt"

func main(){
	var lado, perimetro int

	fmt.Print("Ingrese el lado de un cuadrado: ")
	fmt.Scan(&lado)

	perimetro = lado * 4

	fmt.Print("El perímero del cuadrado es: ", perimetro)
}