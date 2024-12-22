package services

import (
	"book-management/models"
	"encoding/json"
	"errors"
	"os"
)

// Ruta del archivo JSON donde se almacenan los libros
const booksFile = "books.json"

// Servicio que implementa la interfaz BookStore
type BookService struct {
	books []models.Book
}

// NewBookService crea una nueva instancia de BookService
func NewBookService() *BookService {
	// Cargar los libros desde el archivo JSON
	bs := &BookService{}
	bs.loadBooks()
	return bs
}

// loadBooks carga los libros desde un archivo JSON
func (bs *BookService) loadBooks() {
	// Verificar si el archivo existe
	file, err := os.OpenFile(booksFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			// Si el archivo no existe, no hacer nada
			return
		}
		// Si ocurre un error al abrir el archivo, reportar el error
		panic("Error al abrir el archivo de libros: " + err.Error())
	}
	defer file.Close()

	// Leer el contenido del archivo
	data := make([]byte, 0)
	_, err = file.Read(data)
	if err != nil && err.Error() != "EOF" {
		panic("Error al leer el archivo de libros: " + err.Error())
	}

	// Deserializar el contenido JSON en la variable books
	if len(data) > 0 {
		if err := json.Unmarshal(data, &bs.books); err != nil {
			panic("Error al deserializar el archivo de libros: " + err.Error())
		}
	}
}

// saveBooks guarda los libros en un archivo JSON
func (bs *BookService) saveBooks() error {
	data, err := json.MarshalIndent(bs.books, "", "  ")
	if err != nil {
		return errors.New("Error al serializar los libros: " + err.Error())
	}

	// Guardar los datos en el archivo
	err = os.WriteFile(booksFile, data, 0644)
	if err != nil {
		return errors.New("Error al guardar los libros en el archivo: " + err.Error())
	}

	return nil
}

// AddBook agrega un libro al servicio y guarda los cambios en el archivo
func (bs *BookService) AddBook(book models.Book) error {
	if book.Title == "" || book.Author == "" || book.Year <= 0 {
		return errors.New("datos inválidos para el libro")
	}
	bs.books = append(bs.books, book)

	// Guardar los libros actualizados en el archivo
	if err := bs.saveBooks(); err != nil {
		return err
	}
	return nil
}

// ListBooks devuelve todos los libros
func (bs *BookService) ListBooks() ([]models.Book, error) {
	return bs.books, nil
}

// SearchBook busca un libro por su título
func (bs *BookService) SearchBook(title string) (*models.Book, error) {
	for _, b := range bs.books {
		if b.Title == title {
			return &b, nil
		}
	}
	return nil, errors.New("libro no encontrado")
}

// DeleteBook elimina un libro por su título y guarda los cambios en el archivo
func (bs *BookService) DeleteBook(title string) error {
	for i, b := range bs.books {
		if b.Title == title {
			bs.books = append(bs.books[:i], bs.books[i+1:]...)

			// Guardar los libros actualizados en el archivo
			if err := bs.saveBooks(); err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("libro no encontrado para eliminar")
}
