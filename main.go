package main

import (
	_ "github.com/udistrital/sideap_mid/routers"

	api_status "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/auditoria"
	error_controller "github.com/udistrital/utils_oas/customerrorv2"

	beego "github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
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
