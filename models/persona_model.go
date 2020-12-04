package models

type Persona struct {
	Nombre   string `json:nombre`
	Apellido string `json:apellido`
	Celular  string `json:celular`
}
