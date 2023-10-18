/*Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos tipos en Go, 
pero tuvo varias dudas mientras lo hacía. A partir de esto, nos brindó su código y pidió la ayuda 
de un desarrollador experimentado que pueda:
Verificar su código y realizar las correcciones necesarias.
*/

package main

import "fmt"

func main(){
	var apellido string = "Gomez" //ok
	var edad int = 35//"35" //variable edad mal definida. Le sacamos las comillas 
	boolean := false //ok
	var sueldo float64 = 45857.90 //string = 45857.90 //tipo de variable mal definida. debe ser tipo float
	var nombre string = "Julian" //ok

	fmt.Println(apellido, edad, boolean, sueldo, nombre)
}