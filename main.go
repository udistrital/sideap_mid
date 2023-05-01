package main

import (
	_ "sideap_mid/routers"
	"sideap_mid/utils_oas/api_status"
	"sideap_mid/utils_oas/auditoria"
	"sideap_mid/utils_oas/error_controller"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {
	AllowedOrigins := []string{"*.udistrital.edu.co"}
	if beego.BConfig.RunMode == "dev" {
		AllowedOrigins = []string{"*"}
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: AllowedOrigins,
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length", "X-Total-Count"},
		AllowCredentials: true,
	}))

	beego.ErrorController(&error_controller.CustomErrorController{})
	api_status.Init()
	auditoria.InitMiddleware()

	beego.Run()
}
