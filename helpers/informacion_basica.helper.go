package helpers

import (
	"sideap_mid/models"
	"sideap_mid/utils_oas/error_control"
	"sideap_mid/utils_oas/request"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

var path, _ = beego.AppConfig.String("sideapService")
var key, _ = beego.AppConfig.String("sideapKey")
var usuario, _ = beego.AppConfig.String("sideapUser")
var endpointBasicos = "obtenerDatosBasicosUDISTRITAL"

func GetInformacionBasicaByDocumento(codigoDocumento string, numeroDocumento string) (info models.InformacionBasica, outputError map[string]interface{}) {

	funcion := "GetInformacionBasicaByDocumento - "
	defer error_control.ErrorControlFunction(funcion+"Unhandled Error!", "500")

	headers := map[string]string{
		"key":             key,
		"usuario":         usuario,
		"tipoDocumento":   codigoDocumento,
		"numeroDocumento": numeroDocumento,
	}

	urlcrud := "http://" + path + endpointBasicos

	err := request.Get(urlcrud, headers, &info)
	if err != nil {
		logs.Error(urlcrud+", ", err)
		eval := "request.Get(urlcrud, headers, &info)"
		outputError = error_control.Error(funcion+eval, err, "502")
	}

	return
}
