package models

// Book representa un libro en el sistema
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

// BookStore define las operaciones que se pueden realizar sobre los libros
type BookStore interface {
	AddBook(book Book) error
	ListBooks() ([]Book, error)
	SearchBook(title string) (*Book, error)
	DeleteBook(title string) error
}
