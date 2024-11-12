package book

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// AddBook agrega un nuevo libro a la lista
func AddBook() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Título del libro: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Autor del libro: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	fmt.Print("Año de publicación: ")
	var year int
	fmt.Scanf("%d\n", &year)

	fmt.Print("Género del libro: ")
	genre, _ := reader.ReadString('\n')
	genre = strings.TrimSpace(genre)

	// Crear el libro y añadirlo al slice
	book := Book{Title: title, Author: author, Year: year, Genre: genre}
	books = append(books, book)

	fmt.Println("Libro agregado exitosamente.")
}