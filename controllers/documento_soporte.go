package controllers

import (
	"sideap_mid/helpers"
	"sideap_mid/utils_oas/error_control"

	beego "github.com/beego/beego/v2/server/web"
)

// DocumentoSoporteController operations for Documento_soporte
type DocumentoSoporteController struct {
	beego.Controller
}

// URLMapping ...
func (c *DocumentoSoporteController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
}

// GetOne ...
// @Title GetOne
// @Description get Documento_soporte by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Documento_soporte
// @Failure 403 :id is empty
// @router /:id [get]
func (c *DocumentoSoporteController) GetOne() {

	defer error_control.ErrorControlController(c.Controller, "DocumentoSoporteController")

	doc, file, err := helpers.GetSoporte()
	if err == nil {
		resp := map[string]interface{}{
			"doc":  doc,
			"file": file,
		}
		c.Data["json"] = resp
	} else {
		panic(err)
	}

	c.ServeJSON()
}
