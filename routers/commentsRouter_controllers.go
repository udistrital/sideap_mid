package routers

import (
	beego "github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["sideap_mid/controllers:DocumentoSoporteController"] = append(beego.GlobalControllerRouter["sideap_mid/controllers:DocumentoSoporteController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["sideap_mid/controllers:HojaDeVidaController"] = append(beego.GlobalControllerRouter["sideap_mid/controllers:HojaDeVidaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["sideap_mid/controllers:InformacionBasicaController"] = append(beego.GlobalControllerRouter["sideap_mid/controllers:InformacionBasicaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
