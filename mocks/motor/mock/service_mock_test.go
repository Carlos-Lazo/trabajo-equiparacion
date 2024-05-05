package motor

import (
	"testing"
	"trabajo-equiparacion/carlos-lazo/test-dobles/internal/motor"

	"github.com/stretchr/testify/assert"
)

/********************************************************************************************/
// finalmente, en mock. el cual presenta una combinacion de todas las caracteristicas
// de los doubles anteriores. Lo que quiere decir, entre otras cosas, que con un mock
// pordria implementar analogamente a un dummy, así como también una implementación
// más compleja, como la que desarrollaremos aqui en los ejemplos.
/********************************************************************************************/
// a tener en cuenta, la implementacion de mock hará que las funciones,
// naturalmente, se vuelvan más complejas. lo que lo hace una desventaja
// por otro lado, obtenemos ganancias de poner simular componentes
// para el test que son más complejos, y ganamos legitividad,
// confianza y seguridad en nuestros tests.
// ya que cada vez implementarmos objetos
// cada vez más parecidos a los reales.
/********************************************************************************************/

func TestBuscarPorNombre(t *testing.T) {

	t.Run("mock - error por longitud insuficiente", func(t *testing.T) {
		//
		myMock := NewMockObj()
		myMotor := motor.Motor{}
		nombreBuscado := "nv"
		// aqui podemos ver la facilidad y potencia de la implementacion de un mock
		myMock.On("BuscarPorNombre", nombreBuscado).Return(motor.Storage{}, "longitud insuficiente")
		//
		_, errResultado := myMotor.BuscarPorNombre(nombreBuscado)
		_, esperado := myMock.BuscarPorNombre(nombreBuscado)
		//
		assert.EqualError(t, errResultado, esperado)

	})
}

func TestAgregarEntrada(t *testing.T) {
	t.Run("mock - entrada agregada exitosamente", func(t *testing.T) {
		//
		myMock := NewMockObj()
		myMotor := motor.Motor{}
		// entrada de nombre valida
		nombreEntrada := "valid"
		sectorEntrada := "4J"
		// podria manipular los argumentos desde la definicion.
		// es decir, podria no usarlos.
		myMock.On("AgregarEntrada", nombreEntrada, sectorEntrada).Return(nil)
		//
		resultado := myMotor.AgregarEntrada(nombreEntrada, sectorEntrada)
		esperado := myMock.AgregarEntrada(nombreEntrada, sectorEntrada)
		//
		assert.Equal(t, resultado, esperado)
	})
}
