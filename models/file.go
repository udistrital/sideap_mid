package models

type File struct {
	Descripcion     string `json:"descripcion"`
	File            string `json:"file"`
	IdTipoDocumento int
	Metadatos       map[string]interface{} `json:"metadatos"`
	Nombre          string                 `json:"nombre"`
}

type Documento struct {
	Activo      bool
	Descripcion string
	Enlace      string
	Id          int
	Metadatos   string
	Nombre      string
}
