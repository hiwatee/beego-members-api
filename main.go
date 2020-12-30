package main

import (
	"beego-members-api/models"
	_ "beego-members-api/routers"
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	sqlConn, err := beego.AppConfig.String("sqlConn")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", sqlConn)
	if err != nil {
		fmt.Println(err)
	}
	orm.RegisterModel(
		new(models.User),
		new(models.Profile),
	)
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	web.Run()
}
