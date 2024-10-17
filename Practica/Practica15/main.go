package main

import (
	"bufio"
	"fmt"
	"os"
)

// Vamos a crear un listado de tareas

type Tarea struct {
	nombre      string
	descripcion string
	completado  bool
}

type ListadoTareas struct {
	tareas []Tarea // slice de tareas que va a almacenar datos
}

func (lt *ListadoTareas) AgregarTarea(t Tarea) {
	lt.tareas = append(lt.tareas, t) // al slice tarea le agrego una tarea
}

func (lt *ListadoTareas) CompletarTarea(index int) {
	lt.tareas[index].completado = true
}

func (lt *ListadoTareas) EditarTarea(index int, t Tarea) {
	lt.tareas[index] = t
}

func (lt *ListadoTareas) EliminarTarea(index int) {
	lt.tareas = append(lt.tareas[:index], lt.tareas[index+1:]...)
}

/*

 */

func main() {
	lista := ListadoTareas{}

	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println("Seleccione una tarea")
		fmt.Println("1. Agregar tarea")
		fmt.Println("2. Completar tarea")
		fmt.Println("3. Editar tarea")
		fmt.Println("4. Eliminar tarea")
		fmt.Println("5. Salir")

		fmt.Println("Ingrese una opción: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			var tarea Tarea
			fmt.Println("Ingrese el nombre de la tarea: ")
			tarea.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripción de la tarea: ")
			tarea.descripcion, _ = leer.ReadString('\n')
			tarea.completado = false
			lista.AgregarTarea(tarea)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Println("Ingrese el número de la tarea que desea completar: ")
			fmt.Scan(&index)
			lista.CompletarTarea(index)
			fmt.Println("Tarea completada correctamente")
		case 3:
			var index int
			var tarea Tarea
			fmt.Println("Ingrese el número de la tarea que desea editar: ")
			fmt.Scan(&index)
			fmt.Println("Ingrese el nombre de la tarea: ")
			tarea.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripción de la tarea: ")
			tarea.descripcion, _ = leer.ReadString('\n')
			tarea.completado = false
			lista.EditarTarea(index, tarea)
			fmt.Println("Tarea editada correctamente")
		case 4:
			var index int
			fmt.Println("Ingrese el número de la tarea que desea eliminar: ")
			fmt.Scan(&index)
			lista.EliminarTarea(index)
			fmt.Println("Tarea eliminada correctamente")
		case 5:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no válida")
		}

		fmt.Println("Listado de tareas")
		fmt.Println("*****************")
		for i, t := range lista.tareas {
			fmt.Println(i, t.nombre, t.descripcion, t.completado)
		}
		fmt.Println("*****************")
	}

}
