/*Escribir un programa en el cual: dada una lista de tres valores numéricos distintos se calcule e informe su rango de variación 
(debe mostrar el mayor y el menor de ellos)*/

package main

import "fmt"

func main(){
	var valor1, valor2, valor3 int

	fmt.Print("Ingrese un valor: ")
	fmt.Scan(&valor1)

	fmt.Print("Ingrese un valor: ")
	fmt.Scan(&valor2)

	fmt.Print("Ingrese un valor: ")
	fmt.Scan(&valor3)

	if valor1 > valor2 && valor2 > valor3 {
		fmt.Println("Mayor " , valor1, " Menor: ", valor3)
	} else if valor1 > valor3 && valor3 > valor2 {
		fmt.Println("Mayor " , valor1, " Menor: ", valor2)
	} else if valor2 > valor1 && valor1 > valor3 {
		fmt.Println("Mayor " , valor2, " Menor: ", valor3)
	} else if valor2 > valor3 && valor3 > valor1 {
		fmt.Println("Mayor " , valor2, " Menor: ", valor1)
	} else if valor3 > valor1 && valor1 > valor2 {
		fmt.Println("Mayor " , valor3, " Menor: ", valor2)
	} else if valor3 > valor2 && valor2 > valor1 {
		fmt.Println("Mayor " , valor3, " Menor: ", valor1)
	}
}