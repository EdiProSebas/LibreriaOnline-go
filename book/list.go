package book

import "fmt"

// ListBooks muestra todos los libros en la lista
func ListBooks() {
	if len(books) == 0 {
		fmt.Println("No hay libros en la lista.")
		return
	}

	fmt.Println("\n--- Lista de Libros ---")
	for i, book := range books {
		fmt.Printf("%d. Título: %s | Autor: %s | Año: %d | Género: %s\n", i+1, book.Title, book.Author, book.Year, book.Genre)
	}
}