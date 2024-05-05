package motor

import (
	"errors"
)

/********************************************************************************************/
// Interfaz que define las operaciones básicas para la base de datos
type RepositoryStorage interface {
	BuscarTodos() ([]Storage, error)
	BuscarPorNombre(nombre string) (string, error)
	BuscarPorSector(sector string) (string, error)
	AgregarEntrada(nombre, sector string) error
}

/********************************************************************************************/
// Estructura para simular la db.
// Estructura muy simple para definir los atributos del dominio Motor,
// la cual intenta replicar los atributos una tupla de una base de datos.
type AtributosStorage struct {
	Nombre string
	Sector string
}

type Storage struct {
	Atributos AtributosStorage
	repo      RepositoryStorage
}

/********************************************************************************************/
func NewStorage(repo RepositoryStorage) *Storage {
	return &Storage{
		repo: repo,
	}
}

/********************************************************************************************/
// Datos ficticios
var StorageDB = []Storage{{

	Atributos: AtributosStorage{
		Nombre: "Carlos", Sector: "7G"},
}, {

	Atributos: AtributosStorage{
		Nombre: "srt", Sector: "18F"},
}, {

	Atributos: AtributosStorage{
		Nombre: "Richard", Sector: "11S"},
}}

/********************************************************************************************/
var (
	// ErrStorageInterno en caso de ocurrir un error interno en la capa de la base de datos
	ErrStorageInterno = errors.New("error interno")

	// ErrLongitudInsuficiente en caso de que la entrada no cumpla con la longitud.
	ErrLongitudInsuficiente = errors.New("longitud insuficiente")

	// ErrStorageRegistroNoEncontrado en caso de que un registro del storage no sea encontrado.
	// Ya sea por no coincidir algunos de los atributos.
	ErrStorageRegistroNoEncontrado = errors.New("registro no encontrado")
)

/********************************************************************************************/
