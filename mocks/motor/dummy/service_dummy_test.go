package motor

import (
	"testing"
	"trabajo-equiparacion/carlos-lazo/test-dobles/internal/motor"

	"github.com/stretchr/testify/assert"
)

/********************************************************************************************/
// Dummy es un objeto que no tiene ninguna implementación real.
// Se utiliza para reemplazar dependencias que no son relevantes para la prueba que se está realizando.
/********************************************************************************************/
// Las funciones que vamos a testear pertenecen a la capa service del proyecto.
// Por lo tanto, nuestros doubles estaran basados
// en los componentes de la capa repository.
/********************************************************************************************/

type DummyObj struct {
	versionEsperada string
} // No deberia tener más campos para este caso.

// TestGetVersion prueba la función GetVersion del motor
func TestGetVersion(t *testing.T) {
	t.Run("Exitosa obtencion de version", func(t *testing.T) {
		//
		myDummyObj := DummyObj{} // Crea una instancia Dummy
		// Elijo el valor para mi test, recreando como si fuera la función
		// sin importar, sí necesita, tiene o tuvo procesamiento. Entre otras cosas.
		myDummyObj.versionEsperada = "1.0"
		// Es el ejemplo anterior se puede notar la simpleza y ventaja de usar un Dummy.
		// Podemos notar que para este caso, la suplantación la está haciendo
		// a una función propia.
		motor := motor.Motor{}
		//
		resultado := motor.Version()
		versionEsperada := myDummyObj.versionEsperada
		//
		assert.Equal(t, versionEsperada, resultado)
	})
}

// TestBuscarPorNombre prueba la función BuscarPorNombre del motor
func TestBuscarPorNombre(t *testing.T) {
	t.Run("Dummy - Exitosa obtencion de nombre", func(t *testing.T) {
		//
		// Para este ejemplo omitire el modo del test anterior.
		// Ya que este podria tener una resolucion similir.
		// Un dummy tambien podria ver así
		nombreEsperado := "Carlos"
		// Ya que cumple tanto con la definición de dummy
		// Y las necesidades del test.
		// Aquí notamos, caso distinto al anterior
		// donde la suplantación la está haciendo sobre
		// resultados de una función de una capa externa.
		nombreBuscado := "Carlos"
		motor := motor.Motor{}
		//
		registro, _ := motor.BuscarPorNombre(nombreBuscado)
		resultado := registro.Atributos.Nombre
		//
		assert.Equal(t, nombreEsperado, resultado)
	})
}

func TestAgregarEntrada(t *testing.T) {
	//
	motor := motor.Motor{}
	nombreNuevo := "Carlos"
	sectorNuevo := "7G"
	//
	resultado := motor.AgregarEntrada(nombreNuevo, sectorNuevo)
	// Por último.
	// Este segundo argumento también podria ser un dummy.
	// Si uno lo mira desde la perspectiva de solo firmar
	// para cumplir con los argumentos de una función
	// o solo con un valor, como este es el caso.
	// Lo anterior, es lo que hizo basicamente
	// en los dos ejemplos anteriores
	// solo que de otra manera.
	assert.Equal(t, nil, resultado)
}
