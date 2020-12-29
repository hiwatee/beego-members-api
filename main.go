package main

import (
	"beego-members-api/models"
	_ "beego-members-api/routers"
	"fmt"

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

	sqlConn := beego.AppConfig.String("sqlConn")

	orm.RegisterDataBase("default", "mysql", sqlConn)

	orm.RegisterModel(
		new(models.User),
	)

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}

}
