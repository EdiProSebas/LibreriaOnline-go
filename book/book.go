package book

// Book representa la estructura de un libro en el sistema
type Book struct {
	Title  string
	Author string
	Year   int
	Genre  string
}

// books es una lista de libros almacenada en memoria
var books []Book