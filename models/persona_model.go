package models

type Persona struct {
	ID       int    `json:"id_persona" db:"id_persona"`
	Apellido string `json:"apellido"   db:"apellido"`
	Celular  string `json:"celular"    db:"celular"`
	Name     string `json:"nombre"     db:"nombre"`
}
