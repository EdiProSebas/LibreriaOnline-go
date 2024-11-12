package book

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// SearchBook busca un libro por su título en la lista
func SearchBook() {
	if len(books) == 0 {
		fmt.Println("No hay libros en la lista para buscar.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el título del libro a buscar: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("\n--- Resultados de la Búsqueda ---")
	found := false
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
			fmt.Printf("Título: %s | Autor: %s | Año: %d | Género: %s\n", book.Title, book.Author, book.Year, book.Genre)
			found = true
		}
	}

	if !found {
		fmt.Println("No se encontraron libros con ese título.")
	}
}