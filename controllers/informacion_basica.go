package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// InformacionBasicaController operations for Informacion_basica
type InformacionBasicaController struct {
	beego.Controller
}

// URLMapping ...
func (c *InformacionBasicaController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
}

// GetOne ...
// @Title GetOne
// @Description get Informacion_basica by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Informacion_basica
// @Failure 403 :id is empty
// @router /:id [get]
func (c *InformacionBasicaController) GetOne() {

}
