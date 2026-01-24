package main


func main() {
	var a float64 = 7.1
	x, y := true, 3.7

	a, x = 5.5, false // antes estaban inicializas y con la línea a, x := 5.5, false, daba error. Ahora está corregida

	_, _, _ = a, x, y
}
