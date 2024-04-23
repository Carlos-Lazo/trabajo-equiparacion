package motor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Estructura MockDB simula una base de datos para pruebas con control de llamadas
type MockDB struct {
	// Bandera para verificar si se llamó a BuscarPorNombre
	BuscarPorNombreWasCalled bool // Agrego 'WasCalled' para no generar ambiguedad
}

// Busca una entrada por nombre en la base de datos simulada y marca la bandera
func (m *MockDB) BuscarPorNombre(nombre string) string {
	m.BuscarPorNombreWasCalled = true
	return "12345678" // Valor predefinido para la prueba
}

func (m *MockDB) BuscarPorTelefono(telefono string) string {
	return "" // No implementado para esta prueba
}

func (m *MockDB) AgregarEntrada(nombre, telefono string) error {
	return nil // No implementado para esta prueba
}

// TestFindByNameMocked prueba la búsqueda por nombre con un mock
func TestFindByNameMocked(t *testing.T) {
	// Arrange
	myMockDB := MockDB{}          // Crea una instancia del mock de la base de datos
	motor := NewEngine(&myMockDB) // Crea una instancia del motor con el mock

	telefonoEsperado := "12345678"

	// Act
	resultado := motor.FindByName("Carlos")

	// Assert
	assert.Equal(t, telefonoEsperado, resultado)
	assert.True(t, myMockDB.BuscarPorNombreWasCalled)
}

// TestFindByNameMockedNotCalled prueba que BuscarPorNombre no se llama para nombres cortos
func TestFindByNameMockedNotCalled(t *testing.T) {
	// Arrange
	myMockDB := MockDB{}          // Crea una instancia del mock de la base de datos
	motor := NewEngine(&myMockDB) // Crea una instancia del motor con el mock

	telefonoEsperado := "" // Define el teléfono esperado (cadena vacía)

	// Act
	resultado := motor.FindByName("Car") // Busca por "Car" (nombre corto)

	// Assert
	assert.Equal(t, telefonoEsperado, resultado)       // Verifica si se obtuvo la cadena vacía
	assert.False(t, myMockDB.BuscarPorNombreWasCalled) // Verifica si NO se llamó a BuscarPorNombre
}
