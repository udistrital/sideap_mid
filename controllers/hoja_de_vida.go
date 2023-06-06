package controllers

import (
	"errors"

	"github.com/udistrital/sideap_mid/helpers"
	"github.com/udistrital/sideap_mid/utils_oas/error_control"

	beego "github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// HojaDeVidaController operations for Hoja_de_vida
type HojaDeVidaController struct {
	beego.Controller
}

// URLMapping ...
func (c *HojaDeVidaController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetOne ...
// @Title GetOne
// @Description get Hoja_de_vida by id
// @Param	codigo		query	string	false	"Codigo de abreviacion del documento de identidad"
// @Param	numero		query	string	false	"Numero de documento"
// @Param	personaId	query	int		false	"Numero de documento"
// @Success 200 {object} models.HojaVida
// @router / [get]
func (c *HojaDeVidaController) GetAll() {

	defer error_control.ErrorControlController(c.Controller, "HojaDeVidaController")

	personaId, _ := c.GetInt("personaId", 0)
	codigo := c.GetString("codigo")
	numero := c.GetString("numero")

	if personaId == 0 && (codigo == "" || numero == "") {
		err := errors.New("se debe especificar un tipo y número de documento o el id en el SIDEAP")
		logs.Error(err)
		panic(map[string]interface{}{
			"funcion": `GetAll - personaId == 0 && (codigo == "" || numero == "")`,
			"err":     err,
			"status":  "400",
		})
	}

	hv, err := helpers.GetHojaVida(personaId, codigo, numero)
	if err == nil {
		c.Data["json"] = hv
	} else {
		panic(err)
	}

	c.ServeJSON()
}
