package main
import ("fmt")

func main() {
  var a = make(map[string]string) // Crear un map vacío
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  fmt.Println(a)

  delete(a,"year") // Eliminar una clave

  fmt.Println(a)

  // SI queremos agregar una clave
  a["color"] = "red"
  fmt.Println(a) // map[brand:Ford color:red model:Mustang]
}