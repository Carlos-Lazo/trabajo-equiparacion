package motor

import "fmt" // Este es el paquete nativo de Go para implementar un standard output

// Estructura que representa el motor de directorios
type engine struct {
	// Dependencia de la interfaz DB para operaciones con la base de datos
	database DB
}

// Crea una nueva instancia del motor de directorios
//:: Esta es la forma en la que Go implementa los constructores
func NewEngine(db DB) engine {
	// Inicializa el motor con la base de datos proporcionada
	return engine{database: db}
}

// Obtiene la versión del motor
func (e engine) GetVersion() string {
	return "1.0"
}

// Busca una entrada por nombre (con validación mínima de longitud)
func (e engine) FindByName(name string) string {
	if len(name) > 3 {
		return e.database.BuscarPorNombre(name)
	}
	// Nombre demasiado corto, devuelve cadena vacía
	return ""
}

// Busca una entrada por teléfono (con validación mínima de longitud)
func (e engine) FindByTelephone(telephone string) string {
	if len(telephone) > 5 {
		return e.database.BuscarPorTelefono(telephone)
	}
	// Teléfono demasiado corto, devuelve cadena vacía
	return ""
}

// Agrega una nueva entrada al directorio (con validación de longitud)
func (e engine) AddEntry(name, telephone string) error {
	if len(name) > 3 && len(telephone) > 5 {
		return e.database.AgregarEntrada(name, telephone)
	}
	// Nombre o teléfono demasiado cortos, devuelve error personalizado
	return fmt.Errorf("name %s and telephone %s cannot be added to the registry", name, telephone)
}
