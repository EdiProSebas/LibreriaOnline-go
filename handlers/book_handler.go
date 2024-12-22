package handlers

import (
	"book-management/models"
	"book-management/services"
	"encoding/json"
	"net/http"
)

var bookService = services.NewBookService()

// HandleBooks maneja las rutas relacionadas con libros
func HandleBooks(w http.ResponseWriter, r *http.Request) {
	// Asegúrate de que se usen los parámetros
	if r == nil || w == nil {
		return
	}

	switch r.Method {
	case http.MethodPost:
		handleAddBook(w, r)
	case http.MethodGet:
		handleListBooks(w, r)
	case http.MethodDelete:
		handleDeleteBook(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// Agregar un libro
func handleAddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
		return
	}

	if err := bookService.AddBook(book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Libro agregado exitosamente"})
}

// Listar libros
func handleListBooks(w http.ResponseWriter, _ *http.Request) {
	books, err := bookService.ListBooks()
	if err != nil {
		http.Error(w, "Error al obtener los libros", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Eliminar un libro
func handleDeleteBook(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Se requiere un título", http.StatusBadRequest)
		return
	}

	if err := bookService.DeleteBook(title); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Libro eliminado exitosamente"})
}
