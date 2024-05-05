package motor

// repository será la capa que conectara el módulo de servicio con un almacen de datos.
// Qué con fines prácticos elegí trabajar sobre lo que
// Go define como struct, para replicar una db.
func (stg Storage) BuscarPorNombre(nombre string) (stgRespuesta Storage, err error) {

	// Aquí simulo un algo así como un puntero a la base de datos
	var myStgDB = StorageDB

	for _, registro := range myStgDB {
		if nombre == registro.Atributos.Nombre {
			stgRespuesta = registro
			break
		}
	}
	if stgRespuesta.Atributos.Nombre == "" {
		return Storage{}, ErrStorageRegistroNoEncontrado
	}
	return
}

func (stg *Storage) BuscarPorSector(sector string) (stgRespuesta Storage, err error) {

	var myStgDB = StorageDB

	for _, registro := range myStgDB {
		if sector == registro.Atributos.Sector {
			stgRespuesta = registro
		}
	}

	if stgRespuesta.Atributos.Nombre == "" {
		return Storage{}, ErrStorageRegistroNoEncontrado
	}
	return
}

func (stg *Storage) AgregarEntrada(nombre, sector string) error {
	// Aqui voy a omitir cualquier tipo de verificacion de error.
	stgNew := Storage{
		AtributosStorage{
			Nombre: nombre,
			Sector: sector,
		}, nil}
	StorageDB = append(StorageDB, stgNew)
	return nil
}

func (stg *Storage) BuscarTodos() ([]Storage, error) {
	// Aqui voy a omitir cualquier tipo de verificacion de error.
	return StorageDB, nil
}
