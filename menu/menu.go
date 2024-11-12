package menu

import "fmt"

// ShowMenu muestra el menú principal
func ShowMenu() {
	fmt.Println("\n--- Sistema de Gestión de Libros ---")
	fmt.Println("1. Agregar un libro")
	fmt.Println("2. Ver lista de libros")
	fmt.Println("3. Buscar libro por título")
	fmt.Println("4. Salir")
	fmt.Print("Seleccione una opción: ")
}