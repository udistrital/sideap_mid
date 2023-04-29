package helpers

import (
	"fmt"
	"sideap_mid/models"
	"sideap_mid/utils_oas/error_control"
	"sideap_mid/utils_oas/request"

	"github.com/beego/beego/v2/core/logs"
)

var endpointHojaVida = "obtenerHojaVidaUDISTRITAL"

func GetHojaVida(personaId int, codigoDocumento string, numeroDocumento string) (hojaVida models.HojaVida, outputError map[string]interface{}) {

	funcion := "GetHojaVida - "
	defer error_control.ErrorControlFunction(funcion+"Unhandled Error!", "500")

	if personaId == 0 {
		info, outputError_ := GetInformacionBasicaByDocumento(codigoDocumento, numeroDocumento)
		if outputError_ != nil || len(info.Persona) == 0 || info.Persona[0].ID == 0 {
			outputError = outputError_
			return
		}

		personaId = info.Persona[0].ID
	}

	headers := map[string]string{
		"key":       key,
		"usuario":   usuario,
		"idPersona": fmt.Sprint(personaId),
	}

	urlcrud := "http://" + path + endpointHojaVida
	err := request.Get(urlcrud, headers, &hojaVida)
	if err != nil {
		logs.Error(urlcrud+", ", err)
		eval := "request.Get(urlcrud, headers, &info)"
		outputError = error_control.Error(funcion+eval, err, "502")
	}

	return
}
