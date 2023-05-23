package controllers

import (
	beego "github.com/astaxie/beego"
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

}
