package main

import (
	"fmt"
	"log"
)

type book struct {
	Id     int
	Title  string
	Author string
	Price  float64
}

var bookArray = []book{
	{
		Id:     1,
		Title:  "Clean architecture",
		Author: "Robert Martin",
		Price:  150.00,
	},
	{
		Id:     2,
		Title:  "Learning go",
		Author: "Jon Bodner",
		Price:  80.00,
	},
}

func main() {
	fmt.Println("Bienvenido visitante")
	menu()

	isActive := true

	for isActive {

		var options string
		_, err := fmt.Scan(&options)
		exceptions(err)

		switch options {
		case "0":
			menu()
		case "1":
			showBookCatalog()
		case "2":
			var quantity int
			_, err := fmt.Scan(&quantity)
			exceptions(err)
			registerBooks(quantity)
		case "3":
			var id int
			_, err := fmt.Scan(&id)
			exceptions(err)
			searchBookById(id)
		case "9":
			isActive = false
			goodbye()
		default:
			fmt.Println("Si quiere terminar la aplicacion digite {9}")
		}

	}
}

func exceptions(err error) {
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error:: ", err)
	}
}

func registerBooks(quantity int) {
	for i := 1; i <= quantity; i++ {
		var b book
		b.Id = 2 + i

		fmt.Println("1.\t Ingrese el titulo")
		fmt.Scan(&b.Title)

		fmt.Println("2.\t Ingrese el autor")
		fmt.Scan(&b.Author)

		fmt.Println("3.\t Ingrese el precio")
		fmt.Scan(&b.Price)

		bookArray = append(bookArray, b)

		fmt.Println("Guardado con codigo:: ", b.Id)
	}
}

func showBookCatalog() {

	for _, ba := range bookArray {
		fmt.Printf("Id: %d \ntitle: %s \nauthor: %s \n", ba.Id, ba.Title, ba.Author)
	}
}

func searchBookById(id int) {
	for _, ba := range bookArray {
		if id == ba.Id {
			fmt.Printf("Id:%d, \ntitle:%s, \nauthor:%s \n", ba.Id, ba.Title, ba.Author)
		}
	}

}

func menu() {
	fmt.Println("Digite alguna de las opciones:")
	fmt.Println("{0} Ver el menu de opciones")
	fmt.Println("{1} Ver catalogo de libros")
	fmt.Println("{2} Registrar un nuevo libro")
	fmt.Println("{3} Buscar libro por id")
	fmt.Println("{9} Cerrar")
}

func goodbye() {
	fmt.Println("Hasta luego")
}
