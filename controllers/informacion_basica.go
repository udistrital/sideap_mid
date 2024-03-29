package controllers

import (
	"errors"
	"sideap_mid/helpers"
	"sideap_mid/utils_oas/error_control"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// InformacionBasicaController operations for Informacion_basica
type InformacionBasicaController struct {
	beego.Controller
}

// URLMapping ...
func (c *InformacionBasicaController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title GeAll
// @Description get Informacion_basica by id
// @Param	codigo	query	string	true	"Codigo de abreviacion del documento de identidad"
// @Param	numero	query	string	true	"Numero de documento"
// @Success 200 {object} models.Informacion_basica
// @router / [get]
func (c *InformacionBasicaController) GetAll() {

	defer error_control.ErrorControlController(c.Controller, "InformacionBasicaController")

	codigo := c.GetString("codigo")
	numero := c.GetString("numero")

	if codigo == "" || numero == "" {
		err := errors.New("se debe especificar un tipo y número de documento")
		logs.Error(err)
		panic(map[string]interface{}{
			"funcion": `GetAll - codigo == "" || numero == ""`,
			"err":     err,
			"status":  "400",
		})
	}

	info, err := helpers.GetInformacionBasicaByDocumento(codigo, numero)
	if err == nil {
		c.Data["json"] = info
	} else {
		panic(err)
	}

	c.ServeJSON()

}
