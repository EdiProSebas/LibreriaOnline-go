package main

import (
	"book-management/book"
	"book-management/menu"
	"fmt"
)

func main() {
	for {
		menu.ShowMenu()

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			book.AddBook()
		case 2:
			book.ListBooks()
		case 3:
			book.SearchBook()
		case 4:
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}