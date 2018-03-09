/*Realizar un programa que pida cargar una fecha cualquiera, luego verificar si dicha fecha corresponde a Navidad.*/
package main

import "fmt"

func main(){
	var dia, mes, año int

	fmt.Print("Ingrese el día: ")
	fmt.Scan(&dia)

	fmt.Print("Ingrese el mes: ")
	fmt.Scan(&mes)
	
	fmt.Print("Ingrese el año: ")
	fmt.Scan(&año)

	if mes < 0 || mes > 12 {
		fmt.Print("El día esta fuera del intervalo. ")
	} else {
		if dia < 0 || dia > 28 && mes == 2 {
			fmt.Print("El día esta fuera del intervalo. ")
		} else {
			if dia < 0 ||  dia> 31 {
				fmt.Print("El día esta fuera del intervalo. ")
				
			}else{
				if mes == 12 && dia == 25 {
					fmt.Println("La fecha corresponde a navidad. ")
				}else{
					fmt.Println("La fecha no corresponde a navidad. ")
				}
			}
		} 
	}
}