package models

import "time"

type HojaVida struct {
	ID                        int                  `json:"idPersonaSIDEAP"`
	SoporteDocumentoIdentidad string               `json:"soporteDocumentoIdentidad"`
	SoporteEducacionBasica    string               `json:"soporteEducacionBasica"`
	EducacionSuperior         []EducacionSuperior  `json:"educacionSuperior"`
	ExperienciaLaboral        []ExperienciaLaboral `json:"experienciaLaboral"`
	OtroEstudio               []OtroEstudio        `json:"otroEstudio"`
}

type Estudio struct {
	TipoRegistroID int `json:"idTipoRegistro"`
	RegistroID     int `json:"idRegistro"`
}

type EducacionSuperior struct {
	Estudio
	Programa               string `json:"programa"`
	UltimoSemestreAprobado int    `json:"ultimoSemestreAprobado"`
	Graduado               bool   `json:"graduado"`
	Soporte                string `json:"soporteEstudio"`
}

type ExperienciaLaboral struct {
	Estudio
	Entidad      string    `json:"entidad"`
	FechaIngreso time.Time `json:"fechaIngreso"`
	FechaRetiro  time.Time `json:"fechaRetiro"`
	Cargo        string    `json:"cargoOContrato"`
	Soporte      string    `json:"soporte"`
}

type OtroEstudio struct {
	Estudio
	Programa      string `json:"programa"`
	NombreEstudio string `json:"nombreEstudio"`
	Institucion   bool   `json:"institucionFormacion"`
	Soporte       string `json:"soporte"`
}
