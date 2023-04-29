package models

type InformacionBasica struct {
	Persona []Persona
}

type Persona struct {
	ID                    int    `json:"idPersonaSIDEAP"`
	TipoDocumento         string `json:"idTipoDocumento"`
	NumeroDocumento       string `json:"numDocumento"`
	PrimerNombre          string `json:"primerNom"`
	SegundoNombre         string `json:"otroNom"`
	PrimerApellido        string `json:"primerApellido"`
	SegundoApellido       string `json:"segundoApellido"`
	Genero                string `json:"codSexoPersona"`
	PaisNacimientoID      int    `json:"idPaisNacimiento"`
	MunicipioNacimientoID string `json:"idMunicipioNacimiento"`
	FechaNacimiento       string `json:"fechaNacimiento"`
	LibretaMilitar        bool   `json:"tieneLibretaMilitar"`
	DireccionResidencia   string `json:"direccionResidencia"`
	Telefono              string `json:"numTelefonoDomicilio"`
	EmailPersonal         string `json:"correoElectronicoPersonal"`
}
