/*En una empresa trabajan n empleados cuyos sueldos oscilan entre $100 y $500, 
realizar un programa que lea los sueldos que cobra cada empleado e informe cuántos empleados cobran entre $100 y $300 y 
cuántos cobran más de $300. Además el programa deberá informar el importe que gasta la empresa en sueldos al personal.*/

package main

import "fmt"

func main(){
	var cantidadEmpleados, sueldo,sueldosEntre100y300, sueldoMayoresA300, totalSueldos int

	fmt.Println("Ingrese la cantidad de empleados: ")
	fmt.Scan(&cantidadEmpleados)

	for i:= 0; i < cantidadEmpleados; i ++ {
		fmt.Println("Ingrese el sueldo del empleado: ")
		fmt.Scan(&sueldo)

		if sueldo >= 100 && sueldo <= 300 {
			sueldosEntre100y300 = sueldosEntre100y300 + 1
		}

		if sueldo > 300 {
			sueldoMayoresA300 = sueldoMayoresA300 + 1
		}

		totalSueldos = totalSueldos + sueldo
	}

	fmt.Println("La cantidad de sueldos entre 100 y 300: ", sueldosEntre100y300)
	fmt.Println("La cantidad de sueldo mayores a 300: ", sueldoMayoresA300)
	fmt.Println("Importe total que gasta la empresa en sueldos al personal: " , totalSueldos)

}