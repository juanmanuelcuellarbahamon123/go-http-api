package models

type Persona struct {
	ID       int    `json:"id_persona" db:"id_persona"`
	Nombre   string `json:"nombre" db:"nombre"`
	Apellido string `json:"apellido" db:"apellido"`
	Celular  string `json:"celular" db:"celular"`
}
