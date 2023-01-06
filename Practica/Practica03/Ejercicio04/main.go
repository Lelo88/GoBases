/*
Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de 
los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y 
que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una 
cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.

Ejemplo: (en el archivo de la ejercitacion)

*/

package main

import "fmt"

type Operacion func (notas ...float64) (float64)

const (
	minimum = "minimum"
	maximum = "maximum"
	average = "average"
)

func minimo(notas ...float64) (min float64){
	min = notas[0]
	for _,nota:=range notas{
		if nota<min{
			min = nota
		}
	}
	return
}

func maximo(notas ...float64) (max float64){
	max = notas[0]
	for _,nota:=range notas{
		if nota>max{
			max = nota
		}
	}
	return 
}

func calculaPromedio(notas... float64) float64{
	suma := 0.0
	for _,nota := range notas{
		suma+=nota
	}
	return float64(suma)/float64(len(notas))
}

func operacionProfesores(calculo string) (resultado Operacion){
	switch calculo {
		case minimum:
			resultado = minimo
		case maximum:
			resultado = maximo
		case average: 
			resultado = calculaPromedio	
	}
	return 
}

func ingreso() (status bool, notas []float64){
	fmt.Println("Ingrese las notas de los estudiantes")
	for {
		nota := 0
		
		fmt.Scanf("%d", &nota)
		if nota < 0 || nota > 10 {
			break
		}
		notas = append(notas, float64(nota))
	}
	
	status = true
	return 
}

func main(){
	ok, notas := ingreso()

	if !ok {
		fmt.Println("No se pueden realizar los calculos solicitados")
	} else {
		fmt.Println("Estas son las notas ingresadas:",notas)
	}

	promedio, maximo, minimo := operacionProfesores(average), operacionProfesores(maximum), operacionProfesores(minimum)
	fmt.Printf("El promedio de las notas es de %.2f. La nota minima ingresada es de %.2f y la maxima es de %.2f.\n", promedio(notas...), minimo(notas...), maximo(notas...))
}

