//lo que se hace aca es un modelo de ecommerce, como si fuera un almacen

package main

import (
	"errors"
	"fmt"
)

type Item struct{ //este struct contendra a los elementos que iremos añadiendo
	Title, Description string //otra forma de definir los campos de un struct
}

type ItemStore struct {
	Item
	Price, Quantity float64
}

//controlador 
type Shop struct{
	storage []ItemStore
	//se le puede agregar un limite de transacciones 
	transactions int
	limit 		 int 
}

//metodos 

//Sacar o meter articulos
func(s *Shop) CargaStock(items ...ItemStore)error{ //voy a ir restando o cargando el stock que vaya modificando. Apunto a shop para que se pueda modificar la variable
	
	if !s.Available(){
		return errors.New("not available")
	}

	s.storage = append(s.storage, items...)

	s.transactions ++

	return nil
}

//corroboracion de limite de transacciones 
func(s *Shop) Available() bool{
	return s.transactions < s.limit
}

func(s *Shop) IsEmpty() bool{
	return len(s.storage) == 0
}

func main(){
	
	storage := make([]ItemStore, 0)
	
	shop := &Shop{
		limit: 2,
		storage: storage,
	}	

	err:=shop.CargaStock(ItemStore{Item{"manzana",""}, 500,1})
	if err != nil{
		panic(err)
	}

	err=shop.CargaStock(ItemStore{Item{"manzana",""}, 500,1})
	if err != nil{
		panic(err)
	}

	err=shop.CargaStock(ItemStore{Item{"manzana",""}, 500,1})
	if err != nil{
		panic(err)
	}

	fmt.Println(shop)
}