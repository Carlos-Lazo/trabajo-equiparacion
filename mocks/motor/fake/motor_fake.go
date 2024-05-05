package motor

import "errors"

type AtributosFake struct {
	Nombre string
	Sector string
}

type FakeObj struct {
	// simulo de manera más simple
	// el comportamiento del componente
	// storageDB de la capa repository.
	FakeStorage AtributosFake
}

func (m *FakeObj) BuscarPorNombre(nombre string) string {
	// implementación simplificada
	return nombre // Podria tambien devolver un "nombre"
}

func (m *FakeObj) AgregarEntradaExt(nombre string, sector string) error {
	// esto es un ejemplo de las caracteristicas de un fake
	return nil
}

func (m *FakeObj) AgregarEntradaErr(nombre string, sector string) error {
	// un logica simple.
	return errors.New("longitud insuficiente")
}
