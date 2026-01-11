package main
import ("fmt")

func main() {
  var a = make(map[string]string)
  var b map[string]string

  fmt.Println(a == nil) // false
  fmt.Println(b == nil) // true

  // El primer map a está inicializado con make, por lo que no es nil
  // El segundo map b no está inicializado, por lo que es nil
  // Esto se debe a que make crea un map vacío, mientras que no inicializarlo crea un map nil
}
