package models

type HojaVida struct {
	Persona []HojaVida_
}

type HojaVida_ struct {
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
	Entidad      string `json:"entidad"`
	FechaIngreso string `json:"fechaIngreso"`
	FechaRetiro  string `json:"fechaRetiro"`
	Cargo        string `json:"cargoOContrato"`
	Soporte      string `json:"soporte"`
}

type OtroEstudio struct {
	Estudio
	Programa      string `json:"programa"`
	NombreEstudio string `json:"nombreEstudio"`
	Institucion   string `json:"institucionFormacion"`
	Soporte       string `json:"soporte"`
}
