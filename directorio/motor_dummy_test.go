package motor

import (
	"testing"
	// Importa la librería testify para realizar asserts en los tests
	"github.com/stretchr/testify/assert"
)

// Estructura DummyDB simula una base de datos para pruebas
type DummyDB struct{}

// Implementa los métodos de la interfaz DB pero devuelve valores vacíos o nil
func (d DummyDB) BuscarPorNombre(nombre string) string {
	return ""
}

func (d DummyDB) BuscarPorTelefono(telefono string) string {
	return ""
}

func (d DummyDB) AgregarEntrada(nombre, telefono string) error {
	return nil
}

// TestGetVersion prueba la función GetVersion del motor
func TestGetVersion(t *testing.T) {
	//arrange

	myDummyDB := DummyDB{}        // Crea una instancia Dummy de la base de datos simulada
	motor := NewEngine(myDummyDB) // Crea una instancia del motor con la base de datos simulada
	versionEsperada := "1.0"

	//act
	resultado := motor.GetVersion()

	//assert
	assert.Equal(t, versionEsperada, resultado)

}
