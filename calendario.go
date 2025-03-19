package main

import (
	"fmt"
	"time"
)

func main(){
	// Definición de las fechas de inicio y fin
	fechaInicio := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	fechaFin := time.Date(2025, 8, 31, 0, 0, 0, 0, time.UTC)

	// Imprimir  una lista con los dias
	fmt.Println("Lista de días:")

	for fecha := fechaInicio; !fecha.After(fechaFin); fecha = fecha.AddDate(0, 0, 1){
		fmt.Println(fecha.Format("02/01/2006 "))
	}
}