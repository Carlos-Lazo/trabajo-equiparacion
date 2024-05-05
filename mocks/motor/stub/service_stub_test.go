package motor

import (
	"testing"
	"trabajo-equiparacion/carlos-lazo/test-dobles/internal/motor"

	"github.com/stretchr/testify/assert"
)

/********************************************************************************************/
// stub se presenta como una mejora del tipo anterior, el fake
// estas dichas mejoras son evidenciadas en la complejidad
// y el alcance de la implementacion del objeto.
/********************************************************************************************/
// Se puede notar que las implementaciones de los tipos
// doubles se incrementa.
// lo que produce que la implementacion del double
// se pareza aún más a lo que se quiere simular
/********************************************************************************************/

// TestBuscarPorNombre prueba la búsqueda por nombre con un stub
func TestBuscarPorNombre(t *testing.T) {

	t.Run("stub - obtencion exitosa de nombre", func(t *testing.T) {

		//
		// crea una instancia del stub simular las funciones del repository
		myStubObj := StubObj{StubAtributos: AtributosStub{Nombre: "Carlos", Sector: "7G"}}
		motor := motor.Motor{}

		nombreBuscado := "Carlos"

		// Act
		resultado, _ := motor.BuscarPorNombre(nombreBuscado)
		actual, _ := myStubObj.BuscarPorNombre(nombreBuscado)
		// Assert
		assert.Equal(t, resultado.Atributos.Nombre, actual.StubAtributos.Nombre)
	})
	t.Run("stub - error por longitud insuficiente", func(t *testing.T) {

		//
		// crea una instancia del stub simular las funciones del repository
		myStubObj := StubObj{}
		motor := motor.Motor{}

		nombreBuscado := "nv"

		// Act
		_, errResultado := motor.BuscarPorNombre(nombreBuscado)
		_, errActual := myStubObj.BuscarPorNombre(nombreBuscado)
		// Assert
		// assert.Equal(t, resultado, actual)
		assert.EqualError(t, errResultado, errActual)

	})

}
