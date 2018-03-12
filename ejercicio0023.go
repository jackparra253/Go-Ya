/*Se ingresan tres valores por teclado, si todos son iguales se imprime la suma del primero con el segundo y 
a este resultado se lo multiplica por el tercero.*/

package main

import "fmt"

func main(){
	var valor1, valor2, valor3, suma, total int

	fmt.Print("Ingrese un valor")
	fmt.Scan(&valor1)

	fmt.Print("Ingrese un valor")
	fmt.Scan(&valor2)

	fmt.Print("Ingrese un valor")
	fmt.Scan(&valor3)

	if valor1 == valor2 && valor1 == valor3 {
		suma = valor1 + valor2
		total = suma * valor3
		fmt.Println("El valor total es: " , total)
	}
}