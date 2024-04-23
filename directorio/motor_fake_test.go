package motor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Estructura FakeSearchEngine simula una base de datos para pruebas con más control
type FakeSearchEngine struct {
	// Mapa para almacenar entradas nombre-telefono
	DB map[string]string
	// Para más info sobre el tipo: https://gobyexample.com/maps
}

// Busca una entrada por nombre en el mapa simulado
func (m *FakeSearchEngine) BuscarPorNombre(nombre string) string {
	return m.DB[nombre]
}

// Busca una entrada por teléfono en el mapa simulado, recorre todas las entradas
func (m *FakeSearchEngine) BuscarPorTelefono(telefono string) string {
	for key, value := range m.DB {
		if value == telefono {
			return key
		}
	}
	return ""
}

// Agrega una nueva entrada al mapa simulado
func (m *FakeSearchEngine) AgregarEntrada(nombre, telefono string) error {
	if m.DB == nil {
		m.DB = map[string]string{} // Inicializa el mapa si es nil
	}
	m.DB[nombre] = telefono
	return nil
}

// TestFindByNameFaked prueba la búsqueda por nombre con una base de datos Fake
func TestFindByNameFaked(t *testing.T) {
	// Arrange
	myFakeSearchEngine := FakeSearchEngine{} // Crea una instancia del motor de búsqueda simulado
	motor := NewEngine(&myFakeSearchEngine)  // Crea una instancia del motor con el motor simulado

	// Prepara datos de prueba
	testValues := map[string]string{"Carlos": "123456", "Hector": "234567"}
	for key, value := range testValues {
		motor.AddEntry(key, value)
	}

	// Act
	resultadoCarlos := motor.FindByName("Carlos")
	resultadoHector := motor.FindByName("Hector")
	resultadoTelefono := motor.FindByTelephone("123456")

	// Assert - Verificación de los resultados
	assert.Equal(t, testValues["Carlos"], resultadoCarlos)
	assert.Equal(t, testValues["Hector"], resultadoHector)
	assert.Equal(t, "Carlos", resultadoTelefono)
}
