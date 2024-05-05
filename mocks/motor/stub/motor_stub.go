package motor

type AtributosStub struct {
	Nombre string
	Sector string
}

type StubObj struct {
	// simulo de manera más simple
	// el comportamiento del componente
	// storageDB de la capa repository.
	StubAtributos AtributosStub
}

func (stb StubObj) BuscarPorNombre(nombre string) (StubObj, string) {
	if len(nombre) < 3 { // implementacion de logica para el stub
		return StubObj{}, "longitud insuficiente"
	}
	// con esto, la misma funcion puede cubrir
	// mas de un caso de retorno
	// para el testing
	return stb, "" // devuelve un stub
}
