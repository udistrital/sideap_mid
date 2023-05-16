package helpers

import (
	"sideap_mid/models"
	"sideap_mid/utils_oas/error_control"
	"sideap_mid/utils_oas/request"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

var pathGestorDocumental, _ = beego.AppConfig.String("gestorDocumentalService")
var endpointDocumento = "document/"
var endpointUpload = "upload"

func GetSoporte() (documento models.Response, file_ []models.File, outputError map[string]interface{}) {

	funcion := "GetSoporte - "
	defer error_control.ErrorControlFunction(funcion+"Unhandled Error!", "500")

	urlcrud := "http://" + pathGestorDocumental + endpointDocumento + "ad763b6e-debb-4c01-9c73-0da6c48b2de8"

	var file models.File
	err := request.Get(urlcrud, nil, &file)
	if err != nil {
		logs.Error(urlcrud+", ", err)
		eval := "request.Get(urlcrud, headers, &info)"
		outputError = error_control.Error(funcion+eval, err, "502")
	}

	urlcrud = "http://" + pathGestorDocumental + endpointDocumento + endpointUpload

	//
	file.IdTipoDocumento = 13
	file.Nombre = "jgcastellanosjcorreo.udistrital.edu.co"
	file.Descripcion = "jgcastellanosjcorreo.udistrital.edu.co"
	file.Metadatos = map[string]interface{}{"ss": "lalalal"}
	//

	file_ = append(file_, file)
	err = request.SendJson(urlcrud, "POST", &file_, &documento)
	if err != nil {
		logs.Error(urlcrud+", ", err)
		eval := `request.SendJson(urlcrud, "POST", &file, file)`
		outputError = error_control.Error(funcion+eval, err, "502")
	}

	return
}
