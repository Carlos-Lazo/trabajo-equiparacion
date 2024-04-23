package motor

// Interfaz que define las operaciones básicas para la base de datos
type DB interface {
	BuscarPorNombre(nombre string) string
	BuscarPorTelefono(telefono string) string
	AgregarEntrada(nombre, telefono string) error
}
