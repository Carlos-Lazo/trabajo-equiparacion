package motor

import (
	"trabajo-equiparacion/carlos-lazo/test-dobles/internal/motor"

	"github.com/stretchr/testify/mock" // Mock tiene su propia emplementacion
)

type AtributosMock struct {
	Nombre string
	Sector string
}

func NewMockObj() *MockObj {
	return &MockObj{}
}

type MockObj struct {
	mock.Mock
}

// estas funciones las utilizará el mock, como plantillas, cuando haga
// un llamado el método on. por ejemplo, 'mock.On("BuscarTodos)'
// y devolvera los argumentos acorde a la configuracion
// definida en las funciones siguientes.

func (mrt *MockObj) BuscarPorNombre(nombre string) (stg motor.Storage, err string) {
	// Estás implementaciones actuan muy acorde de la implementacion real
	// A su ves son muy sensillas de aplicar, ya que todo es abstracto
	// para quien realiza el test. Lo deja concentrarce en el test.
	args := mrt.Called(nombre)
	return args.Get(0).(motor.Storage), args.Get(1).(string)
}

func (mrt *MockObj) BuscarPorSector(sector string) (stg motor.Storage, err string) {
	args := mrt.Called(sector)
	return args.Get(0).(motor.Storage), args.Get(1).(string)
}
func (mrt *MockObj) AgregarEntrada(nombre, sector string) error {
	args := mrt.Called(nombre, sector)
	return args.Error(0)
}

func (mrt *MockObj) BuscarTodos() ([]motor.Storage, error) {
	args := mrt.Called()
	return args.Get(0).([]motor.Storage), args.Error(1)
}
