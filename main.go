package main

import (
	"beego-members-api/models"
	_ "beego-members-api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	setupDB()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func setupDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPass := beego.AppConfig.String("mysqlpass")
	mysqlHost := beego.AppConfig.String("mysqlurls")
	mysqldb := beego.AppConfig.String("mysqldb")

	orm.RegisterDataBase("default", "mysql", mysqlUser+":"+mysqlPass+"@"+mysqlHost+"/"+mysqldb+"?charset=utf8")

	orm.RegisterModel(
		new(models.User),
	)
}
