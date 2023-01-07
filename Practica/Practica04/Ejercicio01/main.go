package main

import "fmt"

type Product struct {
	Id 		int
	Price 	float64
	Name, Description, Category string
}

type prods []Product

var productos=prods{
	{
		Id: 43,
		Price: 54.5,
		Name: "Banana",
		Description: "Fruits",
		Category: "F",
	},
	{
		Id: 343,
		Price: 105.6,
		Name: "Milanesa de pollo",
		Description: "Carne de pollo",
		Category: "P",
	},
}

func (p *Product) Save(){
	productos = append(productos, *p) //como es un slice, utilizo el append cada vez que vaya a agregar un producto
}

func (p prods) GetAll(){
	for _,values :=range p{
		fmt.Println(values.Id, values.Name, values.Price, values.Category)
	}
}

func getById(id int)(ide int, message string){
	if id<=0{
		ide = id
		message = "El id ingresado no es correcto"
		return
	}

	for _, producto := range productos {
		if id==producto.Id{
			ide = id
			message = producto.Name
		} else {
			ide = id
			message = "No existe ese producto"
		}
	}
	return 
}

func main(){
	
	fmt.Println("Productos precargados")
	productos.GetAll()

	nuevoProducto := Product{
		Id: 987,
		Name: "Yerba",
		Price: 65.5,
		Description: "",
		Category: "Y",
	}

	fmt.Printf("\nNuevo listado de productos")
	nuevoProducto.Save()

	productos.GetAll()

	fmt.Printf("\n")
	otroProducto, mensaje := getById(45)
	otroProducto2, mensaje2 := getById(987)
	fmt.Println(otroProducto, mensaje)
	fmt.Println(otroProducto2, mensaje2)
}

