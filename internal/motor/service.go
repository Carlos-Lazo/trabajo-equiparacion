package motor

import "fmt"

// Este es el paquete nativo de Go para implementar un standard output

// Estructura que representa el Motor de directorios
type Motor struct {
	// Dependencia de la interfaz DB para operaciones con la base de datos
	database Storage
}

// Crea una nueva instancia del Motor de directorios
//:: Esta es la forma en la que Go implementa los constructores
func (mtr Motor) NewMotor(db Storage) Motor {
	// Inicializa el Motor con la base de datos proporcionada
	return Motor{database: db}
}

// Obtiene la versión del Motor
func (mtr Motor) Version() string {
	return "1.0"
}

// Busca una entrada por nombre (con validación mínima de longitud)
func (mtr Motor) BuscarPorNombre(nombre string) (stgRespuesta Storage, err error) {
	if len(nombre) < 3 {
		return Storage{}, ErrLongitudInsuficiente
	}
	stgRespuesta, err = mtr.database.BuscarPorNombre(nombre)
	if err != nil {
		return Storage{}, ErrStorageRegistroNoEncontrado
	}

	return

}

// Agrega una nueva entrada al directorio (con validación de longitud)
func (mtr Motor) AgregarEntrada(nombre, sector string) (err error) {
	if len(nombre) < 3 || len(sector) < 2 {
		return ErrLongitudInsuficiente
	}

	err = mtr.database.AgregarEntrada(nombre, sector)
	if err != nil {
		return ErrStorageInterno
	}
	return
}

func (mtr Motor) BuscarTodos() ([]Storage, error) {

	return mtr.database.BuscarTodos()
}

func (mtr Motor) VistaListarRegistros() error {

	stgRegistros, err := mtr.BuscarTodos()

	if err != nil {
		return err
	}

	for _, registro := range stgRegistros {
		MostrarRegistro(registro)
	}

	return nil

}

func MostrarRegistro(rgt Storage) {

	fmt.Println("nombre:", rgt.Atributos.Nombre)
	fmt.Println("sector:", rgt.Atributos.Sector)
}
