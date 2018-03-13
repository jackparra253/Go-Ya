/*De un operario se conoce su sueldo y los años de antigüedad. Se pide confeccionar un programa que lea los datos de entrada e 
informe:
a) Si el sueldo es inferior a 500 y su antigüedad es igual o superior a 10 años, otorgarle un aumento del 20 %,
 mostrar el sueldo a pagar.
b)Si el sueldo es inferior a 500 pero su antigüedad es menor a 10 años, otorgarle un aumento de 5 %.
c) Si el sueldo es mayor o igual a 500 mostrar el sueldo en pantalla sin cambios.*/
package main

import "fmt"

func main(){
	var sueldo, antiguedad int
	var aumento, total float32

	fmt.Print("Ingrese el sueldo del colborador: ")
	fmt.Scan(&sueldo)

	fmt.Print("Ingrese los años de antigüedad: ")
	fmt.Scan(&antiguedad)

	if sueldo < 500 && antiguedad >= 10 {
		aumento = float32(sueldo) * 0.2

		total = float32(sueldo) + aumento
		fmt.Println("Sueldo a pagar: " , total)
	} else if sueldo < 500 && antiguedad < 5 {
		aumento = float32(sueldo) * 0.05
		total = float32(sueldo) + aumento
		fmt.Println("Sueldo a pagar: " , total)
	} else if sueldo > 500 {
		fmt.Println("Sueldo a pagar: " , sueldo)
	}
}
