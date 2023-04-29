package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// HojaDeVidaController operations for Hoja_de_vida
type HojaDeVidaController struct {
	beego.Controller
}

// URLMapping ...
func (c *HojaDeVidaController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
}

// GetOne ...
// @Title GetOne
// @Description get Hoja_de_vida by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Hoja_de_vida
// @Failure 403 :id is empty
// @router /:id [get]
func (c *HojaDeVidaController) GetOne() {

}
