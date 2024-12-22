package main

import (
	"book-management/handlers"
	"fmt"
	"net/http"
)

func main() {
	// Registrar rutas
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/books", handlers.HandleBooks) // CRUD para libros

	// Iniciar el servidor
	fmt.Println("Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
