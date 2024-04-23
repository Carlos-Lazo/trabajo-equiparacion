package motor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Estructura StubDB simula una base de datos para pruebas con comportamiento predefinido
type StubDB struct{}

func (d StubDB) BuscarPorNombre(nombre string) string {
	return "12345678" // Siempre devuelve un valor predefinido
}

func (d StubDB) BuscarPorTelefono(telefono string) string {
	return "" // No implementado para esta prueba
}

func (d StubDB) AgregarEntrada(nombre, telefono string) error {
	return nil // No implementado para esta prueba
}

// TestFindByName prueba la búsqueda por nombre con un stub
func TestFindByName(t *testing.T) {
	// Arrange
	myStubDB := StubDB{}         // Crea una instancia del stub de la base de datos
	motor := NewEngine(myStubDB) // Crea una instancia del motor con el stub

	telefonoEsperado := "12345678" // Define el teléfono esperado (valor predefinido del stub)

	// Act
	resultado := motor.FindByName("Carlos") // Busca por "Carlos"

	// Assert
	assert.Equal(t, telefonoEsperado, resultado)
}

// TestFindByNameWithNameShorterThan3 prueba que la búsqueda por nombre corto devuelve cadena vacía
func TestFindByNameWithNameShorterThan3(t *testing.T) {
	// Arrange
	myStubDB := StubDB{}         // Crea una instancia del stub de la base de datos
	motor := NewEngine(myStubDB) // Crea una instancia del motor con el stub

	telefonoEsperado := "" // Define el teléfono esperado (cadena vacía)

	// Act
	resultado := motor.FindByName("Hec") // Busca por "Ah" (nombre corto)

	// Assert
	assert.Equal(t, telefonoEsperado, resultado)
}
