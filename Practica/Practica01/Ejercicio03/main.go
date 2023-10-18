/*Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I para poder brindarles las correspondientes devoluciones. Uno de los puntos del examen consiste en declarar distintas variables. 
Necesita ayuda para:
Detectar cuáles de estas variables que declaró el alumno son correctas.
Corregir las incorrectas.
*/

package main

import "fmt"

func main(){
	/*var 1nombre string -> la denominacion de las variables no pueden comenzar con un numero*/
	var nombre1 string 
	var apellido string //definicion de variable correcta
	/*var int edad -> el tipo de variable debe colocarse luego de definir el nombre de la variable*/
	var edad int 
	/*1apellido := 6  -> la denominacion de la variable no puede comenzar con un numero. 
	Tambien es de mala practica definir una variable que posiblemente sea de otro tipo, como es en este caso */

	apellido1 :=  6

	var licencia_de_conducir = true //a pesar de que el tipo no esta puesto, go lo reconoce mediante la definicion

	/*var estatura de la persona int -> forma incorrecta de definir a una variable*/
	
	var estaturaDeLaPersona int

	fmt.Println(nombre1, apellido, apellido1, edad, licencia_de_conducir, estaturaDeLaPersona)
}