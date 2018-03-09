/*
Se carga una fecha (día, mes y año) por teclado. 
Mostrar un mensaje si corresponde al primer trimestre del año (enero, febrero o marzo) 
Cargar por teclado el valor numérico del día, mes y año. 
Ejemplo: dia:10 mes:5 año:2017.
*/
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
				if mes >= 1 && mes <= 3 {
					fmt.Print("corresponde al primer trimestre del año (enero, febrero o marzo). ")
				}
			}
		} 
	}
}