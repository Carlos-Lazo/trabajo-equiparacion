package motor

import (
	"testing"
	"trabajo-equiparacion/carlos-lazo/test-dobles/internal/motor"

	"github.com/stretchr/testify/assert"
)

/********************************************************************************************/
// Fake es un tipo de Test Double que se utiliza para remplazar.
// funcionalidades básicas de un componente que es demandado
// mientras transcurre el test.
/********************************************************************************************/
// con la misma aclaración que con los tipo Dummy, sobre
// la naturaleza de los doubles. Para esta ocasión
// y para todos los demás tipos de doubles.
/********************************************************************************************/

// TestBuscarPorNombre prueba la búsqueda por nombre implementando el double Fake.
func TestBuscarPorNombre(t *testing.T) {
	t.Run("Fake - Exitosa obtencion de nombre", func(t *testing.T) {
		//
		myFakeRepository := FakeObj{} // crea una instancia del Storage simulado
		motor := motor.Motor{}
		// situacion que involucra una estructura un tanto más compleja
		// que en el ejemplo anteriores con el Dummy.
		// nombre esperado
		myFakeRepository.FakeStorage = AtributosFake{Nombre: "Carlos", Sector: "7G"}
		nombreBuscado := "Carlos"
		//
		resultado, _ := motor.BuscarPorNombre(nombreBuscado)
		//
		assert.Equal(t, resultado.Atributos.Nombre, myFakeRepository.FakeStorage.Nombre)
	})
}

func TestAgregarEntrada(t *testing.T) {

	t.Run("Fake - Error de AgregarEntrada por longitud insuficiente", func(t *testing.T) {
		//
		// Entrada de nombre invalida
		nombreEntrada := "nv"
		sectorEntrada := "7I"
		motor := motor.Motor{}
		fakeMotor := FakeObj{}
		//
		resultado := motor.AgregarEntrada(nombreEntrada, sectorEntrada)
		fakeResultado := fakeMotor.AgregarEntradaErr(nombreEntrada, sectorEntrada)
		//
		assert.Equal(t, resultado, fakeResultado)
	})

	t.Run("Fake - AgregarEntrada exitosa", func(t *testing.T) {
		//
		// Entrada de nombre valida
		nombreEntrada := "valid"
		sectorEntrada := "4J"
		motor := motor.Motor{}
		fakeMotor := FakeObj{}
		//
		resultado := motor.AgregarEntrada(nombreEntrada, sectorEntrada)
		fakeResultado := fakeMotor.AgregarEntradaExt(nombreEntrada, sectorEntrada)
		//
		assert.Equal(t, resultado, fakeResultado)
	})
}
