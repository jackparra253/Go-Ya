//Hallar la superficie de un cuadrado conociendo el valor de un lado.
package main

import "fmt"

func main(){
	var lado int
	var valorSuperficie int	

	fmt.Println("Ingrese el lado de un cuadrado");
	fmt.Scan(&lado)

	valorSuperficie = lado * lado

	fmt.Println("El valor de la superficie es: ", valorSuperficie)
}

