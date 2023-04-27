// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"sideap_mid/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/informacion_basica",
			beego.NSInclude(
				&controllers.InformacionBasicaController{},
			),
		),
		beego.NSNamespace("/hoja_de_vida",
			beego.NSInclude(
				&controllers.HojaDeVidaController{},
			),
		),
		beego.NSNamespace("/documento_soporte",
			beego.NSInclude(
				&controllers.DocumentoSoporteController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
